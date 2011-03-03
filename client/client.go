//this is the location for the client library
package client

import (
	"rpc"
	"os"
	"http"
//	"log"
//	"net"
	"fmt"
	"../include/sfs"
)

//NEED TO MAKE A TYPE FOR FILE
/*
	int fd
	list of chunks (preferably in order)
	list of serverlocations, one for each chunk 
	string filename
	int size

*/
//NEED TYPE FOR CHUNK TO HAVE SPACE FOR A FULL CHUNK
//defined in include/sfs???

/*create*/
//func Create(string filename) (int){


//	return 1;
//}

/* requestForAdditionalChunk */
//func RequestAdditionalChunk(String filename) (chunkIDNUMBER, chunkServersThatOwnFile){
//}

/* open */
func Open(filename string , write bool) (int){
	//open the file by the name filename
	//return an int giving the fd#.  if -1, it was a fail!
	//read == false  write == true

	//fmt.Printf("Attempting to Connect\n");
	client,err :=rpc.DialHTTP("tcp", "127.0.0.1:1338"); //IP needs to be changed to Master's IP
	if err != nil{
		fmt.Printf("Dial Failed");
		os.Exit(1);
	}
	fileInfo := new (sfs.OpenReturn);
	fileArgs := new (sfs.OpenArgs);
	fileArgs.Name = filename;
	masterCall := client.Go("Master.ReadOpen", &fileArgs,&fileInfo, nil);
	replyCall:= <-masterCall.Done
	//this is asynchronous, probably want to change it to synchronous
	if replyCall.Error!=nil{
		fmt.Printf("error in reply from rpc\n");
	}

	if fileInfo.New {
		fmt.Printf("\nNew!\n");
	}
	fmt.Printf("File Size = %d\n", fileInfo.Size)
	fmt.Printf("File Chunk = %d\n", fileInfo.Chunk)
	fmt.Printf("ServerLocation Port = %s\n", fileInfo.ServerLocation.String())
//// store info for file from  fileInfo into file type
//// return file descriptor arbitrary for now
	fd := 111; // 
	return fd;
}

/* close */
func Close(fd int) (int){
	//will have to tell this so that the master can be informed of the new file sizes...
	return 1;
}

/* read */
func Read (fd int, chunk int) (sfs.Chunk){
	//goes to chunk and gets a chunk of memory to read...
///*
	client,err :=rpc.Dial("tcp", "127.0.0.1:1338"); //IP needs to be changed to Master's IP
	if err != nil{
		fmt.Printf("Dial Failed");
		os.Exit(1);
	}
	fileInfo := new (sfs.ReadReturn);
	fileArgs := new (sfs.ReadArgs);
	fileArgs.ChunkID= 0;
	fileArgs.Offset = 1;
	fileArgs.Length = 1;
	chunkCall := client.Go("Server.Read", &fileArgs,&fileInfo, nil);
	replyCall:= <-chunkCall.Done
	//this is asynchronous, probably want to change it to synchronous
	if replyCall.Error!=nil{
		fmt.Printf("error in reply from rpc\n");
	}
	fmt.Printf("\nStatus = %d\n",fileInfo.Status);

	fmt.Printf("chunk = %d\n",fileInfo.Status);


//// store info for file from  fileInfo into file type
//// return file descriptor arbitrary for now
//*/


	return fileInfo.Data;
}

/* write */
func Write (fd int , chunk sfs.Chunk  ) (int){
	//we will need to write data to different blocks
	//the return indicates whether it was successful

///*
	client,err :=rpc.Dial("tcp", "127.0.0.1:1338"); //IP needs to be changed to Master's IP
	if err != nil{
		fmt.Printf("Dial Failed");
		os.Exit(1);
	}
	fileInfo := new (sfs.WriteReturn);
	fileArgs := new (sfs.WriteArgs);
//	fileArgs.Name = filename;
//	fileArgs.Name = filename;
	chunkCall := client.Go("Server.Write", &fileArgs,&fileInfo, nil);
	replyCall:= <-chunkCall.Done
	//this is asynchronous, probably want to change it to synchronous
	if replyCall.Error!=nil{
		fmt.Printf("error in reply from rpc\n");
	}

	fmt.Printf("\nChunkID: %d\n", fileInfo.ChunkID);
	fmt.Printf("Offset: %d\n", fileInfo.Offset);
	fmt.Printf("Length: %d\n", fileInfo.Length);
	fmt.Printf(": %d\n", fileInfo.Length);

//// store info for file from  fileInfo into file type
//// return file descriptor arbitrary for now

//*/
	return 1;	
}


/* seek - LATER */


/* append - LATER */


/* remove - LATER */


