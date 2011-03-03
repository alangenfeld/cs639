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
	
	master.AddServer(*host)
	
	rpc.Register(m)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1338")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	fmt.Println("done")
}
