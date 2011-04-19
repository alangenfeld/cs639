package chunkProxy

import (
	"fmt"
	"./chunk"
	"../include/sfs"
	"os"
	"flag"
	"time"
)

type ChunkProxy int
var chunkServ = new(chunk.Server)
var delay int64

func (t *ChunkProxy) ReplicateChunk(args *sfs.ReplicateChunkArgs, ret *sfs.ReplicateChunkReturn) os.Error {
	fmt.Println("Inside ChunkProxy ReplicateChunk")
	time.Sleep(delay)
	return(chunkServ.ReplicateChunk(args,ret))	
}
func (t *ChunkProxy) Write(args *sfs.WriteArgs, ret *sfs.WriteReturn) os.Error {
	fmt.Println("Inside ChunkProxy Write")
	time.Sleep(delay)
	return(chunkServ.Write(args,ret))
}
func (t *ChunkProxy) Read(args *sfs.ReadArgs, ret *sfs.ReadReturn) os.Error {
	fmt.Println("Inside ChunkProxy Read")
	time.Sleep(delay)
	return(chunkServ.Read(args,ret))
}
func Init(masterAddress string){
	flag.Parse()
	//chunkServ := new(chunk.Server)
	chunk.Init(masterAddress,true)
	delay = 000000
	//go chunk.SendHeartBeat(flag.Arg(0))
}

