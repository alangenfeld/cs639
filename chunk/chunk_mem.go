package chunk

import (
	"../include/sfs" 
	"os"
	"rpc"
	"log"
	"net"
//	"fmt"
	"container/vector"
	//"syscall"
	"time"
	"exec"
)

type Server int

const CHUNK_TABLE_SIZE = 1024*1024*1024 / sfs.CHUNK_SIZE
const STATUS_CMD = ""
const STATUS_ARGS = ""

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

	for {
		if len(args.Info.Servers) < 2 {
			return nil
		}
		
		args.Info.Servers = args.Info.Servers[1:len(args.Info.Servers)]

//		var replicationHostAddr net.TCPAddr = args.Info.Servers.At(0).(net.TCPAddr)
//		str := fmt.Sprintf("%s:%d", replicationHostAddr.IP, replicationHostAddr.Port)

		client, err := rpc.Dial("tcp", args.Info.Servers[0].String())
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

func LogStats(){
	//first, open the daily log file
	host,_ := os.Hostname()
	current := time.Seconds()
	filename:= host + string(current)
	logFile, err := os.Open(filename, os.O_CREAT | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("chunk: unable to init logging")
	}
	//now, enter the happy place of logging all our fun stuff
	//var info syscall.Sysinfo_t
	var args []string
	var result []byte
	args[0] = STATUS_ARGS
	for {
		command, err2 := exec.Run(STATUS_CMD, args, nil, "", exec.DevNull, exec.Pipe, exec.MergeWithStdout)
		if err2 != nil{
			log.Fatal("chunk: unable to obtain remote command")
		}
		command.Stdout.Read(result)
		logFile.Write(result);
		logFile.WriteString(string(result));
		time.Sleep(2000000000);
	}		
}

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

	for i := 0; i < len(args.Servers); i++ {
		replicationHost, err := rpc.Dial("tcp", args.Servers[i].String())
		if err != nil {
			log.Println("chunk: replication call:", err)
			continue
		}
		
		var readArgs sfs.ReadArgs
		var readRet sfs.ReadReturn
		readArgs.ChunkIDs = args.ChunkID
		
		log.Printf("replication request for site %s and chunk %d\n",
			args.Servers[0].String(),args.ChunkID);
		
		err = replicationHost.Call("Server.Write", &readArgs, &readRet)
		if err != nil {
			log.Println("chunk: replication call:", err)
			continue
		}
		
		chunkTable[args.ChunkID] = readRet.Data
		break
	}
	return nil
}
