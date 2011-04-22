//this is the location for the client library
package client

import (
	"container/vector"
	"rpc"
	"os"
	"log"
	"../include/sfs"
	"math"
	"time"
)

const(
 FAIL = -1
 WIN  = 0
 O_RDONLY = 1
 O_WRONLY = 2
 O_RDWR = 3
 O_CREATE = 4
 SEEK_SET = 1
 SEEK_CURR = 2
 SEEK_END = 4
)

type nameAndPtr struct {
	name  string
	filePtr uint64
	permissions int
}

type file struct {
	size uint64
	chunkInfo *vector.Vector
	name string
}

var master string
var fd int
var openFiles = map[string] (*file){}
var openDescriptors = map[int] (*nameAndPtr){}

func Initialize(masterAddr string){
	fd = 0
	log.Printf("Client: Master IP = %s", masterAddr);
	master = masterAddr
}

func dialServer(address string) (*rpc.Client, os.Error){
	server,err := rpc.Dial("tcp", address)
	if(err != nil){
		log.Printf("Client: Dial Error to %s - %s", address, err.String())
	}
	return server,err
}

func Open(filename string , flag int ) (int){
	log.Printf("Client: opening %s!!\n", filename)
	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	defer client.Close()
	if err != nil{
		log.Printf("Client: Dial Error %s", err.String());
		return FAIL
	}else{
		_,openF :=  openFiles[filename]

		if !openF {
			fileInfo := new (sfs.OpenReturn)
			fileArgs := new (sfs.OpenArgs)
			fileArgs.Lock = false
			fileArgs.Name = filename
			if((flag & O_CREATE) == O_CREATE){
				log.Printf("Client: Permissions for New file!\n")
				fileArgs.NewFile = true
			} else {
				fileArgs.NewFile = false
			}
			err := client.Call("Master.ReadOpen", &fileArgs,&fileInfo)
			if(err != nil){
				log.Printf("Client: Open fail ", err )
				return sfs.FAIL
			}
			if fileInfo.New && (flag & O_CREATE) == O_CREATE {
				log.Printf("Client: New file!\n")

			}else if !fileInfo.New  && (flag & O_CREATE) != O_CREATE   {
				log.Printf("Client: Old file!\n")
			}else {
				return sfs.FAIL
			}

			fd++
			var nextFile file
			nextFile.size = 0
			nextFile.name = filename

			nextFile.chunkInfo = new(vector.Vector)
			for i := 0 ; i < cap(fileInfo.Chunk); i ++ {
				nextFile.chunkInfo.Push(fileInfo.Chunk[i])
				nextFile.size +=fileInfo.Chunk[i].Size
			}
			openFiles[filename] = &nextFile
			var d nameAndPtr
			d.name = filename
			d.filePtr = 0
			d.permissions = flag
			openDescriptors[fd] = &d
		}else{
			log.Printf("Client: else of open!\n")
			fd++
			var d nameAndPtr
			d.name = filename
			d.permissions = flag
			d.filePtr = 0
			openDescriptors[fd] = &d
		}
		return fd;
	}

	return sfs.FAIL
}

