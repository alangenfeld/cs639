package main

import (
	"http"
	"rpc"
	"os"
	"log"
	"net"
	"fmt"
)

type Args struct {
	ChunkID string
	BlockOffset int
}

type Block struct {
	Data [64 * 1024]byte
}

type Chunk struct {
	Blocks [10]Block
	Checksums [10]uint32
}

type BFS int

const BLOCK_SIZE int = 1024 * 64
const BLOCKS_PER_CHUNK int = 1024

func (t *BFS) Read(args *Args, block *Block) os.Error {

	file, err := os.Open(args.ChunkID, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("os.Open error: ", err)
	}

//	_, err = file.ReadAt(block.Data[:], (int64)(args.BlockOffset*BLOCK_SIZE))
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
