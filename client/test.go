package main

import (
	"rpc"
	"net"
	"./client"
	"fmt"
	"os"
	"../include/sfs"
)

type Master  int

func (m *Master) ReadOpen (or *sfs.OpenArgs, reply *sfs.OpenReturn) os.Error {
	fmt.Printf("name: \n");
	fmt.Printf("name: \n", or.Name);
	reply.ServerLocation.Port = 9876;
	reply.Chunk = 22;
	reply.Size =12345;
	reply.New = true;
	return nil
}

func main(){

	switch os.Args[1]{
	case "s":
		ctm := new(Master)
		rpc.Register(ctm)
		l,e:=net.Listen("tcp", "127.0.0.1:1338");
		fmt.Printf("connected\n");	
		if e!=nil {
			fmt.Printf("listen error\n"+ e.String());	
			os.Exit(1)
		}
		fmt.Printf("About To Serve\n");	
		rpc.Accept(l);

	case "c":
		fmt.Printf("calling hello\n");	
		_, serveLoc := client.Open("hello", true);
		client.Read(0,0, serveLoc);
		//fmt.Printf("done\n");	
	case "r":
		fmt.Printf("calling read\n");	
/*
		serveLoc:= new (net.TCPAddr);
		serveLoc.IP = "127.0.0.1" 
		serveLoc.Port = "1337" 
		client.Read(0,0, serveLoc);
*/
		fmt.Printf("done\n");	

	default:
	}


}
