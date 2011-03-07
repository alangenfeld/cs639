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

var host = flag.String("h", "localhost:1337", "host to connect to")

func main(){
	m := new(master.Master)
	
	flag.Parse()
	
	//addr, _ := net.ResolveTCPAddr(*host)
	
	//master.AddServer(*addr, 1000)
	
	rpc.Register(m)
	//rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1338")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	rpc.Accept(l)
	fmt.Println("done")
	
	//http.Serve(l, nil)
	
}
