package main

import (
	"./master"
	"fmt"
	"flag"
	"rpc"
	"http"
	"net"
	"log"
)

func main(){
	m := new(master.Master)
	
	flag.Parse()
	
	rpc.Register(m)

	l, e := net.Listen("tcp", ":1338")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	rpc.Accept(l)
	fmt.Println("done")
	
	
}
