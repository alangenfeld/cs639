package main

import (
	"rpc"
	"net"
	"./client"
	"log"
	"os"
	"../include/sfs"
)

type Master  int

func (m *Master) ReadOpen (or *sfs.OpenArgs, reply *sfs.OpenReturn) os.Error {
	log.Printf("Master: name: \n");
	reply.ServerLocation.Port = 1337;
	reply.ServerLocation.IP = net.IPv4(127,0,0,1);  //(net.IP) ((127).(0).(0).(1))
	reply.Chunk = 22;
	reply.Size =12345;
	reply.New = true;
	return nil
}


type Server  int

func (m *Server) Read (or *sfs.ReadArgs, reply *sfs.ReadReturn) os.Error {
/*	log.Printf("Master: name: \n");
	reply.ServerLocation.Port = 1337;
	reply.ServerLocation.IP = net.IPv4(127,0,0,1);  //(net.IP) ((127).(0).(0).(1))
	reply.Chunk = 22;
	reply.Size =12345;
	reply.New = true;
*/
	return nil
}

func main(){


	if len(os.Args) < 2 {
		log.Printf("Error: Run as \"client c [masters ip]\"");
	}

	switch os.Args[1]{
	case "s":
		ctm := new(Master)
		rpc.Register(ctm)
		l,e:=net.Listen("tcp", ":1330");
		log.Printf("connected\n");
		if e!=nil {
			log.Printf("listen error\n"+ e.String());
			os.Exit(1)
		}

		log.Printf("About To Serve\n");
		rpc.Accept(l);

	case "c":
		var fd int
		if len(os.Args) != 3 {
			//log.Printf("Test: correct usage is: client c masterMachine\n");
			fd = client.Open("hello", true, "127.0.0.1");
		} else{
			fd = client.Open("hello", true, os.Args[2]);
		}
		log.Printf("Test: creating hello\n");
		log.Printf("Test: writing hello\n");
		var data [32]byte
		st :=  "stealthy... ";
		for i := 0; i < len(st) ; i++{
			data[i] = st[i]
		}
		_ =  client.Write(fd,  data );
		log.Printf("Test: reading hello\n")
		_,_ =  client.Read(fd)
		log.Printf("Test: done\n")	
	case "r":
		log.Printf("Test: calling read\n");	
/*
		client.Read(os.Args[2]);
*/
		log.Printf("done\n");	

	default:
	}
}
