package main

import (
//	"../include/sfs" 
	"./chunk"
	"http"
	"rpc"
//	"os"
	"log"
	"net"
	"fmt"
//	"strconv"
)


func main() {
	chunkServ := new(chunk.Server)
	chunk.Init()
	rpc.Register(chunkServ)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1337")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	fmt.Println("done")
}
