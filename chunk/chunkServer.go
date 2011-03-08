package main

import (
//	"../include/sfs" 
	"./chunk"
//	"http"
	"rpc"
	"os"
	"log"
	"net"
//	"fmt"
	"flag"
//	"strconv"
)


func main() {

	thisAddr,e := os.Hostname()
	if e != nil {
		log.Fatal("chunk error:", e)
	}

	masterAddress := flag.Arg(0)

	chunkServ := new(chunk.Server)
	chunk.Init(thisAddr, masterAddress)
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
