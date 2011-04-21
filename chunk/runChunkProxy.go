package main

import (
//	"../include/sfs" 
	"./chunk"
	"./chunkProxy"
//	"http"
	"rpc"
	"os"
	"log"
	"net"
	"fmt"
	"flag"
	"strconv"
)

var logging *bool = flag.Bool("log", false, "enables logging")

func main() {

	flag.Parse()
	masterAddress := flag.Arg(0)
	delaystr := flag.Arg(1)
	delay_prob_str := flag.Arg(2)
	delay,e := strconv.Atoi64(delaystr)
	if e != nil{
		log.Fatal("string conversion error",e)
		fmt.Print(e)
		os.Exit(2)
	}
	delay_prob,e := strconv.Atoi64(delay_prob_str)
	if e != nil{
		fmt.Print(e)
		os.Exit(2)
	}
	chunkServ := new(chunkProxy.ChunkProxy)

	if *logging {
		chunkProxy.Init(masterAddress,delay,delay_prob)
	}else{
		chunkProxy.Init(masterAddress,delay,delay_prob)
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