/* read */
func Read (fd int, size int) ([]byte, int ){

	log.Printf("Client: **********READ BEGIN***********\n");
	fileInfo := new (sfs.ReadReturn)
	fileArgs := new (sfs.ReadArgs)
	nameAndPointer, present :=  openDescriptors[fd]
	var entireRead []byte

	if !present {
		log.Printf("Client: fd does not exist\n");
		return entireRead, sfs.FAIL
	}

	filename := nameAndPointer.name
	filePtr := nameAndPointer.filePtr
	fdFile, present := openFiles[filename]

	if((nameAndPointer.permissions & O_RDONLY) != O_RDONLY){ //check if file was opened with Read permissions
		log.Printf("Client: Cannot read without read permissions\n")
		return entireRead, FAIL
	}


	if (int(filePtr)+(size)) > int(fdFile.size) {
		entireRead = make([]byte,(fdFile.size - filePtr))
	}else {
		entireRead = make([]byte,size)
	}
	if !present {
		log.Printf("Client: File not in open list!\n")
		return entireRead, FAIL
	}
	index := 0
	startIndex :=int( filePtr % sfs.CHUNK_SIZE)
	endIndex := int(cap(entireRead) + startIndex)
	if endIndex > int(fdFile.size) {
		endIndex = int(fdFile.size)
	}
	endChunk := int(math.Ceil((float64(filePtr)+float64(size))/sfs.CHUNK_SIZE))
	log.Printf("Client: size = %d\n", size)
	log.Printf("Client: openFiles = %v  openDescriptors = %v\n", openFiles, openDescriptors)
	log.Printf("Client: index = %d, startIndex = %d", index, startIndex)
	log.Printf("Client: endIndex = %d, endChunk = %d", endIndex, endChunk)
	log.Printf("Client: filestarting size = %d",fdFile.size)
	log.Printf("Client: fileName   %s",fdFile.name)
	if endChunk >  int(math.Ceil(float64(fdFile.size)/float64(sfs.CHUNK_SIZE))) {
		endChunk = int(math.Ceil(float64(fdFile.size)/float64(sfs.CHUNK_SIZE)))
		log.Printf("Client: file was smaller than read size")
	}

	for i := int(filePtr/sfs.CHUNK_SIZE); i<endChunk; i++ {

		chunkServerMirrors := fdFile.chunkInfo.At(i).(sfs.ChunkInfo).Servers
		numChunkServers := len(chunkServerMirrors)

		if (numChunkServers < 1) {
				log.Printf("Client: Dial Failed in Read")
				return entireRead, FAIL
		}

		fileArgs.Nice = 1 //default to reading "nicely"

		for j:=0; j < (numChunkServers*2); j++ {

			if(j >= numChunkServers) {//if we have tried all servers "nicely", start forcing reads
				fileArgs.Nice = 0
			}

			client,err := dialServer(chunkServerMirrors[j % numChunkServers].String())
			log.Printf( " dialed server : %s", chunkServerMirrors[j % numChunkServers].String())
			if err != nil || client == nil{
				log.Printf("Client: returned error or nil client\n")
				log.Printf("Client: retrying dial in one second\n")
				time.Sleep(sfs.HEARTBEAT_WAIT/15)
				if i == endChunk-1 {
					return entireRead, fileInfo.Status
				}
				continue
			}
			defer client.Close()

			fileArgs.ChunkID= fdFile.chunkInfo.At(i).(sfs.ChunkInfo).ChunkID;
			if fileArgs.ChunkID == 0 {
				log.Printf("Client: ChunkID = 0, invalid number so this shouldn't happen")
			}
		log.Printf("Client: 212")
			err = client.Call("Server.Read", &fileArgs, &fileInfo);
		log.Printf("Client: 214")

			if err!=nil{
				log.Printf("Client: error reading from Chunk Server - %s\n",chunkServerMirrors[j % numChunkServers].String());
				if i == endChunk-1 {
					log.Printf("Client: i == endChunk -1 ???")
					return entireRead, fileInfo.Status

				}
				continue
			}
			if fileInfo.Status == sfs.BUSY{
				log.Printf("Client: retrying dial to new server in one second\n")
				time.Sleep(sfs.HEARTBEAT_WAIT/15)
				client.Close()
			}

	log.Printf("Client: fileArgs.Nice = %d", fileArgs.Nice)
	log.Printf("Client: fileInfo.status = %d", fileInfo.Status)
			if(fileInfo.Status == sfs.SUCCESS ){
				for k:=0; k<sfs.CHUNK_SIZE ; k++{
					if (index< endIndex && index>= startIndex ) {
						entireRead[index-startIndex] = fileInfo.Data.Data[k]
					}
					index++
				}

				break //successfully read chunk.
			}
			if fileInfo.Status != sfs.SUCCESS {
				client.Close()
			}
			log.Printf("Client: Chunk returned status %d\n",fileInfo.Status);

		}


		if fileInfo.Status == sfs.FAIL { //case where chunk was never read successfully. 
			log.Printf("Client: chunk returned bad status\n");
			return entireRead, fileInfo.Status //return early b/c we failed to read chunk
		}
	}

	openDescriptors[fd].filePtr += uint64(size)
	if openDescriptors[fd].filePtr > fdFile.size {
		openDescriptors[fd].filePtr	=openFiles[filename].size
	}
	printByteSlice(entireRead)
	log.Printf("Client: *********READ END*************\n");
	return  entireRead, fileInfo.Status;
}

