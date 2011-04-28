//this is the location for the client library
package client

import (
	"container/vector"
	"rpc"
	"os"
	"log"
	"../include/sfs"
	"math"
//	"time"
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
	log.Println("Client: Master IP = ", masterAddr);
	master = masterAddr
}

func dialServer(address string) (*rpc.Client, os.Error){
	server,err := rpc.Dial("tcp", address)
	if(err != nil){
		log.Println("Client: Dial Error to", address, err.String())
	}
	return server,err
}

func Open(filename string , flag int ) (int){
	log.Println("Client: opening ", filename)
	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	defer client.Close()
	if err != nil{
		log.Println("Client: Dial Error ", err);
		return FAIL
	}else{
		_,openF :=  openFiles[filename]

		if !openF {
			fileInfo := new (sfs.OpenReturn)
			fileArgs := new (sfs.OpenArgs)
			fileArgs.Lock = false
			fileArgs.Name = filename
			if((flag & O_CREATE) == O_CREATE){
				log.Println("Client: Permissions for New file!")
				fileArgs.NewFile = true
			} else {
				fileArgs.NewFile = false
			}
			err := client.Call("Master.ReadOpen", &fileArgs,&fileInfo)
			if(err != nil){
				log.Println("Client: Open fail ", err)
				return sfs.FAIL
			}
			if fileInfo.New && (flag & O_CREATE) == O_CREATE {
				log.Println("Client: New file!")

			}else if !fileInfo.New  && (flag & O_CREATE) != O_CREATE   {
				log.Println("Client: Old file!")
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
			log.Println("Client: else of open!")
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

	log.Println("Client: **********READ BEGIN***********");
	nameAndPointer, present :=  openDescriptors[fd]
	var entireRead []byte

	if !present {
		log.Println("Client: fd does not exist");
		return entireRead, sfs.FAIL
	}

	filename := nameAndPointer.name
	filePtr := nameAndPointer.filePtr
	fdFile, present := openFiles[filename]

	if((nameAndPointer.permissions & O_RDONLY) != O_RDONLY){ //check if file was opened with Read permissions
		log.Println("Client: Cannot read without read permissions")
		return entireRead, FAIL
	}


	if (int(filePtr)+(size)) > int(fdFile.size) {
		entireRead = make([]byte,(fdFile.size - filePtr))
	}else {
		entireRead = make([]byte,size)
	}
	if !present {
		log.Println("Client: File not in open list!")
		return entireRead, FAIL
	}
	index := 0
	startIndex :=int( filePtr % sfs.CHUNK_SIZE)
	endIndex := int(cap(entireRead) + startIndex)
	if endIndex > int(fdFile.size) {
		endIndex = int(fdFile.size)
	}
	endChunk := int(math.Ceil((float64(filePtr)+float64(size))/sfs.CHUNK_SIZE))
	log.Println("Client: size = ", size)
	log.Println("Client: openFiles =", openFiles, "openDescriptors =", openDescriptors)
	log.Println("Client: index = %d, startIndex = ", index, startIndex)
	log.Println("Client: endIndex = %d, endChunk = ", endIndex, endChunk)
	log.Println("Client: filestarting size = ",fdFile.size)
	log.Println("Client: fileName ",fdFile.name)
	if endChunk >  int(math.Ceil(float64(fdFile.size)/float64(sfs.CHUNK_SIZE))) {
		endChunk = int(math.Ceil(float64(fdFile.size)/float64(sfs.CHUNK_SIZE)))
		log.Println("Client: file was smaller than read size")
	}
	var returned int
	for i := int(filePtr/sfs.CHUNK_SIZE); i<endChunk; i++ {
		chunkServerMirrors := fdFile.chunkInfo.At(i).(sfs.ChunkInfo).Servers
		numChunkServers := len(chunkServerMirrors)
		if (numChunkServers < 1) {
			log.Println("Client: Dial Failed in Read")
			return entireRead, sfs.FAIL
		}

		returned, bytesRead := GetChunk(*fdFile, i)

		if(returned == sfs.SUCCESS ){
			for k:=0; k<sfs.CHUNK_SIZE ; k++{
				if (index< endIndex && index>= startIndex ) {
					entireRead[index-startIndex] = bytesRead[k]
				}
				index++
			}
		}else {
			return entireRead, sfs.FAIL
		}
	}

	openDescriptors[fd].filePtr += uint64(size)
	if openDescriptors[fd].filePtr > fdFile.size {
		openDescriptors[fd].filePtr	=openFiles[filename].size
	}
	printByteSlice(entireRead)
	log.Println("Client: *********READ END*************");
	return  entireRead, returned;
}

/* write */
func Write (fd int, data []byte) (int){
	log.Println("Client: *************WRITE BEGIN**********");

	masterServ,masterErr  := rpc.Dial("tcp",master + ":1338")
	if masterErr != nil {
		log.Println("Client: dial fail: ", masterErr)
		return FAIL
	}
	if masterServ == nil {
		return FAIL
	}

	nameAndPointer, inMap :=  openDescriptors[fd]
	if !inMap {
		log.Println("Client: fd does not exist");
		return sfs.FAIL
	}
	filename := nameAndPointer.name
	filePtr := nameAndPointer.filePtr

	fdFile, inMap := openFiles[filename]
	if !inMap {
		log.Println("Client: File not in open list!")
		return FAIL
	}

	log.Println("Client: filestarting size = ",fdFile.size)
	log.Println("Client: fileName ",fdFile.name)
	if((nameAndPointer.permissions & O_WRONLY) != O_WRONLY){
		log.Println("Client: Cannot write without write permissions")
		return FAIL
	}

	indexWithinChunk := int( filePtr)%int(sfs.CHUNK_SIZE)
	chunkOffset := int(filePtr)/int(sfs.CHUNK_SIZE)
	var toWrite sfs.Chunk

	//special case if writing to middle of first chunk
	if  indexWithinChunk  > 0 {
		returned, bytesRead := GetChunk(*fdFile, chunkOffset)
		if(returned == sfs.SUCCESS){
			for i:= 0 ; i < indexWithinChunk ; i++ {
				toWrite.Data[i] = bytesRead[i]
			}
		}else{
			log.Println("Client: Dial Failed in GetChunk trying to get beginning of first chunk")
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
					log.Println("Client: Dial Failed in GetChunk trying to get beginning of first chunk")
					return FAIL
				}
				
			}

			if(fdFile.chunkInfo.Len() <= chunkOffset+1) {
				returned, toPush := AddChunks(fdFile.name, 1)
				if returned == sfs.FAIL {
					log.Println("AddChunk failed 1")
					return sfs.FAIL
				}
				fdFile.chunkInfo.Push(toPush)
				log.Println("Client: adding chunk x")
			}else{
				returned, toSet := AddChunks(fdFile.name, 1)
				if returned == sfs.FAIL {
					log.Println("AddChunk failed 1")
					return sfs.FAIL
				}
				fdFile.chunkInfo.Set(chunkOffset, toSet )
				log.Println("Client: adding chunk y")
			}
			fileArgs.Info = fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo)
			fileArgs.Data = toWrite

			if(len(fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers)<1){
				log.Println("fdFile.chunkInfo ", fdFile.chunkInfo)
			}

			servers := fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers;
			numChunkServers := len(servers)
			if (numChunkServers < 1) {
				log.Println("Client: Dial Failed in Write")
				return sfs.FAIL
			}

			log.Println("Client: numChunkServers ", numChunkServers);
			for j:=0; j < (numChunkServers); j++ {
				log.Println("Client:  j", j);
				client,err  := rpc.Dial("tcp",servers[0].String())
				for i:= 0 ; i < len(servers) ; i ++ {
					log.Println("Client: servers contains : " + servers[i].String() )
				}
				if err != nil ||  client == nil{
					log.Println("Client: Dial to chunk failed, returned bad client", err);
					//log.Println("Client: retrying dial in one second");
					numServers := len(fileArgs.Info.Servers)
					tmp := fileArgs.Info.Servers[numServers-1]
					for n:=0 ; n<numServers -1 ; n++{
						fileArgs.Info.Servers[n+1] = fileArgs.Info.Servers[n]
					}
					fileArgs.Info.Servers[0] = tmp
				//	time.Sleep(sfs.HEARTBEAT_WAIT/15)

					if j == numChunkServers-1 {
						return sfs.FAIL
					}
					continue
				}

				err = client.Call("Server.Write", &fileArgs,&fileInfo);
				client.Close()
				if err != nil{
					log.Println("Client: Server.Write failed:", err);
					if j == numChunkServers-1 {
						return sfs.FAIL
					}
					continue
				}
				if(fileInfo.Status!=0){
					log.Println("Client: Server.Write status non zero=",fileInfo.Status)
					if j == numChunkServers-1 {
						return sfs.FAIL
					}
					continue
				}
				break
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
				returned, toPush := AddChunks(fdFile.name, 1)
				if returned == sfs.FAIL {
					return sfs.FAIL
				}
				log.Println("Client: adding chunk zz")
				fdFile.chunkInfo.Push(toPush)
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
	log.Println("Client: file ending size ", openFiles[filename].size);
	log.Println("Client: fileInfo.Status = ", fileInfo.Status);
	log.Println("Client: ************WRITE END**************");
	return fileInfo.Status
}

func GetChunk(fdFile file,  chunkOffset int)(int, [sfs.CHUNK_SIZE]byte){
	log.Println("Client: Getting Chunk", fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).ChunkID)
	fileArgsRead := new (sfs.ReadArgs)
	fileInfoRead := new (sfs.ReadReturn)
	fileArgsRead.Nice = 1 // try things nicely first
	Servers := fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers
	numServers := len(Servers)
	for i:= 0 ; i < (numServers*2) ; i ++ {
		if i >= numServers {
			fileArgsRead.Nice = 0
		}

		client,err :=rpc.Dial("tcp",Servers[i%numServers].String())
		if client == nil {
			log.Println("Client: Dial Failed in GetChunk client == nil ")
			if i != (numServers*2)-1 {
				continue
			} else {
				log.Println("Client: missed continue in GetChunk ")
				return sfs.FAIL, fileInfoRead.Data.Data
			}
		}
		//defer client.Close()
		if err != nil {
			log.Println("Client: Dial Failed in GetChunk err != nil ", err)
			if i != (numServers*2-1) {
				continue
			}else{
				log.Println("Client: missed continue in GetChunk ")
				return sfs.FAIL, fileInfoRead.Data.Data
			}
		}
		
		fileArgsRead.ChunkID= fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).ChunkID;
		if fileArgsRead.ChunkID == 0 {
			log.Println("Client: ChunkID = 0, Chunk ID should never be 0")
		}
		
		err = client.Call("Server.Read", &fileArgsRead,&fileInfoRead);
		client.Close();
		if err!=nil{
			log.Printf("Client: Server.Read failed: %s, on server %s\n", err.String(),Servers[i%numServers].String());
			log.Printf("On try %d out of %d with %d# of servers\n",i,(numServers*2-1), numServers);
			if i != (numServers*2-1) {
				continue
			}
			return sfs.FAIL, fileInfoRead.Data.Data
		}

		if fileInfoRead.Status != sfs.SUCCESS {
			log.Println("Client: Read failed with status", fileInfoRead.Status)

		}else{
			break
		}

		log.Println("At end of function with status ", fileInfoRead.Status)
	}
	return sfs.SUCCESS, fileInfoRead.Data.Data

}

/* delete */
func Delete(filename string) (int){

	_ , present := openFiles[filename]
	if (!present ){
		log.Println("Client: filename does not exist ", filename);
		return FAIL
	}
	var dummy file
	openFiles[filename] = &dummy,false


	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	defer client.Close()

	if err != nil{
		log.Println("Client: Dial Error", err);
		return FAIL
	}else{
		fileArgs := new (sfs.DeleteArgs)
		fileInfo := new (sfs.DeleteReturn)
		fileArgs.Name = filename
		err := client.Call("Master.DeleteFile", &fileArgs,&fileInfo)
		if(err != nil || !fileInfo.Status){
			log.Println("Client: Delete fail ", fileInfo.Status, err)
			return FAIL
		}
	}
	return WIN
}

func Close(fd int) (int){
	n, inMap :=  openDescriptors[fd]
	filename := n.name
	if !inMap {
		log.Println("Client: fd does not exist");
		return FAIL
	}
	_ , present := openFiles[filename]
	if (!present ){
		log.Println("Client: filename does not exist", filename);
		return FAIL
	}
	var dummy nameAndPtr
	openDescriptors[fd]= &dummy,false
	return WIN
}

func ReadDir(path string) ([]string, int){

	readDirArgs := new (sfs.ReadDirArgs)
	readDirRet := new (sfs.ReadDirReturn)
	readDirArgs.Prefix = path
	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	defer client.Close()
	if err != nil{
		log.Println("Client: Dial Error", err);
		return readDirRet.FileNames,  FAIL
	}else{
		err = client.Call("Master.ReadDir", &readDirArgs, &readDirRet)
		if(err != nil){
			log.Println("Client: Read Dir fail ", err )
			return readDirRet.FileNames,  FAIL
		}
	}
	for i:= 0 ; i < len(readDirRet.FileNames) ; i++ {
		log.Println("Client: Got back directory: ", readDirRet.FileNames[i])
	}
	return readDirRet.FileNames,  WIN
}

/* seek */
func Seek (fd int, offset int, whence int) (int){
	n, inMap :=  openDescriptors[fd]
	filename := n.name
	log.Println("Client: filestarting size = ", openFiles[filename].size)
	log.Println("Client: fileName ",openFiles[filename].name)

	if !inMap {
		log.Println("Client: fd does not exist");
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
	log.Println("Client: bytes Read length = ", len(toPrint))
	log.Println("Client: bytes Read:")
	var st string

	for i := 0; i < len(toPrint); i++{
		st += string(toPrint[i])
	}

	log.Println(st)
}

func MakeDir(path string) (int) {

	var args sfs.MakeDirArgs
	var returnVal sfs.MakeDirReturn

	args.DirName = path

	masterConn,err := rpc.Dial("tcp", master + ":1338")
	defer masterConn.Close()
	if(err != nil){
		log.Println("Error Dialing Master(MakeDir):", err)
		return sfs.FAIL
	}

	err = masterConn.Call("Master.MakeDir",&args,&returnVal)
	if(err != nil){
		log.Println("Error Calling Master(MakeDir):", err)
		return sfs.FAIL
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
		log.Println("Error Dialing Master(Remove Dir):", err)
		return sfs.FAIL
	}

	err = masterConn.Call("Master.RemoveDir",&args,&returnVal)
	if(err != nil){
		log.Println("Error Calling Master(RemoveDir):", err)
		return sfs.FAIL
	}

	return returnVal.Status

}

func AddChunks(fileName string, numChunks uint64) (int, sfs.ChunkInfo) {

	var args sfs.GetNewChunkArgs
	var returnVal sfs.GetNewChunkReturn

	args.Name = fileName
	args.Count = numChunks

	masterConn,err := rpc.Dial("tcp", master + ":1338")

	if(err != nil){
		log.Println("Error Dialing Master(AddChunks):", err)
		return sfs.FAIL, returnVal.Info
	}

	err = masterConn.Call("Master.GetNewChunk",&args,&returnVal)
	if(err != nil){
		log.Println("Error Calling Master(AddChunks):", err)
		return sfs.FAIL, returnVal.Info
	}
	masterConn.Close()
	return sfs.SUCCESS, returnVal.Info

}

