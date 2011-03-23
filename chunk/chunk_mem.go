package chunk

import (
	"../include/sfs" 
	"os"
	"rpc"
	"log"
	"net"
	"fmt"
	"container/vector"
	"syscall"
	"time"
)

type Server int

const CHUNK_TABLE_SIZE = 1024*1024*1024 / sfs.CHUNK_SIZE

var chunkTable = map[uint64] sfs.Chunk {}
var capacity uint64
var addedChunks vector.Vector
var chunkServerID uint64

func Init(masterAddress string) {

	var args sfs.ChunkBirthArgs
	var ret sfs.ChunkBirthReturn 

	args.Capacity = 5
	host,_ := os.Hostname()
	_,iparray,_ := net.LookupHost(host)
	tcpAddr,_ := net.ResolveTCPAddr(iparray[0] + ":1337")
	args.ChunkServerIP = *tcpAddr
	log.Println(args.ChunkServerIP)

	master, err := rpc.Dial("tcp", masterAddress + ":1338")
	if err != nil {
		log.Fatal("chunk dial error:", err)
	}

	err = master.Call("Master.BirthChunk", &args, &ret)
	if err != nil {
		log.Fatal("chunk call error: ", err)
	}
	chunkServerID = ret.ChunkServerID
}

func (t *Server) Read(args *sfs.ReadArgs, ret *sfs.ReadReturn) os.Error {
	data,present := chunkTable[args.ChunkID]
	if !present{
		ret.Status = -1
		return nil
	}
	log.Println("chunk: Reading from chunk ", args.ChunkID)

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

	log.Println("chunk: Writing to chunk ", args.ChunkID)

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

	host,_ := os.Hostname()
	_,iparray,_ := net.LookupHost(host)
	tcpAddr,_ := net.ResolveTCPAddr(iparray[0] + ":1337")
	args.ChunkServerIP = *tcpAddr
	args.ChunkServerID = chunkServerID 

	for {
		var info syscall.Sysinfo_t
		syscall.Sysinfo(&info)
		var mem_usage float32
		mem_usage = float32 (info.Freeram) / float32 (info.Totalram)
		if mem_usage > .8 { 
			capacity --
			//TODO: This is super sketchy and needs to be fixed, free up current blocks, etc
		}
		args.Capacity = capacity
		args.AddedChunks = addedChunks
		err = master.Call("Master.BeatHeart", &args, &ret)
		if err != nil {
			log.Fatal("chunk: heartbeat error: ", err)
		}
		addedChunks.Resize(0, 0)
		time.Sleep(sfs.HEARTBEAT_WAIT)		
	}
	return
}

func (t *Server) ReplicateChunk(args *sfs.ReplicateChunkArgs, ret *sfs.ReplicateChunkReturn) os.Error {

	if args.Servers.At(0) == nil {
		log.Printf("chunk: replication call: nil address.")
	}
	var replicationHostAddr net.TCPAddr = args.Servers.At(0).(net.TCPAddr)
	
	str := fmt.Sprintf("%s:%d", replicationHostAddr.IP, replicationHostAddr.Port)

replicationHost, err := rpc.Dial("tcp", str)
	if err != nil {
		log.Fatal("chunk: replication call:", err)
	}

	var readArgs sfs.ReadArgs
	var readRet sfs.ReadReturn
	readArgs.ChunkID = args.ChunkID

	log.Printf("replication request for site %s and chunk %d\n",
		replicationHostAddr.IP.String(),args.ChunkID);

	err = replicationHost.Call("Server.Write", &readArgs, &readRet)
	if err != nil {
		log.Fatal("chunk: replication call:", err)
	}

	chunkTable[args.ChunkID] = readRet.Data

	return nil
}
