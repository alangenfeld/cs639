package main

import (
	"../include/sfs" 
	"http"
	"rpc"
	"os"
	"log"
	"net"
	"fmt"
	"strconv"
)

func (t *sfs.Handle) Read(args *sfs.ReadArgs, chunk *sfs.Chunk) os.Error {

	fileName := strconv.Uitob64(args.ChunkID, 10)
	file, err := os.Open(fileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("os.Open error: ", err)
	}

	_, err = file.Read(chunk.Data[:])
	if err != nil {
		log.Fatal("file.Read error: ", err)
	}

	file.Close()

	return nil
}

func main() {
	sfs := new(sfs.Handle)
	rpc.Register(sfs)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	fmt.Println("done")
}
