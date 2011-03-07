package chunk

import (
	"../include/sfs" 
	"os"
	"rpc"
	"log"
	"net"
)

type Server int

const CHUNK_TABLE_SIZE = 1024*1024*1024 / sfs.CHUNK_SIZE

var chunkTable = map[uint64] sfs.Chunk {}

func Init(thisAddr string, serverAddress string) {

	var args sfs.PingArgs
	var ret sfs.PingReturn 

	addr, err := net.ResolveTCPAddr(thisAddr+ ":1337")
	if err != nil {
		log.Fatal("ping error: ", err)
	}

	args.ChunkServer = *addr
	log.Println(args.ChunkServer)

	client, err := rpc.DialHTTP("tcp", serverAddress + ":1338")
	if err != nil {
		log.Fatal("dialing:", err)
	}

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
	log.Println("Reading from chunk ", args.ChunkID)

	ret.Data.Data = data.Data

	ret.Status = 0
	return nil	
}

func (t *Server) Write(args *sfs.WriteArgs, ret *sfs.WriteReturn) os.Error {

	data,present := chunkTable[args.ChunkID]
	if !present{
		//ret.Status = -1
	}

	log.Println("Writing to chunk ", args.ChunkID)

	data.Data = args.Data.Data
	chunkTable[args.ChunkID] = data

	return nil	
}
