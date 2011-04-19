package main

import (
        "./proxy"
	"fmt"
	//"flag"
	"rpc"
	"net"
	"os"
	)
func main(){
	fmt.Println("Creating a proxy object")
	m := new(proxy.Proxy)
	fmt.Println("Registering the proxy object")
        rpc.Register(m)
	fmt.Println("Creating a listener")
	l, e := net.Listen("tcp", ":1338")
	if e != nil {
		//log.Fatal("listen error:", e)
		fmt.Println("I am here")
		fmt.Print(e)
		os.Exit(2)
	}
	fmt.Println("Accepting the listener")
	rpc.Accept(l)
	fmt.Println("done")
}

