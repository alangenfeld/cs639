package main

import (
//	"../include/sfs" 
	"./chunk"
//	"http"
	"rpc"
//	"os"
	"log"
	"net"
//	"fmt"
	"flag"
//	"strconv"
)

var logging *bool = flag.Bool("log", false, "enables logging")

func main() {

	flag.Parse()
	masterAddress := flag.Arg(0)

	chunkServ := new(chunk.Server)
	if *logging {
		chunk.Init(masterAddress, true)
	}else{
		chunk.Init(masterAddress, false)
	}
	go chunk.SendHeartbeat(masterAddress)
	rpc.Register(chunkServ)
	
	log.Println("chunk: Server Online.")

	l, e := net.Listen("tcp", ":1337")
	if e != nil {
		log.Fatal("chunk error:", e)
	}
	rpc.Accept(l)
	log.Println("chunk: done.")
}
