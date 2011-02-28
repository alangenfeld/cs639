package main

import (
	"http"
	"../include/chunk"
	"rpc"
	"os"
	"log"
	"net"
	"fmt"
)

func (t *BFS) Read(args *Args, block *Block) os.Error {

	file, err := os.Open(args.ChunkID, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("os.Open error: ", err)
	}

	_, err = file.Read(block.Data[:])
	if err != nil {
		log.Fatal("file.Read error: ", err)
	}

	file.Close()

	return nil
}

func main() {

	bfs := new(BFS)
	rpc.Register(bfs)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	fmt.Println("done")
}
