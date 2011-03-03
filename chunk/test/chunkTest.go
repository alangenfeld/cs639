package main

import (
	"../../include/sfs"
	"fmt"
	"rpc"
	"log"
	"flag"
//	"time"
//	"os"
 )

func main() {

	serverAddress := flag.Arg(0)
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1337")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var ret sfs.WriteReturn
	var args sfs.WriteArgs
	args.ChunkID = 12
	args.Data.Data[0] = 'a'

	err = client.Call("Server.Write", &args, &ret)
	if err != nil {
		log.Fatal("chunk server error: ", err)
	}

	var rdRet sfs.ReadReturn
	var rdArgs sfs.ReadArgs
	
	rdArgs.ChunkID = 12
	
	err = client.Call("Server.Read", &rdArgs, &rdRet)
	if err != nil {
		log.Fatal("chunk server error: ", err)
	}

	if rdRet.Status == -1 {
		log.Fatal("FAILURE of read call")
	}

	if args.Data.Data[0] != rdRet.Data.Data[0] {
		fmt.Println(args.Data.Data[0], " VS ", rdRet.Data.Data[0])
		log.Fatal("Data doesnt match!!!")
	}


}
