package main

import (
	"../include/chunk"
	"http"
	"rpc"
	"os"
	"log"
	"net"
	"fmt"
)

func (t *SFS) Read(args *ReadArgs, chunk *Chunk) os.Error {

	file, err := os.Open(args.ChunkID, os.O_RDONLY, 0666)
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
	sfs := new(SFS)
	rpc.Register(sfs)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	fmt.Println("done")
}
