//this is the location for the client library
package client

import (
	"container/vector"
	"rpc"
	"os"
	"log"
//	"fmt"
//	"io/ioutil"
	"../include/sfs"
	"math"
)

const(
 FAIL = -1
 WIN  = 0
 O_RDONLY = 1
 O_WRONLY = 2
 O_RDWR = 3
 O_CREATE =4
 SEEK_SET = 1
 SEEK_CURR = 2
 SEEK_END = 4
)


type file struct {
	size uint64
	filePtr  uint64
	chunkInfo *vector.Vector
	name string
	permissions int
}

var master string
var fd int
var openFiles = map[int] (*file){}

func Initialize(masterAddr string){
	fd = 0
	log.Printf("Client: Master IP = %s", masterAddr);
	master = masterAddr
}

func Open(filename string , flag int ) (int){
	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	if err != nil{
		log.Printf("Client: Dial Error %s", err.String());
		return FAIL
	}else{
		fileInfo := new (sfs.OpenReturn)
		fileArgs := new (sfs.OpenArgs)
    	fileArgs.Lock = false
		fileArgs.Name = filename
		if((flag & O_CREATE) == O_CREATE){
			fileArgs.NewFile = true	
		} else {
			fileArgs.NewFile = false
		}
		err := client.Call("Master.ReadOpen", &fileArgs,&fileInfo)
		if(err != nil){
			log.Printf("Client: Open fail ", err )
			return FAIL
		}
		if fileInfo.New {
			log.Printf("Client: New file!\n")

		}else{
			log.Printf("Client: Old file!\n")
		}
		fd++
		var nextFile file
		nextFile.size = 0
		nextFile.filePtr = 0
		nextFile.permissions = flag
		nextFile.name = filename
		nextFile.chunkInfo = new(vector.Vector)
		for i := 0 ; i < cap(fileInfo.Chunk); i ++ {
			nextFile.chunkInfo.Push(fileInfo.Chunk[i])
			nextFile.size +=fileInfo.Chunk[i].Size
		}
		openFiles[fd] = &nextFile
		return fd;
	}
	return FAIL
}

/* read */
func Read (fd int, size int) ([]byte, int ){
	//goes to chunk and gets a chunk of memory to read...
	log.Printf("Client: **********READ BEGIN***********\n");
	fileInfo := new (sfs.ReadReturn)
	fileArgs := new (sfs.ReadArgs)
	fdFile, inMap := openFiles[fd]

//	log.Printf("Client: READ openFiles = %v",openFiles)
//	log.Printf("Client: READ FILESTARTING SIZE = %d",fdFile.size)

	var entireRead []byte
	if((fdFile.permissions & O_RDONLY) != O_RDONLY){ //check if file was opened with Read permissions
		log.Printf("Client: Cannot read without read permissions\n")
		return entireRead, FAIL
	}


	if (int(fdFile.filePtr)+(size)) > int(fdFile.size) {
		entireRead = make([]byte,(fdFile.size - fdFile.filePtr))
	}else {
		entireRead = make([]byte,size)
	}
	if !inMap {
		log.Printf("Client: File not in open list!\n")
		return entireRead, FAIL
	}
	index := 0
	startIndex :=int( fdFile.filePtr % sfs.CHUNK_SIZE)
	endIndex := int(int(size) + startIndex)
	if endIndex > int(fdFile.size) {
		endIndex = int(fdFile.size)
	}
	endChunk := int(math.Ceil((float64(fdFile.filePtr)+float64(size))/sfs.CHUNK_SIZE))
	log.Printf("Client: index = %d, startIndex = %d", index, startIndex)
	log.Printf("Client: endIndex = %d, endChunk = %d", endIndex, endChunk)
	log.Printf("Client: filestarting size = %d",fdFile.size)
	log.Printf("Client: fileName   %s",fdFile.name)
	if endChunk >  int(math.Ceil(float64(fdFile.size)/float64(sfs.CHUNK_SIZE))) {
		endChunk = int(math.Ceil(float64(fdFile.size)/float64(sfs.CHUNK_SIZE)))
		log.Printf("Client: file was smaller than read size")
	}
//	log.Println("Client:READ what is this %v", fileInfo)
	for i := int(fdFile.filePtr/sfs.CHUNK_SIZE); i<endChunk; i++ {
		if (len(fdFile.chunkInfo.At(i).(sfs.ChunkInfo).Servers) < 1) {
			log.Fatal("Client: No servers listed")
		}
		client,err :=rpc.Dial("tcp",fdFile.chunkInfo.At(i).(sfs.ChunkInfo).Servers[0].String())
		if err != nil{
			log.Printf("Client: Dial Failed in Read")
			return entireRead, FAIL
		}
		fileArgs.ChunkID= fdFile.chunkInfo.At(i).(sfs.ChunkInfo).ChunkID;
		if fileArgs.ChunkID == 0 {
			log.Printf("Client: ChunkID = 0, this shouldn't happen")
		}
		chunkCall := client.Go("Server.Read", &fileArgs,&fileInfo, nil);
		replyCall:= <-chunkCall.Done


		if replyCall.Error!=nil{
			log.Printf("Client: error in reply from rpc in read\n");
			return entireRead, FAIL
		}
		if(fileInfo.Status!=0){
			log.Printf("Client: CHUNK RETURNED BAD STATUS = %d\n",fileInfo.Status);
			break;
		}
		for j:=0; j<sfs.CHUNK_SIZE ; j++{
			if (index< endIndex && index>= startIndex ) {
				entireRead[index-startIndex] = fileInfo.Data.Data[j];
			}
			index++;
		}
	}
	openFiles[fd].filePtr += uint64(size)
	if openFiles[fd].filePtr > fdFile.size {
		openFiles[fd].filePtr =openFiles[fd].size
	}
	printByteSlice(entireRead)
	log.Printf("Client: *********READ END*************\n");
	return  entireRead, fileInfo.Status;
}

