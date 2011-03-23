package main

import (
//	"rpc"
//	"net"
	"./client"
	"log"
	"bufio"
	"os"
	"../include/sfs"
	"fmt"
)

func main(){
	if len(os.Args) < 1 {
		log.Printf("Error: Run as \"client [masters ip]\"");
	}else {
		var fd int
		log.Printf("Test: creating hello\n");
		if len(os.Args) != 2 {
			//log.Printf("Test: correct usage is: client c masterMachine\n");
			fd = client.Open("hello", true, "127.0.0.1")
		} else{
			log.Printf("os.args[1] = " +  os.Args[1])
			fd = client.Open("hello", true, os.Args[1])
		}
		log.Printf("Test: writing hello\n");
		var data [32]byte
	//	st :=  "stealthy... "
		log.Printf("Test: Enter text to write to file:\n")
		var fromIn string 
		in := bufio.NewReader(os.Stdin) 
		fromIn, _ = in.ReadString('\n') 

		for i := 0; i < len(fromIn); i++{
			data[i] = fromIn[i]
		}
		_ =  client.Write(fd,  data );
		log.Printf("Test: reading hello\n")
		var readFromFile sfs.Chunk
		readFromFile,ret :=  client.Read(fd)
		if(ret==-1){
			log.Printf("Test: Read Failed")
		}
		printChunk(readFromFile)
		log.Printf("Test: done\n")	
	}
}

func printChunk ( toPrint sfs.Chunk ){
	fmt.Printf("Test: Chunk Read: ")
	fmt.Printf("\n")
	for i := 0; i < sfs.CHUNK_SIZE ; i++{
		fmt.Printf("%c", toPrint.Data[i])
	}
	fmt.Printf("\n")
}


