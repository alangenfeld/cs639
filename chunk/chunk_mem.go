package chunk

import (
	"../include/sfs" 
	"os"
	"rpc"
	"log"
	"net"
//	"fmt"
	"container/vector"
	"syscall"
	"time"
)

type Server int

const CHUNK_TABLE_SIZE = 1024*1024*1024 / sfs.CHUNK_SIZE
const NUM_USR_COMMAND = "who"
const NUM_USR_ARGS = "| awk '{print $1;}'| uniq | wc -l |awk '{print $1;}'"

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
	data,present := chunkTable[args.ChunkIDs]
	if !present{
		ret.Status = -1
		return nil
	}
	log.Println("chunk: Reading from chunk ", args.ChunkIDs)

	ret.Data.Data = data.Data

	ret.Status = 0
	return nil	
}

func (t *Server) Write(args *sfs.WriteArgs, ret *sfs.WriteReturn) os.Error {

	log.Println("chunk: Writing to chunk ", args.Info.ChunkID)
	data,present := chunkTable[args.Info.ChunkID]
	if !present{
		addedChunks.Push(args.Info)
		capacity --
	}

	log.Println("chunk: Writing to chunk ", args.Info.ChunkID)

	data.Data = args.Data.Data
	chunkTable[args.Info.ChunkID] = data

	if args.Info.Servers == nil || args.Info.Servers.At(0) == nil {
		return nil
	}

	for {
		if args.Info.Servers.At(1) == nil {
			return nil
		}
		
		args.Info.Servers.Slice(1,args.Info.Servers.Len())

//		var replicationHostAddr net.TCPAddr = args.Info.Servers.At(0).(net.TCPAddr)
//		str := fmt.Sprintf("%s:%d", replicationHostAddr.IP, replicationHostAddr.Port)

		client, err := rpc.Dial("tcp", args.Info.Servers.At(0).(*net.TCPAddr).String())
		if err != nil {
			log.Printf("chunk: dialing:", err)
			continue
		}
		
		err = client.Call("Server.Write", &args, &ret)
		if err != nil {
			log.Fatal("chunk: server error: ", err)
		}
		break
	}
	return nil	
}

/*func (t *Server) Get(args *sfs.PingArgs, ret *sfs.PingReturn) os.Error {
	return nil
}*/

/*func LogStats(){
	//first, open the daily log file
	host,_ := os.Hostname()
	current := time.Seconds()
	filename := host
	logFile, err := os.Open(filename, os.O_CREAT | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("chunk: unable to init logging");
	}
	
	//now, enter the happy place of logging all our fun stuff
	var info syscall.Sysinfo_t
	var args []string
	args[0] = NUM_USR_ARGS
	for {
		syscall.Sysinfo(&info)
		p, err := os.StartProcess(NUM_USR_COMMAND, args, nil, nil, nil)
	}		
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
		if mem_usage > .8  && capacity > 0{ 
			capacity --
			//TODO: This is super sketchy and needs to be fixed, free up current blocks, etc
		}else if mem_usage < .2{
			capacity ++
		}
		args.Capacity = capacity
		addedChunkSlice := make([]sfs.ChunkInfo, addedChunks.Len())
		for i := 0; i < addedChunks.Len(); i++ {
			addedChunkSlice[i] = addedChunks.At(i).(sfs.ChunkInfo)
		}
		args.AddedChunks = addedChunkSlice
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

	if args.Servers == nil {
		log.Printf("chunk: replication call: nil address.")
		return nil
	}

//	var replicationHostAddr net.TCPAddr = args.Servers.At(0).(net.TCPAddr)
//	str := fmt.Sprintf("%s:%d", replicationHostAddr.IP, replicationHostAddr.Port)
	
	replicationHost, err := rpc.Dial("tcp", args.Servers.At(0).(*net.TCPAddr).String())
	if err != nil {
		log.Fatal("chunk: replication call:", err)
	}

	var readArgs sfs.ReadArgs
	var readRet sfs.ReadReturn
	readArgs.ChunkIDs = args.ChunkID

	log.Printf("replication request for site %s and chunk %d\n",
		args.Servers.At(0).(*net.TCPAddr).String(),args.ChunkID);

	err = replicationHost.Call("Server.Write", &readArgs, &readRet)
	if err != nil {
		log.Fatal("chunk: replication call:", err)
	}

	chunkTable[args.ChunkID] = readRet.Data

	return nil
}
