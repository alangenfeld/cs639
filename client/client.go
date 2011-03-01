//this is the location for the client library
package main

import (
	"http"
	"rpc"
	"os"
	"log"
	"net"
	"fmt"
)

//NEED TO MAKE A TYPE FOR FILE
/*
	int fd
	list of chunks (preferably in order)
	string filename

*/
//NEED TYPE FOR CHUNK TO HAVE SPACE FOR A FULL CHUNK

/* open */
func Open(String filename) (int){
	//open the file by the name filename
	//return an int giving the fd#.  if -1, it was a fail!
}

/* close */
func Close(int fd) (int){
	//will have to tell this so that the master can be informed of the new file sizes...
}

/* seek - LATER */


/* read */
func Read (int fd, int chunk) (chunk){
	//goes to chunk and gets a chunk of memory to read...
}

/* write */
func Write (int fd, chunk chunkData) (int){
	//we will need to write data to different blocks
	//the return indicates whether it was successful
}

/* append - LATER */


/* remove - LATER */


