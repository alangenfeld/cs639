package main


import (
	"http"
	"rpc"
	"os"
	"log"
	"net"
	"fmt"
	"container/list"
	"../include/sfs"
)

type HeartbeatArgs struct {
        error uint //reserved, but can't think of anything to put here yet
}

type Status struct {
        ChunkCount uint
        ChunkIDs list.List
}

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

func (t *SFS) Heartbeat(args *HeartbeatArgs, status *Status) os.Error {
	return nil
}

func Init(status Status) {
	status.ChunkCount = 0
	status.ChunkIDs = ChunkIDs.New(List)
	return 
}
	

func main() {
	sfs := new(SFS)
	status := new(Status)
	Init(status)
	rpc.Register(sfs)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	fmt.Println("done")
}
