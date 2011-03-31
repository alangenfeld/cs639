package main

import (
	"./client"
	"log"
	"bufio"
	"os"
	"fmt"
)

func main(){
	if len(os.Args) < 1 {
		log.Printf("Error: Run as \"client [masters ip]\"");
	}else {
		var fd int
		var data []byte
		var fromIn string
		log.Printf("Test: Enter text to write to file:\n")
		in := bufio.NewReader(os.Stdin)
		fromIn, _ = in.ReadString('\n')
		data = make([]byte,len(fromIn))
		for i := 0; i < len(fromIn); i++{
			data[i] = fromIn[i]
		}
		log.Printf("Test: creating hello\n");
		if len(os.Args) != 2 {
			client.Initialize("127.0.0.1")
		} else{
			client.Initialize(os.Args[1])
		}
		fd = client.Open("hello", 1)
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

func printChunk ( toPrint []byte ){
	fmt.Printf("Test: Chunk Read: ")
	log.Printf("Test: Chunk Read: ")
	fmt.Printf("\n")
	log.Printf("\n")
	for i := 0; i < len(toPrint); i++{
		fmt.Printf("%c", toPrint[i])
		log.Printf("%c %d", toPrint[i], i)
	}
	fmt.Printf("\n")
}


