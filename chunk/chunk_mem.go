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
//	"flag"
	"../logger/logger"
	"os/signal"
)

type Server int

const CHUNK_TABLE_SIZE = 1024*1024*1024 / sfs.CHUNK_SIZE
const STATUS_CMD = "../stats.sh"
const STATUS_ARGS = ""
const STATUS_LEN = 17
const THRESHOLD = 15 //represents value out of 20

var chunkTable = map[uint64] sfs.Chunk {}
var capacity uint64
var addedChunks vector.Vector
var chunkServerID uint64
var logging bool
var requestLoad int
var loadArray []int
var loadArrayIndex int
var tcpAddr *net.TCPAddr

func Init(masterAddress string, loggingFlag bool) {

	var args sfs.ChunkBirthArgs
	var ret sfs.ChunkBirthReturn 

	requestLoad = 0
	loadArray = make([]int, 3)
	for i:=0; i < 3; i ++ {
		loadArray[i] = 0
	}
	loadArrayIndex = 0
	logging = loggingFlag

	capacity = CHUNK_TABLE_SIZE
	args.Capacity = capacity

	host,_ := os.Hostname()
	_,iparray,_ := net.LookupHost(host)
	tcpAddr,_ = net.ResolveTCPAddr(iparray[0] + ":1337")
	args.ChunkServerIP = *tcpAddr
	log.Println("Chunk: Server addr: ", args.ChunkServerIP)
	
	if logging {
		err := os.Mkdir("log", 0777);
		if err != nil {
			log.Println(err.String())
			logging = false
		}
		err = logger.Init("log/chunk-log-" + host + ".txt", "../logger/")
		if err != nil {
			log.Println(err.String())
			logging = false
		}
	}
    logger.QuickInit()

	master, err := rpc.Dial("tcp", masterAddress + ":1338")
	if master != nil {
		defer master.Close()
	}
	if err != nil {
		log.Fatal("chunk dial error:", err)
	}

	err = master.Call("Master.BirthChunk", &args, &ret)
	if err != nil {
		log.Fatal("chunk call error: ", err)
	}
	chunkServerID = ret.ChunkServerID
	go sigHandler();
}

func (t *Server) Read(args *sfs.ReadArgs, ret *sfs.ReadReturn) os.Error {
	requestLoad++
	//var id logger.TaskId
	if logging {
	//	id = logger.Start("Read")
	}
	log.Println("chunk: Reading from chunk ", args.ChunkID)

	data,present := chunkTable[args.ChunkID]
	if !present{
		ret.Status = sfs.FAIL
		log.Println("chunk: Invalid read request chunk ", args.ChunkID)
		return nil
	}


	/*if args.Nice == sfs.NICE && ServerBusy() {
		log.Println("chunk: BUSY")
		ret.Status = sfs.BUSY
		return nil
	}*/
	
	ret.Data.Data = data.Data
	ret.Status = sfs.SUCCESS
	log.Println("chunk: Read success")
	/*if logging {
		errString := logger.End(id, false)
		if errString != "" {
			logging = false
		}
	}*/
	return nil	
}

func (t *Server) Write(args *sfs.WriteArgs, ret *sfs.WriteReturn) os.Error {
	requestLoad++
	ret.Status = sfs.FAIL
	var id logger.TaskId

	id = logger.Start("Write")
	log.Println("chunk: Writing to chunk ", args.Info.ChunkID)
	if (capacity < 1) {
		log.Println("chunk: Server Full!")
		return nil
	}
	
	data,present := chunkTable[args.Info.ChunkID]
	if !present{
		addedChunks.Push(args.Info)
		capacity --
	}

	data.Data = args.Data.Data
	chunkTable[args.Info.ChunkID] = data

	tempServ := args.Info.Servers[0]
	var inRet sfs.WriteReturn

	for {
		if len(args.Info.Servers) < 2 {
			break
		}
		args.Info.Servers = args.Info.Servers[1:len(args.Info.Servers)]
		client, err := rpc.Dial("tcp", args.Info.Servers[0].String())
		if err != nil {
			log.Println("chunk: dialing error: ", err)
			continue
		}
		log.Println("chunk: forwarding write to ", args.Info.Servers[0])
		err = client.Call("Server.Write", &args, &inRet)
	        client.Close()
		if err != nil {
			log.Println("chunk: server error: ", err)
			continue
		}
		break
	}

	ret.Info.Servers = make([]net.TCPAddr, len(inRet.Info.Servers) + 1)
	
	for i:=0;i<len(ret.Info.Servers);i++ {
		if(i < len(inRet.Info.Servers)) {
			ret.Info.Servers[i] = inRet.Info.Servers[i]		
		} else {
			ret.Info.Servers[i] = tempServ		
		}
	}

    logger.End(id, false)

	ret.Status = sfs.SUCCESS
	return nil	
}

/*func (t *Server) Get(args *sfs.PingArgs, ret *sfs.PingReturn) os.Error {
	return nil
}*/

func LogStats(){
	//first, open the daily log file
	
	//log.Println("We're alive")
	host,_ := os.Hostname()
	//current := time.Seconds()
	filename:= host
	logFile,err := os.Open(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("chunk: unable to init logging")
		log.Println("chunk fails opening file")
	}
	//now, enter the happy place of logging all our fun stuff
	//var info syscall.Sysinfo_t
	args := make([]string, 1)
	var result []byte
	result = make([]byte, STATUS_LEN)
	for {
		command, err2 := exec.Run(STATUS_CMD, args, nil, ".", exec.PassThrough, exec.Pipe, exec.PassThrough)
		if err2 != nil{
			log.Println("chunk fails in command:" + err2.String())
			log.Fatal("chunk: unable to obtain remote command")
		}
		/*length,err3 := command.Stdout.Read(result)
		if err3 != nil{
			log.Println("chunk fails read from command: " + err3.String())
		}*/
		length := 0
		var err3 os.Error
		for j :=0; j < 4; j ++ {
			if (length == 0) {
				time.Sleep(4000000000)
				length,err3 =command.Stdout.Read(result)
				if err3 != nil{
				log.Println("chunk fails read from command: " + err3.String())
				}
			}
		}
		logFile.Write(result)
		logFile.WriteString("\nNextRecord\n");
	}
	logFile.Close();		
}

