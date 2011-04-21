package chunkProxy

import (
	"fmt"
	"./chunk"
	"../include/sfs"
	"os"
	"flag"
	"time"
	"rand"
)

type ChunkProxy int
var chunkServ = new(chunk.Server)
var delay int64
var delay_prob int64

func (t *ChunkProxy) ReplicateChunk(args *sfs.ReplicateChunkArgs, ret *sfs.ReplicateChunkReturn) os.Error {
	fmt.Println("Inside ChunkProxy ReplicateChunk")
	//time.Sleep(delay)
	Delay()
	return(chunkServ.ReplicateChunk(args,ret))	
}
func (t *ChunkProxy) Write(args *sfs.WriteArgs, ret *sfs.WriteReturn) os.Error {
	fmt.Println("Inside ChunkProxy Write")
	//time.Sleep(delay)
	Delay()
	return(chunkServ.Write(args,ret))
}
func (t *ChunkProxy) Read(args *sfs.ReadArgs, ret *sfs.ReadReturn) os.Error {
	fmt.Println("Inside ChunkProxy Read")
	//time.Sleep(delay)
	Delay()
	return(chunkServ.Read(args,ret))
}
func Init(masterAddress string,del int64,del_prob int64){
	flag.Parse()
	//chunkServ := new(chunk.Server)
	chunk.Init(masterAddress,true)
	delay = del
	delay_prob = del_prob
	//go chunk.SendHeartBeat(flag.Arg(0))
}
func Delay(){
        m := rand.Int63n(100)
	if m < delay_prob {
		fmt.Println("Delaying...")
		time.Sleep(delay);
	}
}

