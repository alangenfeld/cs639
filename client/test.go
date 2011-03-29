package main

import (
//	"rpc"
//	"net"
//	"../include/sfs"
	"./client"
	"log"
	"bufio"
	"os"
	"fmt"
	"container/vector"
)

func main(){
	if len(os.Args) < 1 {
		log.Printf("Error: Run as \"client [masters ip]\"");
	}else {
		var fd int
		var data vector.Vector
		var fromIn string
		log.Printf("Test: Enter text to write to file:\n")
		in := bufio.NewReader(os.Stdin)
		fromIn, _ = in.ReadString('\n')
		for i := 0; i < len(fromIn); i++{
			data.Push(fromIn[i])
		}
		log.Printf("Test: creating hello\n");
		if len(os.Args) != 2 {
			//log.Printf("Test: correct usage is: client c masterMachine\n");
			fd = client.Open("127.0.0.1", "hello", 1) )
		} else{
			log.Printf("os.args[1] = " +  os.Args[1])
			fd = client.Open(os.Args[1], "hello", 1))
		}
		log.Printf("Test: writing hello\n");
		_ =  client.Write(fd,  data );
		log.Printf("Test: reading hello\n")
		readFromFile,ret :=  client.Read(fd, 100)
		if(ret==-1){
			log.Printf("Test: Read Failed")
		}
		printChunk(readFromFile)
		log.Printf("Test: done\n")
	}
}

func printChunk ( toPrint vector.Vector ){
	fmt.Printf("Test: Chunk Read: ")
	fmt.Printf("\n")
	for i := 0; i < toPrint.Len(); i++{
		fmt.Printf("%c", toPrint.At(i).(byte))
	}
	fmt.Printf("\n")
}