/* write */
func Write (fd int, data []byte) (int){
///*
	log.Printf("Client: *************WRITE BEGIN**********\n");

	fdFile, inMap := openFiles[fd]

	if((fdFile.permissions & O_WRONLY) != O_WRONLY){
		log.Printf("CLient: Cannot write without write permissions\n")
		return FAIL
	}

	if !inMap {
		log.Printf("Client: File not in open list!\n")
		return FAIL
	}
	var indexWithinChunk int
	indexWithinChunk = int( fdFile.filePtr)%int(sfs.CHUNK_SIZE)
	chunkOffset := int(fdFile.filePtr)/int(sfs.CHUNK_SIZE)
	sizeToWrite := uint64(len(data));
	var toWrite sfs.Chunk
	//this block prevents overwriting old blocks at indexes less than indexwithin chunk with 0s

	if  indexWithinChunk  > 0 {
		fileArgsRead := new (sfs.ReadArgs)
		fileInfoRead := new (sfs.ReadReturn)
		client,err :=rpc.Dial("tcp",fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers[0].String())
		if err != nil{
			log.Printf("Client: Dial Failed in Write 1")
			return FAIL
		}
		fileArgsRead.ChunkID= fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).ChunkID;
		if fileArgsRead.ChunkID == 0 {
			log.Printf("Client: ChunkID = 0, this shouldn't happen")
		}
		chunkCall := client.Go("Server.Read", &fileArgsRead,&fileInfoRead, nil);
		replyCall:= <-chunkCall.Done
		if replyCall.Error!=nil{
			log.Printf("Client: error in reply from rpc in read\n");
			return FAIL
		}
		for i:= 0 ; i < indexWithinChunk ; i++ {
			toWrite.Data[i] = fileInfoRead.Data.Data[i]
		}
	}
//	log.Printf("Client: WRITE IndexWithinChunk=%d, chunkOffset=%d",indexWithinChunk,chunkOffset)
//	log.Printf("Client: file size = %d, file fileptr = %d",fdFile.size, fdFile.filePtr)
//	log.Printf("Client: size To write  %d",sizeToWrite)
	fileArgs := new (sfs.WriteArgs);
	fileInfo := new (sfs.WriteReturn);
	for i:=0 ; i < len(data) ; i++  {
		toWrite.Data[int(indexWithinChunk)] = data[i]
		indexWithinChunk++
		if (indexWithinChunk == sfs.CHUNK_SIZE || i == len(data)-1 ){
			if(fdFile.chunkInfo.Len() <= chunkOffset+1) {
				fdFile.chunkInfo.Push(AddChunks(fdFile.name, 1))
				log.Printf("Client: adding chunk x")
			}else{
				fdFile.chunkInfo.Set(chunkOffset, AddChunks(fdFile.name, 1))
				log.Printf("Client: adding chunk y")
			}
//			if(fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).ChunkID ==0  ){
//				log.Printf("Client: adding chunk y")
//			}
			if((i != fdFile.chunkInfo.Len()-1)|| (sizeToWrite%sfs.CHUNK_SIZE==0)){
				//fileArgs.Length = sfs.CHUNK_SIZE;
			}else{
				//fileArgs.Length =uint(sizeToWrite)% uint(sfs.CHUNK_SIZE)
			}

			fileArgs.Info = fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo)
			fileArgs.Data = toWrite

			if(len(fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers)<1){
				log.Fatal("fdFile.chunkInfo ", fdFile.chunkInfo)

			}
			client,err  := rpc.Dial("tcp",fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers[0].String())
			if err != nil {
				log.Fatal("Client: dial fail:", err)
			}
			err = client.Call("Server.Write", &fileArgs,&fileInfo);
			if err != nil{
				log.Fatal("Client: error in reply from rpc in write\n");
			}
			if(fileInfo.Status!=0){
				log.Fatal("Client: Status non zero =",fileInfo.Status)
			}

			// reply to master
			fileInfo.Info.ChunkID = fileArgs.Info.ChunkID
			fileInfo.Info.Size = uint64(((indexWithinChunk-1)%int(sfs.CHUNK_SIZE)) +1)
			mapArgs := &sfs.MapChunkToFileArgs{fdFile.name, chunkOffset, fileInfo.Info}
			var mapRet sfs.MapChunkToFileReturn

			masterServ,err  := rpc.Dial("tcp",master + ":1338")
			if err != nil {
				log.Fatal("Client: dial fail:", err)
			}
			err = masterServ.Call("Master.MapChunkToFile", &mapArgs,&mapRet);
			if err != nil{
				log.Fatal("Client: error in reply from rpc in write\n");
			}

			chunkOffset++;
			indexWithinChunk =0

			if(fdFile.chunkInfo.Len() <  chunkOffset-1 ){
				log.Printf("Client: adding chunk zz")
				fdFile.chunkInfo.Push(AddChunks(fdFile.name, 1))
				for c := 0; c<sfs.CHUNK_SIZE ; c++ {
					toWrite.Data[c] = 0;
				}
			}

		}
	}
