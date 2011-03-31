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


func main() {

	masterAddress := flag.Arg(0)

	chunkServ := new(chunk.Server)
	chunk.Init(masterAddress)
	go chunk.SendHeartbeat(masterAddress)
	//go chunk.LogStats()
	rpc.Register(chunkServ)
	
	log.Println("chunk: Server Online.")

	l, e := net.Listen("tcp", ":1337")
	if e != nil {
		log.Fatal("chunk error:", e)
	}
	rpc.Accept(l)
	log.Println("chunk: done.")
}
