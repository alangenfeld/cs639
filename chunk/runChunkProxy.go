package main

import (
//	"../include/sfs" 
	"./chunk"
	"./chunkProxy"
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

	chunkServ := new(chunkProxy.ChunkProxy)

	if *logging {
		chunkProxy.Init(masterAddress)
	}else{
		chunkProxy.Init(masterAddress)
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
