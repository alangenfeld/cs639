package main

import (
	"net"
	"fmt"
	"flag"
	"log"
//	"time"
)


func main() {

	servName := flag.Arg(0)

	localAddr,err := net.ResolveUDPAddr(servName + ":1234")
	if err != nil {
		log.Fatal("UDP Resolve Failure: ", err)
	}

	local,err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatal("UDP Listen Failure: ", err)
	}

	var buff [1]byte
	var n int

	n,remoteAddr,err := local.ReadFromUDP(buff[:])
	if err != nil {
		log.Fatal("listen ReadFromUDP: ", err, n)
	}
	
	remote,err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatal("DialUDP : ", err)
	}
	
	//	fmt.Println("bytes read", n, "from",  remoteAddr)
	
	n,err = remote.WriteToUDP(buff[:], remoteAddr)
	if err != nil {
		log.Fatal("listen WriteFromUDP: ", err, n)
	}
	
	for ;; {
		n,remoteAddr,err := local.ReadFromUDP(buff[:])
		if err != nil {
			log.Fatal("listen ReadFromUDP: ", err, n)
		}

		remote,err := net.DialUDP("udp", nil, remoteAddr)
		if err != nil {
			log.Fatal("DialUDP : ", err)
		}
		
		//	fmt.Println("bytes read", n, "from",  remoteAddr)
		
		n,err = remote.WriteToUDP(buff[:], remoteAddr)
		if err != nil {
			log.Fatal("listen WriteFromUDP: ", err, n)
		}
	}
	fmt.Println("bytes written", n, "to", remoteAddr)
}
