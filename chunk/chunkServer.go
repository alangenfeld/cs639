package main

import (
//	"../include/sfs" 
	"./chunk"
//	"http"
	"rpc"
//	"os"
	"log"
	"net"
	"fmt"
	"flag"
//	"strconv"
)


func main() {

	thisAddr := flag.Arg(0)
	masterAddress := flag.Arg(1)

	chunkServ := new(chunk.Server)
	chunk.Init(thisAddr, masterAddress)
	rpc.Register(chunkServ)

	l, e := net.Listen("tcp", ":1337")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	rpc.Accept(l)
	fmt.Println("done")
}
