package chunk

import (
	"../include/sfs" 
	"os"
	"rpc"
	"log"
	"net"
	"container/vector"
	"time"
)

type Server int

const CHUNK_TABLE_SIZE = 1024*1024*1024 / sfs.CHUNK_SIZE

var chunkTable = map[uint64] sfs.Chunk {}
var capacity uint64
var addedChunks vector.Vector

func Init(thisAddr string, masterAddress string) {

	var args sfs.ChunkBirthArgs
	var ret sfs.ChunkBirthReturn 

	addr, err := net.ResolveTCPAddr(thisAddr+ ":1337")
	if err != nil {
		log.Fatal("chunk: ping error: ", err)
	}

	capacity = 5
	args.ChunkServer = *addr
	log.Println(args.ChunkServer)

	master, err := rpc.Dial("tcp", masterAddress + ":1338")
	if err != nil {
		log.Fatal("chunk: dialing:", err)
	}

	err = master.Call("Master.ReadChunkPing", &args, &ret)
	if err != nil {
		log.Fatal("chunk: ping error: ", err)
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
		addedChunks.Push(args.ChunkID)
		capacity --
	}

	log.Println("Writing to chunk ", args.ChunkID)

	data.Data = args.Data.Data
	chunkTable[args.ChunkID] = data

	return nil	
}

/*func (t *Server) Get(args *sfs.PingArgs, ret *sfs.PingReturn) os.Error {
	return nil
}*/

func SendHeartbeat(masterAddress string){
	var args sfs.HeartbeatArgs
	var ret  sfs.HeartbeatReturn
	
	master, err := rpc.Dial("tcp", masterAddress + ":1338")
	if err != nil {
		log.Fatal("chunk: dialing:", err)
	}
	

	for {
		args.Capacity = capacity
		args.AddedChunks = addedChunks
		err = master.Call("Master.ReadChunkHeartbeat", &args, &ret)
		if err != nil {
			log.Fatal("chunk: ping error: ", err)
		}
		addedChunks.Resize(0, 0)
		time.Sleep(sfs.HEARTBEAT_WAIT)		
	}
	return
}