func SendHeartbeat(masterAddress string){
	var args sfs.HeartbeatArgs
	var ret  sfs.HeartbeatReturn
	
	master, err := rpc.Dial("tcp", masterAddress + ":1338")
	if(master != nil){
		defer master.Close()
	}
	if err != nil {
		log.Fatal("chunk: dialing:", err)
	}

	host,_ := os.Hostname()
	_,iparray,_ := net.LookupHost(host)
	tcpAddr,_ := net.ResolveTCPAddr(iparray[0] + ":1337")
	args.ChunkServerIP = *tcpAddr
	args.ChunkServerID = chunkServerID

	for {
		var id logger.TaskId	
		if logging {
			id = logger.Start("Heart")
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
		if ret.Accepted == false {
			var bArgs sfs.ChunkBirthArgs
			var bRet sfs.ChunkBirthReturn
			host,_ := os.Hostname()
			_,iparray,_ := net.LookupHost(host)
			tcpAddr,_ := net.ResolveTCPAddr(iparray[0] + ":1337")
			bArgs.ChunkServerIP = *tcpAddr
			bArgs.ChunkIDs = make([]uint64, len(chunkTable))
			i := 0
			for k, _ := range chunkTable {
				bArgs.ChunkIDs[i] = k
				i++
			}
			log.Println("chunk: heartbeat")
			err = master.Call("Master.BirthChunk", &bArgs, &bRet)
			if err != nil {
				log.Fatal("chunk call error: ", err)
			}
		} else if ret.ChunksToRemove != nil {
			for i := 0; i < ret.ChunksToRemove.Len(); i++ {
				chunkTable[ret.ChunksToRemove.At(i).(uint64)] = 
					sfs.Chunk{}, false
				capacity++
			}
		}
		addedChunks.Resize(0, 0)
		if logging {
			errString := logger.End(id, false)
			if errString != "" {
				logging = false
			}
		}
		//addCount(requestLoad)
		requestLoad = 0
		time.Sleep(sfs.HEARTBEAT_WAIT)	
	}
	return
}

func (t *Server) ReplicateChunk(args *sfs.ReplicateChunkArgs, ret *sfs.ReplicateChunkReturn) os.Error {
	requestLoad++
	if args.Servers == nil {
		log.Println("chunk: replication call: nil address.")
		return nil
	}
	
	log.Println("chunk: replication request chunk ", args.ChunkID);
	_,present := chunkTable[args.ChunkID]
	if present {
		log.Println("chunk: already have it!");
		return nil
	}
	
	for i := 0; i < len(args.Servers); i++ {
		if(args.Servers[i].String() == tcpAddr.String()) {
			continue;
		}
		
		replicationHost, err := rpc.Dial("tcp", args.Servers[i].String())
		if(replicationHost != nil){
			defer replicationHost.Close()
		}
		if err != nil {
			log.Println("chunk: replication error ", err)
			continue
		}
		
		var readArgs sfs.ReadArgs
		var readRet sfs.ReadReturn
		readArgs.ChunkID = args.ChunkID
		log.Println("chunk: replicating from ", args.Servers[i]);
		err = replicationHost.Call("Server.Read", &readArgs, &readRet)
		if err != nil {
			log.Println("chunk: replication error ", err)
			continue
		}
		
		chunkTable[args.ChunkID] = readRet.Data
		capacity--
		break
	}
	return nil
}

func ServerBusy() bool {
	log.Println("Chunk: calculating load...")
    loggerLoad := logger.GetLoad()
    chunkLoad := getAvgReq()
    log.Println("Chunk: logger metric says: ", loggerLoad)
    log.Println("Chunk: chunk metric says: ", chunkLoad)

	index := loggerLoad * 2
//	index := requestLoad
	log.Println("Chunk: server load index", index)

	return index > THRESHOLD
}

func getAvgReq() int {
	var avg float64
	avg = 1
	for i:=0; i < 3; i ++ {
		avg += float64(loadArray[i])
	}
    avg = avg/3.0
    log.Println("avg is: ", avg)
	//use loadArrayIndex -1 for most recent reading
	var currentVal float64
	if loadArrayIndex == 0 {
		currentVal = float64(loadArray[2])
	} else {
		currentVal = float64(loadArray[loadArrayIndex -1])
	}
    log.Println("currentVal is: ", currentVal)
	if currentVal > avg {
		retVal := int(10.0 * (currentVal - avg)/avg)
        if (retVal > 10){
           return 10
        }
        return retVal
	}
	return 0
}

func addCount(currReq int) {
    log.Println("in addCount: Index is: ", loadArrayIndex)
	loadArray[loadArrayIndex] = currReq
	loadArrayIndex ++
	if loadArrayIndex >= len(loadArray) {
        log.Println("We flip when index is: ", loadArrayIndex)
		loadArrayIndex = 0
	}
}

func sigHandler() {
	for {
		sig := <- signal.Incoming
		
		log.Println("Signal received: %d!\n", sig)
		
		if sig.String() == "SIGTERM: termination" || sig.String() == "SIGINT: interrupt" {
			log.Println("Chunk Server going down ", (*tcpAddr).String())
			os.Exit(1)
		}

	}
}
