//this is the location for the client library
package main

import (
	"rpc"
	"os"
	"log"
	"net"
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
func Create(String filename) (int){
		//fmt.Printf("Attempting to Connect\n");
		client,err :=rpc.Dial("tcp", "127.0.0.1:12345"); //IP needs to be changed to Master's IP
		if err != nil{
			fmt.Printf("Dial Failed");
			os.Exit(1);
		}
		fileInfo := new (CreateReturn);
		fileArgs := new (CreateArgs);
		fileArgs.Name = filename;
		masterCall := client.Go("type.functionname this needs to be decided on", &fileArgs,&fileInfo, nil);
		replyCall:= <-masterCall.Done
		//this is asynchronous, probably want to change it to synchronous
		if replyCall.Error!=nil{
			fmt.Printf("error in reply from rpc\n");
		}
		if fileInfo.Confirmation !=0 {
			fmt.Printf("Create Failed!\n");
			return -1;
		}
//// store info for file from  fileInfo into file type
//// return file descriptor arbitrary for now
		fd := 111; // 
		return fd;

}

/* requestForAdditionalChunk */
//func RequestAdditionalChunk(String filename) (chunkIDNUMBER, chunkServersThatOwnFile){
//}

/* open */
func Open(string filename, bool write) (int){
	//open the file by the name filename
	//return an int giving the fd#.  if -1, it was a fail!
	//read == false  write == true
}

/* close */
func Close(int fd) (int){
	//will have to tell this so that the master can be informed of the new file sizes...
}

/* read */
func Read (int fd, int chunk) (chunk){
	//goes to chunk and gets a chunk of memory to read...
}

/* write */
func Write (int fd, chunk chunkData) (int){
	//we will need to write data to different blocks
	//the return indicates whether it was successful
}


/* seek - LATER */


/* append - LATER */


/* remove - LATER */


