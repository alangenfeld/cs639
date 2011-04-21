package main

import (
        "./proxy"
	"fmt"
	"flag"
	"rpc"
	"net"
	"os"
	"strconv"
	)
func main(){
	fmt.Println("Creating a proxy object")
        flag.Parse()
	delaystr := flag.Arg(0)
	delayprobstr := flag.Arg(1)
	m := new(proxy.Proxy)
	delay,err := strconv.Atoi64(delaystr)
	if err != nil{
		//log.Fatal("String conversion error:",err)
		fmt.Print(err)
		os.Exit(2)
	}
	delay_prob, err := strconv.Atoi64(delayprobstr)
	if err!= nil{
		fmt.Print(err)
		os.Exit(2)
	}
	proxy.Init(delay,delay_prob)
	fmt.Println("Registering the proxy object")
        rpc.Register(m)
	fmt.Println("Creating a listener")
	l, e := net.Listen("tcp", ":1338")
	if e != nil {
		//log.Fatal("listen error:", e)
		//fmt.Println("I am here")
		fmt.Print(e)
		os.Exit(2)
	}
	fmt.Println("Accepting the listener")
	rpc.Accept(l)
	fmt.Println("done")
}

