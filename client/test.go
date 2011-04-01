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
		fd := client.Open("hello", 1)
		log.Printf("Test: writing hello from file\n");
		_ = WriteFromFile("client/test.in", fd, master)
	//	_ =  client.Write(fd,  data );
		log.Printf("Test: reading hello\n")
		readFromFile,ret :=  client.Read(fd, 80)
		_ = WriteFromFile("client/test2.in", fd, master)
	//	_ =  client.Write(fd,  data );
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
//	fmt.Printf("Test: Chunk Read: ")
	log.Printf("Test: ******************Chunk Read: ")
//	fmt.Printf("\n")
	var st string
	for i := 0; i < len(toPrint); i++{
//		fmt.Printf("%c", toPrint[i])
		st += string(toPrint[i])
//		log.Printf("%c %d", toPrint[i], i)
	}
	log.Printf(st)
	log.Printf("***************** END CHUNK READ")
//	fmt.Printf(st)
//	fmt.Printf("\n")
}

func WriteFromFile( fileNameLocal string, fd int, master string)( int){
	f, err := ioutil.ReadFile(fileNameLocal)
	if err != nil {
		log.Fatal( "Test: error reading file in write from file")
	}
	log.Printf("Test: len f =  %d", len(f))
	toWrite := make([]byte, len(f))
	for j := 0 ; j<len(f) ; j++ {
		//log.Printf("j = %d", j)
		toWrite[j] = f[j]
	}
	return client.Write(fd,toWrite)
}


