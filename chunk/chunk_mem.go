package chunk

import (
	"../include/sfs" 
	"os"
	"fmt"
)

type Server int

const CHUNK_TABLE_SIZE = 1024*1024*1024 / sfs.CHUNK_SIZE

var chunkTable = map[uint64] sfs.Chunk {}

func Init() {
	return
}

func (t *Server) Read(args *sfs.ReadArgs, ret *sfs.ReadReturn) os.Error {
	data,present := chunkTable[args.ChunkID]
	if !present{
		ret.Status = -1
		return nil
	}
	fmt.Println("Reading from chunk ", args.ChunkID)

	ret.Data.Data[0] = data.Data[0]

	return nil	
}

func (t *Server) Write(args *sfs.WriteArgs, ret *sfs.WriteReturn) os.Error {

	data,present := chunkTable[args.ChunkID]
	if !present{
		//ret.Status = -1
	}

	fmt.Println("Writing to chunk ", args.ChunkID)
	data.Data[0] = args.Data.Data[0]
	chunkTable[args.ChunkID] = data

	return nil	
}
