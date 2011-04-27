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
			err := client.Call("Proxy.ReadOpen", &fileArgs,&fileInfo)
			//err := client.Call("Master.ReadOpen", &fileArgs,&fileInfo)
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
	var returned int
	for i := int(filePtr/sfs.CHUNK_SIZE); i<endChunk; i++ {
		chunkServerMirrors := fdFile.chunkInfo.At(i).(sfs.ChunkInfo).Servers
		numChunkServers := len(chunkServerMirrors)
		if (numChunkServers < 1) {
				log.Printf("Client: Dial Failed in Read")
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
	log.Printf("Client: *********READ END*************\n");
	return  entireRead, returned;
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

//special case if writing to middle of first chunk
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
				returned, toPush := AddChunks(fdFile.name, 1)
				if returned == sfs.FAIL {
					log.Printf("AddChunk failed 1")
					return sfs.FAIL
				}
				fdFile.chunkInfo.Push(toPush)
				log.Printf("Client: adding chunk x")
			}else{
				returned, toSet := AddChunks(fdFile.name, 1)
				if returned == sfs.FAIL {
					log.Printf("AddChunk failed 1")
					return sfs.FAIL
				}
				fdFile.chunkInfo.Set(chunkOffset, toSet )
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
					log.Printf("Client: Dial Failed in Write")
					return sfs.FAIL
			}

			log.Printf("Client: numChunkServers %d", numChunkServers);
			for j:=0; j < (numChunkServers); j++ {
				log.Printf("Client:  j %d", j);
				client,err  := rpc.Dial("tcp",servers[j].String())
				if err != nil ||  client == nil{
					log.Println("Client: Dial to chunk failed, returned bad client");
					log.Println("Client: retrying dial in one second");
					numServers := len(fileArgs.Info.Servers)
					tmp := fileArgs.Info.Servers[numServers-1]
					for n:=0 ; n<numServers-1 ; n++{
						fileArgs.Info.Servers[n+1] = fileArgs.Info.Servers[n]
					}
					fileArgs.Info.Servers[0] = tmp
					time.Sleep(sfs.HEARTBEAT_WAIT/15)
					if j == numChunkServers-1 {
						return sfs.FAIL
					}
					continue
				}

				err = client.Call("ChunkProxy.Write", &fileArgs,&fileInfo);
				//err = client.Call("Server.Write", &fileArgs,&fileInfo);
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

			err := masterServ.Call("Proxy.MapChunkToFile", &mapArgs,&mapRet);
			//err := masterServ.Call("Master.MapChunkToFile", &mapArgs,&mapRet);
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
				log.Printf("Client: adding chunk zz")
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
	log.Printf("Client: file ending size %d", openFiles[filename].size);
	log.Printf("Client: fileInfo.Status = %d \n", fileInfo.Status);
	log.Printf("Client: ************WRITE END**************\n");
	return fileInfo.Status
}

func GetChunk(fdFile file,  chunkOffset int)(int, [sfs.CHUNK_SIZE]byte){
	log.Printf("Client: in Get Chunk\n")
	fileArgsRead := new (sfs.ReadArgs)
	fileInfoRead := new (sfs.ReadReturn)
	fileArgsRead.Nice = 1 // try things nicely first
	Servers := fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).Servers
	numServers := len(Servers)
		for i:= 0 ; i < (numServers*2) ; i ++ {
		log.Printf("trying server %s\n",Servers[i%numServers].String())
		log.Printf("On try %d out of %d with %d# of servers\n",i,((numServers*2)-1),numServers);
			if i >= numServers {
				fileArgsRead.Nice = 0
			}
//			for j := 0 ; j < numServers ; j ++ {
//				log.Printf("Client: server: %s", Servers[j%numServers].String())
//			}
			client,err :=rpc.Dial("tcp",Servers[i%numServers].String())
			if client == nil {
				log.Printf("Client: Dial Failed in GetChunk client == nil ")
					log.Printf("On try %d out of %d with %d# of servers\n",i,(numServers*2-1),numServers);
				if i != (numServers*2)-1 {
					continue
				}else {
					log.Printf("Client: missed continue in GetChunk ")
					return sfs.FAIL, fileInfoRead.Data.Data
				}
			}
			//defer client.Close()
			if err != nil {
				log.Printf("Client: Dial Failed in GetChunk err != nil %s", err.String())
				if i != (numServers*2-1) {
					continue
				}else{
					log.Printf("Client: missed continue in GetChunk ")
					return sfs.FAIL, fileInfoRead.Data.Data
				}
			}
			log.Println("Dial succeeded")
			fileArgsRead.ChunkID= fdFile.chunkInfo.At(chunkOffset).(sfs.ChunkInfo).ChunkID;
			if fileArgsRead.ChunkID == 0 {
				log.Printf("Client: ChunkID = 0, Chunk ID should never be 0")
			}
			log.Println("calling")
			err = client.Call("ChunkProxy.Read", &fileArgsRead,&fileInfoRead);
			//err = client.Call("Server.Read", &fileArgsRead,&fileInfoRead);
			log.Println("back from call")
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
				log.Printf("Client: retrying dial to new server in one second\n")
			//	time.Sleep(sfs.HEARTBEAT_WAIT/15)
				//client.Close()
			}else{
				break
			}

			log.Printf("At end of function with status %d\n",fileInfoRead.Status)
		}
		return sfs.SUCCESS, fileInfoRead.Data.Data

}

/* delete */
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
		err := client.Call("Proxy.DeleteFile", &fileArgs,&fileInfo)
		//err := client.Call("Master.DeleteFile", &fileArgs,&fileInfo)
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
		err = client.Call("Proxy.ReadDir", &readDirArgs, &readDirRet)
		//err = client.Call("Master.ReadDir", &readDirArgs, &readDirRet)
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
		log.Printf("Error Dialing Master(MakeDir):", err.String())
		return sfs.FAIL
	}

	err = masterConn.Call("Proxy.MakeDir",&args,&returnVal)
	//err = masterConn.Call("Master.MakeDir",&args,&returnVal)
	if(err != nil){
		log.Printf("Error Calling Master(MakeDir):", err.String())
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
		log.Printf("Error Dialing Master(Remove Dir):", err.String())
		return sfs.FAIL
	}

	err = masterConn.Call("Proxy.RemoveDir",&args,&returnVal)
	//err = masterConn.Call("Master.RemoveDir",&args,&returnVal)
	if(err != nil){
		log.Printf("Error Calling Master(RemoveDir):", err.String())
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
		log.Printf("Error Dialing Master(AddChunks):", err.String())
		return sfs.FAIL, returnVal.Info
	}

	err = masterConn.Call("Proxy.GetNewChunk",&args,&returnVal)
	//err = masterConn.Call("Master.GetNewChunk",&args,&returnVal)
	if(err != nil){
		log.Printf("Error Calling Master(AddChunks):", err.String())
		return sfs.FAIL, returnVal.Info
	}
	masterConn.Close()
	return sfs.SUCCESS, returnVal.Info

}

