package chunk

import (
	"../include/sfs" 
	"os"
	"fmt"
	"rpc"
	"log"
	"net"
)

type Server int

const CHUNK_TABLE_SIZE = 1024*1024*1024 / sfs.CHUNK_SIZE

var chunkTable = map[uint64] sfs.Chunk {}

func Init(serverAddress string) {

	client, err := rpc.DialHTTP("tcp", serverAddress + ":1338")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var addr net.TCPConn
	
	var args sfs.PingArgs
	var ret sfs.PingReturn 
	args.ChunkServer = (addr.LocalAddr())

	err = client.Call("Master.ReadChunkPing", &args, &ret)
	if err != nil {
		log.Fatal("ping error: ", err)
	}
}

func (t *Server) Read(args *sfs.ReadArgs, ret *sfs.ReadReturn) os.Error {
	data,present := chunkTable[args.ChunkID]
	if !present{
		ret.Status = -1
		return nil
	}
	fmt.Println("Reading from chunk ", args.ChunkID)

	ret.Data.Data = data.Data

	ret.Status = 0
	return nil	
}

func (t *Server) Write(args *sfs.WriteArgs, ret *sfs.WriteReturn) os.Error {

	data,present := chunkTable[args.ChunkID]
	if !present{
		//ret.Status = -1
	}

	fmt.Println("Writing to chunk ", args.ChunkID)

	data.Data = args.Data.Data
	chunkTable[args.ChunkID] = data

	return nil	
}