//*/

//	log.Printf("Client: Write openFiles = %v",openFiles)
	if fileInfo.Status !=FAIL {
		openFiles[fd].size = openFiles[fd].filePtr + uint64(len(data))
	}
	if openFiles[fd].size > fdFile.size {
		openFiles[fd].size = fdFile.size
	}
	openFiles[fd].filePtr += uint64(len(data))
	if openFiles[fd].filePtr > fdFile.size {
		openFiles[fd].filePtr = fdFile.size
	}
	log.Printf("Client: ************WRITE END**************\n");
	return fileInfo.Status
}

/* delete */
//TODO
func Delete(filename string) (int){
	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	if err != nil{
		log.Printf("Client: Dial Error %s", err.String());
		return FAIL
	}else{
		fileArgs := new (sfs.DeleteArgs)
		fileInfo := new (sfs.DeleteReturn)
		fileArgs.Name = filename
		err := client.Call("Master.DeleteFile", &fileArgs,&fileInfo)
		if(err != nil || !fileInfo.Status){
			log.Printf("Client: Delete fail ", fileInfo.Status, err )
			return FAIL
		}

	}



	return WIN
}

func Close(fd int) (int){
	_ , present := openFiles[fd]
	if (!present ){
		return FAIL
	}
	var x file
	openFiles[fd] = &x,false
	return WIN
}

//TODO
func ReadDir(path string) ([]string, int){

	readDirArgs := new (sfs.ReadDirArgs)
	readDirRet := new (sfs.ReadDirReturn)
	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	if err != nil{
		log.Printf("Client: Dial Error %s", err.String());
		return readDirRet.FileNames,  FAIL
	}else{

			err = client.Call("Master.ReadDir", &readDirArgs, &readDirRet)
			if(err != nil){
				log.Printf("Client: Read Dir fail ", err )
				return readDirRet.FileNames,  FAIL
			}
			
	}
		return readDirRet.FileNames,  WIN
}

/* seek */
//TODO
func Seek (fd int, offset int, whence int) (int){
	_ , present := openFiles[fd]
	if !present{
		return FAIL
	}
	if whence == SEEK_SET {
		openFiles[fd].filePtr = 0 + uint64(offset)
	}else if whence == SEEK_CURR {
		openFiles[fd].filePtr = openFiles[fd].filePtr + uint64(offset)

	}else if whence == SEEK_END {
		openFiles[fd].filePtr = openFiles[fd].size-1 - uint64(offset)
	}
	if openFiles[fd].filePtr > openFiles[fd].size {
		openFiles[fd].size = openFiles[fd].size
	}
	if openFiles[fd].filePtr < 0 {
		openFiles[fd].size = 0
	}
	return  int(openFiles[fd].filePtr);
}

func printByteSlice(toPrint []byte){
//	fmt.Printf("Client: bytes Read:")
//	fmt.Printf("\n")
	log.Printf("Client: bytes Read length = %d", len(toPrint))
	log.Printf("Client: bytes Read:")
	var st string
	for i := 0; i < len(toPrint); i++{
 //	  fmt.Printf("%c", toPrint[i])
	  st += string(toPrint[i])
	}
	log.Printf(st)
//	fmt.Printf(st)
//	fmt.Printf("\n")

}

func AddChunks(fileName string, numChunks uint64) (sfs.ChunkInfo) {

	var args sfs.GetNewChunkArgs
	var returnVal sfs.GetNewChunkReturn

	args.Name = fileName
	args.Count = numChunks

	masterConn,err := rpc.Dial("tcp", master + ":1338")
	if(err != nil){
		log.Printf("Error Dialing Master(AddChunks):", err.String())
		os.Exit(1)
	}
	
	err = masterConn.Call("Master.GetNewChunk",&args,&returnVal)
	if(err != nil){
		log.Printf("Error Calling Master(AddChunks):", err.String())
		os.Exit(1)
	}

	return returnVal.Info
	

}

//func ReadToFile(fileNameLocal string, fileNameServer string, master string)(int){
//	fd := Open(fileNameServer, true, master, uint64(len(f)))
// read , create local file, write to local file, close local file
//	return 0;
//}

