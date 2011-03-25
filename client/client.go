//this is the location for the client library
package client

import (
	"rpc"
	"os"
	"log"
	"../include/sfs"
	"net"
)

//NEED TO MAKE A TYPE FOR FILE
/*
	int fd
	list of chunks (preferably in order)
	list of serverlocations, one for each chunk 
	string filename
	int size

*/
type file struct {
	chunk uint64
	serverAddress net.TCPAddr
}
var fd = 0
var openFiles = map[int] file{}

/* requestForAdditionalChunk */
//func RequestAdditionalChunk(String filename) (chunkIDNUMBER, chunkServersThatOwnFile){
//}
/* open */
func Open(filename string , write bool, master string  ) (int){
	//open the file by the name filename
	//return an int giving the fd#.  if -1, it was a fail!
	//read == false  write == true

	client,err :=rpc.Dial("tcp", master + ":1338"); //IP needs to be changed to Master's IP
	if err != nil{
		log.Printf("Client: Dial Failed");
		log.Printf("Client: Error", err.String());
		os.Exit(1);
	}
	fileInfo := new (sfs.OpenReturn);
	fileArgs := new (sfs.OpenArgs);
	fileArgs.Name = filename;
	client.Call("Master.ReadOpen", &fileArgs,&fileInfo);

	if fileInfo.New {
		log.Printf("\nClient: New file!\n");
	}
	log.Printf("Client: File Size = %d\n", fileInfo.Size)
	log.Printf("Client: File Chunk = %d\n", fileInfo.Chunk)
	log.Printf("Client: ServerLocation Port = %s\n", fileInfo.ServerLocation.String())
	fd++
	var nextFile file
	nextFile.chunk = fileInfo.Chunk
	nextFile.serverAddress = fileInfo.ServerLocation
	openFiles[fd] = nextFile
	return fd;

}

/* read */
func Read (fd int) (sfs.Chunk, int ){
	//goes to chunk and gets a chunk of memory to read...
///*
	fileInfo := new (sfs.ReadReturn);
	fileArgs := new (sfs.ReadArgs);
	fdFile, inMap := openFiles[fd]
	if !inMap {
		log.Printf("Client: File not in open list!\n")
		return fileInfo.Data, -1
	}
	client,err :=rpc.Dial("tcp",fdFile.serverAddress.String())
	if err != nil{
		log.Printf("Client: Dial Failed")
		return fileInfo.Data, -1
	}
	fileArgs.ChunkID= fdFile.chunk;
	fileArgs.Offset = 0;
	fileArgs.Length = sfs.CHUNK_SIZE;
	chunkCall := client.Go("Server.Read", &fileArgs,&fileInfo, nil);
	replyCall:= <-chunkCall.Done
	//this is asynchronous, probably want to change it to synchronous
	if replyCall.Error!=nil{
		log.Printf("Client: error in reply from rpc in read\n");
		return fileInfo.Data, -1
	}
	log.Printf("\nClient: Status = %d\n",fileInfo.Status);
	log.Printf("Client: Data = %d\n",fileInfo.Data);

	return  fileInfo.Data, fileInfo.Status;
}

/* write */
func Write (fd int , data [sfs.CHUNK_SIZE]byte  ) (int){
	//we will need to write data to different blocks
	//the return indicates whether it was successful
	fileInfo := new (sfs.WriteReturn);
	fileArgs := new (sfs.WriteArgs);
	fdFile, inMap := openFiles[fd]
	if !inMap {
		log.Printf("Client: File not in open list!\n")
		return -1
	}
	client,err :=rpc.Dial("tcp",fdFile.serverAddress.String()); //IP needs to be changed to Master's IP
	if err != nil{
		log.Printf("Client: Dial Failed");
		os.Exit(1);
	}
	fileArgs.Data.Data = data;
	fileArgs.ChunkID = fdFile.chunk;
	fileArgs.Offset = 0;
	fileArgs.Length = sfs.CHUNK_SIZE;
	chunkCall := client.Go("Server.Write", &fileArgs,&fileInfo, nil);
	replyCall:= <-chunkCall.Done
	//this is asynchronous, probably want to change it to synchronous
	if replyCall.Error!=nil{
		log.Printf("Client: error in reply from rpc in write\n");
		return -1
	}
//	log.Printf("\nClient: ChunkID: %d\n", fileArgs.ChunkID);
//	log.Printf("Client: Offset: %d\n", fileArgs.Offset);
//	log.Printf("Client: Length: %d\n", fileArgs.Length);
	log.Printf("Client: Data: %d\n", fileArgs.Length);

	return fileInfo.Status;
}



/* close */
//TODO

func Close(fd int) (int){
	//will have to tell this so that the master can be informed of the new file sizes...
// also remove from open files list 
	return 1;
}


/* seek */
//TODO
func Seek (fd int, chunkIndex int) (sfs.Chunk, int){
	//BASED Off of READ as they should be fairly similar...
	//goes to chunk and gets a chunk of memory to read...
///*
	fileInfo := new (sfs.ReadReturn);
	fileArgs := new (sfs.ReadArgs);
	fdFile, inMap := openFiles[fd]
	if !inMap {
		log.Printf("Client: File not in open list!\n")
		return fileInfo.Data, -1
	}
	client,err :=rpc.Dial("tcp",fdFile.serverAddress.String())
	if err != nil{
		log.Printf("Client: Dial Failed")
		return fileInfo.Data, -1
	}
	fileArgs.ChunkID= fdFile.chunk;
	fileArgs.Offset = 0;
	fileArgs.Length = sfs.CHUNK_SIZE;
	//ChunkCall := client.Go("Server.Read", &fileArgs,&fileInfo, nil);
	//replyCall:= <-chunkCall.Done
	//this is asynchronous, probably want to change it to synchronous
	//if replyCall.Error!=nil{
	//	log.Printf("Client: error in reply from rpc in read\n");
	//	return fileInfo.Data, -1
	//}
	//log.Printf("\nClient: Status = %d\n",fileInfo.Status);
	//log.Printf("Client: Data = %d\n",fileInfo.Data);

	return  fileInfo.Data, fileInfo.Status;
}

/* append - LATER */
//TODO
//func Append(fd int, data ???) (int){
//}

/* remove - LATER */
//TODO
func Remove(fd int) (int){
	return -1
}

/*create*/
//TODO
func Create(string filename) (int){
	return -1
}