/* write */
func Write (fd int, data []byte) (int){
	log.Printf("Client: *************WRITE BEGIN**********\n");

	masterServ,masterErr  := rpc.Dial("tcp",master + ":1338")
	if masterErr != nil {
		log.Printf("Client: dial fail: %s", masterErr)
		return FAIL
	}
	if masterServ == nil {
		return FAIL
	}

	nameAndPointer, inMap :=  openDescriptors[fd]
	if !inMap {
		log.Printf("Client: fd does not exist\n");
		return sfs.FAIL
	}
	filename := nameAndPointer.name
	filePtr := nameAndPointer.filePtr

	fdFile, inMap := openFiles[filename]
	if !inMap {
		log.Printf("Client: File not in open list!\n")
		return FAIL
	}

	log.Printf("Client: filestarting size = %d",fdFile.size)
	log.Printf("Client: fileName   %s",fdFile.name)
	if((nameAndPointer.permissions & O_WRONLY) != O_WRONLY){
		log.Printf("Client: Cannot write without write permissions\n")
		return FAIL
	}

	indexWithinChunk := int( filePtr)%int(sfs.CHUNK_SIZE)
	chunkOffset := int(filePtr)/int(sfs.CHUNK_SIZE)
	var toWrite sfs.Chunk

	if  indexWithinChunk  > 0 {
		returned, bytesRead := GetChunk(*fdFile, chunkOffset)
		if(returned == sfs.SUCCESS){
			for i:= 0 ; i < indexWithinChunk ; i++ {
				toWrite.Data[i] = bytesRead[i]
			}
		}else{
			log.Printf("Client: Dial Failed in GetChunk trying to get beginning of first chunk")
			return FAIL
		}
	}
	fileArgs := new (sfs.WriteArgs);
	fileInfo := new (sfs.WriteReturn);
	for i:=0 ; i < len(data) ; i++  {
		toWrite.Data[int(indexWithinChunk)] = data[i]
		indexWithinChunk++
		if (indexWithinChunk == sfs.CHUNK_SIZE || i == len(data)-1 ){

//special case if write to middle of last chunk
			if  i == len(data)-1  && fdFile.size > filePtr && indexWithinChunk != sfs.CHUNK_SIZE {

				returned, bytesRead := GetChunk(*fdFile, chunkOffset)
				if(returned == sfs.SUCCESS){
					for i:= indexWithinChunk ; i < sfs.CHUNK_SIZE ; i++ {
						toWrite.Data[i] = bytesRead[i]
					}
				}else{
					log.Printf("Client: Dial Failed in GetChunk trying to get beginning of first chunk")
					return FAIL
				}

			}

			if(fdFile.chunkInfo.Len() <= chunkOffset+1) {
				fdFile.chunkInfo.Push(AddChunks(fdFile.name, 1))
				log.Printf("Client: adding chunk x")
			}else{
				fdFile.chunkInfo.Set(chunkOffset, AddChunks(fdFile.name, 1))
				log.Printf("Client: adding chunk y")
			}
			fileArgs.Info = fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo)
			fileArgs.Data = toWrite

			if(len(fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers)<1){
				log.Printf("fdFile.chunkInfo %v", fdFile.chunkInfo)
			}


			servers := fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers;

			numChunkServers := len(servers)

			if (numChunkServers < 1) {
					log.Printf("Client: Dial Failed in Read")
					return FAIL
			}

			for j:=0; j < (numChunkServers); j++ {

				client,err  := rpc.Dial("tcp",servers[j].String())
				if err != nil ||  client == nil{
					log.Println("Client: Dial to chunk failed, returned bad client");
					log.Println("Client: retrying dial in one second");

					tmp := fileArgs.Info.Servers[sfs.NREPLICAS-1]
					for n:=0 ; n<sfs.NREPLICAS-1 ; n++{
						fileArgs.Info.Servers[n+1] = fileArgs.Info.Servers[n]
					}
					fileArgs.Info.Servers[0] = tmp
					time.Sleep(sfs.HEARTBEAT_WAIT/15)
					continue
				}


				err = client.Call("Server.Write", &fileArgs,&fileInfo);
				client.Close()
				if err != nil{
					log.Println("Client: Server.Write failed:", err);
					continue
				}
				if(fileInfo.Status!=0){
					log.Println("Client: Server.Write status non zero=",fileInfo.Status)
					return FAIL
				}
			}


			// reply to master
			fileInfo.Info.ChunkID = fileArgs.Info.ChunkID
			fileInfo.Info.Size = uint64(((indexWithinChunk-1)%int(sfs.CHUNK_SIZE)) +1)
			mapArgs := &sfs.MapChunkToFileArgs{fdFile.name, chunkOffset, fileInfo.Info}
			var mapRet sfs.MapChunkToFileReturn



			err := masterServ.Call("Master.MapChunkToFile", &mapArgs,&mapRet);
			if err != nil{
				log.Println("Client: Master.MapChunkToFile failed:", err);
				return FAIL
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

	if masterServ != nil {
		masterServ.Close()
	}


	if fileInfo.Status !=FAIL && (openDescriptors[fd].filePtr+uint64(len(data)) > openFiles[filename].size) {
		openFiles[filename].size = openDescriptors[fd].filePtr + uint64(len(data))
	}
	if openFiles[filename].size > fdFile.size {
		openFiles[filename].size = fdFile.size
	}
	openDescriptors[fd].filePtr += uint64(len(data))
	if openDescriptors[fd].filePtr > fdFile.size {
		openDescriptors[fd].filePtr = fdFile.size
	}
	log.Printf("Client: file ending size %d", openFiles[filename].size);
	log.Printf("Client: ************WRITE END**************\n");
	log.Printf("Client: fileInfo.Status = %d \n", fileInfo.Status);
	log.Printf("Client: ************WRITE END**************\n");
	return fileInfo.Status
}

func GetChunk(fdFile file,  chunkOffset int)(int, [32]byte){
		log.Printf("Client: in Get Chunk\n")
		fileArgsRead := new (sfs.ReadArgs)
		fileInfoRead := new (sfs.ReadReturn)
		fileArgsRead.Nice = 1 // try things nicely first
		for i:= 0 ; i < (sfs.NREPLICAS*2) ; i ++ {
			if i == sfs.NREPLICAS {
				fileArgsRead.Nice = 0
			}
			client,err :=rpc.Dial("tcp",fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers[i%sfs.NREPLICAS].String())
			if client == nil {
				log.Printf("Client: Dial Failed in GetChunk client == nil %s", err.String())
				if i != (sfs.NREPLICAS*2-1) {
					continue
				}
				return FAIL, fileInfoRead.Data.Data
			}
			defer client.Close()
			if err != nil {
				log.Printf("Client: Dial Failed in GetChunk err != nil %s", err.String())
				if i != (sfs.NREPLICAS*2-1) {
					continue
				}
				return FAIL, fileInfoRead.Data.Data
			}
			fileArgsRead.ChunkID= fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).ChunkID;
			if fileArgsRead.ChunkID == 0 {
				log.Printf("Client: ChunkID = 0, Chunk ID should never be 0")
			}
			chunkCall := client.Go("Server.Read", &fileArgsRead,&fileInfoRead, nil);
			replyCall:= <-chunkCall.Done
			if replyCall.Error!=nil{
				log.Printf("Client: Server.Read failed:\n");
				return FAIL, fileInfoRead.Data.Data
			}
		}
		return sfs.SUCCESS, fileInfoRead.Data.Data

}

/* delete */
//TODO
func Delete(filename string) (int){

	_ , present := openFiles[filename]
	if (!present ){
		log.Printf("Client: filename does not exist  %s\n", filename);
		return FAIL
	}
	var dummy file
	openFiles[filename] = &dummy,false


	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	defer client.Close()

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
	n, inMap :=  openDescriptors[fd]
	filename := n.name
	if !inMap {
		log.Printf("Client: fd does not exist\n");
		return FAIL
	}
	_ , present := openFiles[filename]
	if (!present ){
		log.Printf("Client: filename does not exist  %s\n", filename);
		return FAIL
	}
	var dummy nameAndPtr
	openDescriptors[fd]= &dummy,false
	return WIN
}

//TODO
func ReadDir(path string) ([]string, int){

	readDirArgs := new (sfs.ReadDirArgs)
	readDirRet := new (sfs.ReadDirReturn)
	readDirArgs.Prefix = path
	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	defer client.Close()
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
	for i:= 0 ; i < len(readDirRet.FileNames) ; i++ {
		log.Printf("Client: Got back directory: %s", readDirRet.FileNames[i])
	}
	return readDirRet.FileNames,  WIN
}

/* seek */
//TODO
func Seek (fd int, offset int, whence int) (int){
	n, inMap :=  openDescriptors[fd]
	filename := n.name
	log.Printf("Client: filestarting size = %d", openFiles[filename].size)
	log.Printf("Client: fileName   %s",openFiles[filename].name)

	if !inMap {
		log.Printf("Client: fd does not exist\n");
		return FAIL
	}
	_ , present := openFiles[filename]
	if !present{
		return FAIL
	}
	filePtr := int(openDescriptors[fd].filePtr)
	if whence == SEEK_SET {
		filePtr = 0 + offset
	}else if whence == SEEK_CURR {
		filePtr = filePtr + offset
	}else if whence == SEEK_END {
		filePtr = int(openFiles[filename].size) + offset
	}
	openDescriptors[fd].filePtr = uint64(filePtr)
	if filePtr > int(openFiles[filename].size) {
		openDescriptors[fd].filePtr = openFiles[filename].size
	}
	if filePtr < 0 {
		openDescriptors[fd].filePtr = 0
	}
	return  int(openDescriptors[fd].filePtr);
}

func printByteSlice(toPrint []byte){
	log.Printf("Client: bytes Read length = %d", len(toPrint))
	log.Printf("Client: bytes Read:")
	var st string

	for i := 0; i < len(toPrint); i++{
	  st += string(toPrint[i])
	}

	log.Printf(st)
}

func MakeDir(path string) (int) {

	var args sfs.MakeDirArgs
	var returnVal sfs.MakeDirReturn

	args.DirName = path

	masterConn,err := rpc.Dial("tcp", master + ":1338")
	defer masterConn.Close()
	if(err != nil){
		log.Printf("Error Dialing Master(AddChunks):", err.String())
		os.Exit(1)
	}

	err = masterConn.Call("Master.MakeDir",&args,&returnVal)
	if(err != nil){
		log.Printf("Error Calling Master(MakeDir):", err.String())
		os.Exit(1)
	}

	return returnVal.Status

}


func RemoveDir(path string) (int) {

	var args sfs.RemoveDirArgs
	var returnVal sfs.RemoveDirReturn

	args.DirName = path

	masterConn,err := rpc.Dial("tcp", master + ":1338")
	defer masterConn.Close()
	if(err != nil){
		log.Printf("Error Dialing Master(AddChunks):", err.String())
		os.Exit(1)
	}

	err = masterConn.Call("Master.RemoveDir",&args,&returnVal)
	if(err != nil){
		log.Printf("Error Calling Master(RemoveDir):", err.String())
		os.Exit(1)
	}

	return returnVal.Status

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
	masterConn.Close()
	return returnVal.Info

}

