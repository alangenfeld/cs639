package main

import (
	"./client"
	"log"
//	"bufio"
	"io/ioutil"
	"os"
//	"fmt"
)

func main(){
	if len(os.Args) < 1 {
		log.Printf("Error: Run as \"client [masters ip]\"");
	}else {
/*
		log.Printf("Test: Enter text to write to file:\n")
		var data []byte
		var fromIn string
		in := bufio.NewReader(os.Stdin)
		fromIn, _ = in.ReadString('\n')
		data = make([]byte,len(fromIn))
		for i := 0; i < len(fromIn); i++{
			data[i] = fromIn[i]
		}
*/
		log.Printf("Test: creating hello\n");
		var master string
		if len(os.Args) != 2 {
			master = "127.0.0.1"
		} else{
			master = os.Args[1]
		}
		client.Initialize(master)

		fd := client.Open("hello", client.O_CREATE | client.O_RDWR)

		log.Printf("Test: writing hello from file\n");
		_ = WriteFromFile("client/test.in", fd, master)
		client.Seek(fd, 0, client.SEEK_SET)

		log.Printf("Test: reading hello\n")
		readFromFile,ret :=  client.Read(fd, 80)

		_ = WriteFromFile("client/test2.in", fd, master)
		client.Seek(fd, 80, client.SEEK_SET)
		readFromFile2,ret :=  client.Read(fd, 102)

		if(ret==-1){
			log.Printf("Test: Read Failed")
		}
		printChunk(readFromFile)
		printChunk(readFromFile2)
		log.Printf("Test: done\n")
	}
}

func printChunk ( toPrint []byte ){
	log.Printf("Test: ******************Chunk Read: ")
	var st string
	for i := 0; i < len(toPrint); i++{
		st += string(toPrint[i])
	}
	log.Printf(st)
	log.Printf("***************** END CHUNK READ")
}

func WriteFromFile( fileNameLocal string, fd int, master string)( int){
	f, err := ioutil.ReadFile(fileNameLocal)
	if err != nil {
		log.Fatal( "Test: error reading file in write from file")
	}
	log.Printf("Test: len f =  %d", len(f))
	toWrite := make([]byte, len(f))
	for j := 0 ; j<len(f) ; j++ {
		toWrite[j] = f[j]
	}
	return client.Write(fd,toWrite)
}


