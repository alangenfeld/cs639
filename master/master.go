package master

import (
	"net"
	"./trie"
	"log"
	"os"
	"fmt"
	"time"
	"container/vector"
	"container/heap"
	"../include/sfs"
	"rpc"
	"rand"
	"math"
	"os/signal"
)

var t *trie.Trie
var nextChunk uint64 = 1
var nextChunkServerID uint64 = 0
var serverIndex uint64 = 0
var sHeap *serverHeap
//var sMap map[net.TCPAddr](*server)
var servers map[uint64](*server)
var addrToServerMap map[string](*server)
var chunks map[uint64](*chunk)

var heartbeatMonitors map[uint64](chan int64)

type inode struct {
	name        string
	permissions uint64
	size        uint64
	lock        bool
	chunks      *vector.Vector
}

type chunk struct {
	chunkID uint64
	size    uint64
	servers *vector.Vector
}

type Master int

func (m *Master) ReadOpen(args *sfs.OpenArgs, info *sfs.OpenReturn) os.Error {

	if sHeap.vec.Len() == 0 {
		err := os.NewError("No chunk servers!")
		return err
	}
	file, newFile, err := OpenFile(args.Name, args.NewFile)
	
	if file == nil {
		return err
	}
	
	if newFile && !args.Lock {
		file.lock = false
	} else if file.lock && args.Lock {
		err := os.NewError("Cannot get lock")
		return err
	} else if args.Lock && !file.lock {
		file.lock = true
	}

	info.New = newFile
	info.Size = file.size

	info.Chunk = make([]sfs.ChunkInfo, file.chunks.Len())

	return err
}

func (m *Master) MapChunkToFile(args *sfs.MapChunkToFileArgs, ret *sfs.MapChunkToFileReturn) os.Error {
	file, ok := QueryFile(args.Name)

	if !ok {
		log.Printf("master: MapChunkToFile: File %s does not exist\n", args.Name)
		return os.NewError("File does not exist! Herp derp")
	}

	log.Printf("master: MapChunkToFile: ChunkID: %d  Offset %d  nservers: %d\n", args.Chunk.ChunkID, args.Offset, len(args.Chunk.Servers))
	var newChunk chunk

	newChunk.chunkID = args.Chunk.ChunkID
	newChunk.size = args.Chunk.Size
	newChunk.servers = new(vector.Vector)
	for i := 0; i < len(args.Chunk.Servers); i++ {
		newChunk.servers.Push(addrToServerMap[args.Chunk.Servers[i].String()])
	}

	_, err := file.MapChunk(args.Offset, &newChunk)

	if err != nil {
		return os.NewError("Could not add chunk! Ruh roh")
	}

	return nil
}

func (m *Master) GetNewChunk(args *sfs.GetNewChunkArgs, ret *sfs.GetNewChunkReturn) os.Error {
	ret.Info.ChunkID = nextChunk

	nextChunk++

	ret.Info.Servers = make([]net.TCPAddr, sfs.NREPLICAS)
	for i := 0; i < sfs.NREPLICAS && i < sHeap.vec.Len(); i++ {
		ret.Info.Servers[i] = sHeap.vec.At(i).(*server).addr
	}

	return nil
}

func (m *Master) ReportWrite(args *sfs.ReportWriteArgs, ret *sfs.ReportWriteReturn) os.Error {

	return nil
}

func (m *Master) ReadDir(args *sfs.ReadDirArgs, ret *sfs.ReadDirReturn) os.Error {
	var files map[string]interface{}
	var dirs *vector.StringVector
	var err os.Error

	if args.Prefix[len(args.Prefix)-1] == byte('/') {
		dirs, files, err = t.ReadDir(args.Prefix[0:len(args.Prefix)-2])
	} else {
		dirs, files, err = t.ReadDir(args.Prefix)
	}
	
	if err != nil {
		return err
	}
	
	log.Printf("ReadDir: prefix %s\n", args.Prefix)
	
	cnt := dirs.Len()
	retSlice := make([]string, cnt + len(files))

	var i int
	for i = 0; i < cnt; i++ {
		retSlice[i] = dirs.At(i) + "/"
	}
	
	for k, _ := range files {
		retSlice[i] = k
		i++
	}
	
	ret.FileNames = retSlice
	
	log.Printf("ReadDir: retSlice -- %+v\n", retSlice)

	return nil
}

func (m *Master) MakeDir(args *sfs.MakeDirArgs, ret *sfs.MakeDirReturn) os.Error {
	err := t.AddDir(args.DirName)
	
	return err
}

func (m *Master) RemoveFile(args *sfs.RemoveArgs, result *sfs.RemoveReturn) os.Error {
	result.Success = true
	name := args.Name

	i, exists := QueryFile(name)
	if !exists {
		log.Printf("RemoveFile: file %s does not exist\n", name)
		err := os.NewError("You are trying to delete a file that doesn't exist.")
		result.Success = false
		return err
	} else {
		for j := 0; j < i.chunks.Len(); j++ {
			chunks[i.chunks.At(j).(*chunk).chunkID] = nil, false
		}
		empty := t.Remove(name)

		if empty {
			log.Printf("There are no more files in the trie\n")
		}
	}
	return nil
}

func (m *Master) BirthChunk(args *sfs.ChunkBirthArgs, info *sfs.ChunkBirthReturn) os.Error {
	s := AddServer(args.ChunkServerIP, args.Capacity)
	go s.monitorServerBeats(heartbeatMonitors[s.id])
	
	if args.ChunkIDs != nil {
		for _, id := range args.ChunkIDs {
			c, ok := chunks[id]
			
			if ok {
				AssociateChunkAndServer(c, s)
			}
		}
	}	

	info.ChunkServerID = s.id

	log.Println("Birthed a Chunk Server!\n")

	return nil
}

func (m *Master) ReleaseLock(args *sfs.LockReleaseArgs, ret *sfs.LockReleaseReturn) os.Error {
	file, exists := QueryFile(args.Name)
	if !exists {
		ret.Status = -1
		return os.NewError("File does not exist")
	}
	file.lock = false
	ret.Status = -1
	return nil
}

func (m *Master) DeleteFile(args *sfs.DeleteArgs, ret *sfs.DeleteReturn) os.Error {
	log.Printf("DeleteFile: args -- %+v\n", args)
	err := DeleteFile(args.Name)
	
	ret.Status = (err == nil)
	 
	return err
}

func (m *Master) BeatHeart(args *sfs.HeartbeatArgs, info *sfs.HeartbeatReturn) os.Error {
	str := fmt.Sprintf("%s:%d", args.ChunkServerIP.IP.String(), args.ChunkServerIP.Port)
	log.Printf("BeatHeart: %s's HEART IS BEATING\n", str)

	//find the server who's heart is beating
	server, servOK := servers[args.ChunkServerID]
	if servOK == false {
		log.Printf("BeatHeart: Server (%s) not in server map; telling it to rebirth itself a la Madonna\n", str)
		info.Accepted = false
		return nil
	}
	
	info.Accepted = true

	//if somethings changed, update the server, heapify
	if server.capacity != args.Capacity || args.AddedChunks != nil {
		server.capacity = args.Capacity
		//make sure added chunks are valid, add them
		chunkRange := len(args.AddedChunks)
		for cnt := 0; cnt < chunkRange; cnt++ {
			//chunk , chunkOK := chunks[args.AddedChunks.At(cnt).(*chunk).chunkID]
			chunk, chunkOK := chunks[args.AddedChunks[0].ChunkID]
			//log.Printf("Herp dDerp %d\n", args.AddedChunks[0].ChunkID)
			if chunkOK == true {
				server.chunks.Push(args.AddedChunks[cnt])
				chunk.servers.Push(server)
			} /*else{
				log.Printf("BeatHeart: Error chunk %s does not exist\n",
				chunks[args.AddedChunks[0].ChunkID)
			}*/
		}
	}
	
	info.ChunksToRemove = server.evictedChunks
	
	server.evictedChunks = new(vector.Vector)

	heartbeatMonitors[args.ChunkServerID] <- time.Nanoseconds()

	return nil
}

func (s *server) monitorServerBeats(beats chan int64) int {
	for {
		t := time.NewTicker(sfs.HEARTBEAT_WAIT * 2)
		defer t.Stop()

		select {
		case <-beats:
			continue
		case <-t.C:
			RemoveServer(s)
			return -1
		}
	}
	return 0
}

func AddServer(servAddr net.TCPAddr, capacity uint64) *server {
	str := fmt.Sprintf("%s:%d", servAddr.IP.String(), servAddr.Port)
	log.Printf("AddServer: adding %s\n", str)

	s := new(server)

	s.id = nextChunkServerID
	nextChunkServerID += 1
	s.addr = servAddr
	s.capacity = capacity
	s.chunks = new(vector.Vector)
	s.evictedChunks = new(vector.Vector)

	heartbeatMonitors[s.id] = make(chan int64)
	heap.Push(sHeap, s)
	servers[s.id] = s
	addrToServerMap[servAddr.String()] = s

	return s
}

func RemoveServer(serv *server) os.Error {

	//Remove the Server
	sHeap.Remove(serv)

	for i := 0; i < sHeap.vec.Len(); i++ {
		if sHeap.vec.At(i).(*server).id == serv.id {
			log.Printf("master: RemoveServer: remove didn't actually remove server %d! Busto Rhymes\n", serv.id)
			return nil
		}
	}

	servers[serv.id] = &server{}, false
	addrToServerMap[serv.addr.String()] = &server{}, false

	str1 := fmt.Sprintf("removing server %s:%d", serv.addr.IP.String(), serv.addr.Port)
	
	network_size := float64(sHeap.vec.Len())
	if network_size <= 0 { 
		return nil
	}
	//chunk_size := serv.chunks.Len()
	rep_factor := .9

	//map the size of the network to some fraction of that size to replicate to
	if network_size < 10 {
		rep_factor = .5
	} else {
		rep_factor = .3
	}

	rep_size := int(math.Floor(rep_factor * network_size))

	//for each chunk in the server, make a replication call.
	for cnt := 0; cnt < serv.chunks.Len(); cnt++ {

		index := rand.Intn(rep_size)

		otherserver := sHeap.vec.At(index).(*server)

		str := fmt.Sprintf("%s:%d", otherserver.addr.IP.String(), otherserver.addr.Port)

		log.Printf("master: RemoveServer: dialing %s\n", str)

		client, err := rpc.Dial("tcp", str)

		if err != nil {
			log.Printf("master: RemoveServer: unable to dial %s\n", str)
		} else {
			log.Printf("master: RemoveServer: dial %s succeeded\n", str)
		}

		chunk := serv.chunks.At(cnt).(*chunk)

		//populate chunk location vector
		chunklist := make([]net.TCPAddr, chunk.servers.Len())
		for cnt1 := 0; cnt1 < chunk.servers.Len(); cnt1++ {
			chunklist[cnt1] = chunk.servers.At(cnt1).(*server).addr
		}

		//send rpc call off
		args := &sfs.ReplicateChunkArgs{chunk.chunkID, chunklist}
		reply := new(sfs.ReplicateChunkReturn)
		log.Printf("master: RemoveServer: issuing replication req to %s\n", str)
		err = client.Call("Server.ReplicateChunk", args, reply)
		if err != nil {
			log.Printf("master: RemoveServer: unable to call %s\n", str)
		}
		log.Printf("%s", reply)
	}

	log.Printf("RemoveServer: removing %s\n", str1)
	return nil
}

func OpenFile(name string, create bool) (i *inode, newFile bool, err os.Error) {
	err = nil

	i, exists := QueryFile(name)

	if !exists && create {
		log.Printf("OpenFile: file %s does not exist\n", name)
		i, err = AddFile(name)
	} else if !exists && !create {
		return nil, true, os.NewError("file doesn't exist")
	}

	newFile = !exists

	return i, newFile, err
}

func AddFile(name string) (i *inode, err os.Error) {
	i = new(inode)

	log.Printf("AddFile: nextChunk %d, len(servers) %d\n", nextChunk, sHeap.Len())

	i.size = 0
	//i.addr = *(servers.At(int(nextChunk) % servers.Len()).(*net.TCPAddr))
	//i.addr = servers[0]

	i.chunks = new(vector.Vector)

	//i.AddChunk()

	err = t.AddFile(name, i) // trie insert
	
	if err != nil {
		return nil, err
	}
	
	log.Printf("AddFile: %d nodes in trie\n", t.Size())

	return i, nil
}

func QueryFile(name string) (i *inode, fileExists bool) {
	inter, exists := t.QueryFile(name)

	if !exists {
		log.Printf("QueryFile: file %s does not exist\n", name)
		return nil, exists
	}

	log.Printf("QueryFile: %d nodes in trie\n", t.Size())

	return inter.(*inode), exists
}

func DeleteFile(name string) (err os.Error) {
	inode, exists := QueryFile(name)
	if exists {
		ok := t.Remove(name)
		
		if !ok {
			log.Printf("Delete: file %s does not exist\n", name)
			return os.NewError("file does not exist")
		}

		cnt1 := inode.chunks.Len()
		//for each chunk in the server, make an unmap call.
		for i := 0; i < cnt1; i++ {
			chunk := inode.chunks.At(i).(*chunk)

			chunk.unmapChunk()
		}
	}

	log.Printf("DeleteFile: %d nodes in trie\n", t.Size())

	return nil
}

func (i *inode) AppendChunk() (chunkID uint64, err os.Error) {
	//var serv *server = heap.Pop(sHeap).(*server)
	thisChunk := new(chunk)
	thisChunk.chunkID = nextChunk
	nextChunk += 1

	thisChunk.servers = new(vector.Vector)

	//thisChunk.servers.Push(serv)
	//serv.chunks.Push(thisChunk)
	i.chunks.Push(thisChunk)

	//heap.Push(sHeap, serv)

	chunks[thisChunk.chunkID] = thisChunk

	return thisChunk.chunkID, nil
}

func (i *inode) AddChunk() (chunkID uint64, err os.Error) {
	//var serv *server = heap.Pop(sHeap).(*server)
	thisChunk := new(chunk)
	thisChunk.chunkID = nextChunk
	nextChunk += 1

	thisChunk.servers = new(vector.Vector)

	/*
		for i := 0; i < sfs.NREPLICAS; i++ {
			thisChunk.servers.Push(sHeap.vec.At(i).(*server))
		}*/

	//thisChunk.servers.Push(serv)
	//serv.chunks.Push(thisChunk)
	i.chunks.Push(thisChunk)

	//heap.Push(sHeap, serv)

	chunks[thisChunk.chunkID] = thisChunk

	return thisChunk.chunkID, nil
}

func (i *inode) MapChunk(offset int, newChunk *chunk) (chunkID uint64, err os.Error) {
	var oldID uint64

	if offset < i.chunks.Len() {
		oldID = i.chunks.At(offset).(*chunk).chunkID
		
		c, ok := chunks[oldID]
		
		if ok {
			c.unmapChunk()
		}
		
		i.chunks.Set(offset, newChunk)
	} else if offset == i.chunks.Len() {
		i.chunks.Push(newChunk)
	} else {
		return 0, os.NewError("Fucking A.")
	}

	for j := 0; j < newChunk.servers.Len(); j++ {
		s := newChunk.servers.At(j).(*server)
		
		//if the server dies after right replication but before dying..
		if s != nil {
			s.chunks.Push(newChunk)
		}else {
			newChunk.servers.Delete(j)
		}
	}

	return newChunk.chunkID, nil
}

func (c *chunk) unmapChunk() (err os.Error){
	cnt := c.servers.Len()	
	for j := 0; j < cnt; j++ {
		c.servers.At(j).(*server).evictedChunks.Push(c.chunkID)
		
		cnt2 := c.servers.At(j).(*server).chunks.Len()
		for k := 0; k < cnt2; k++ {
			if c.servers.At(j).(*server).chunks.At(k).(*chunk) == c {
				c.servers.At(j).(*server).chunks.Delete(k)
			}
		}
	}
	
	//chunks[c.chunkID] = &chunk{}, false
	
	return nil
}

func (c *chunk) AssociateServer(s *server) os.Error {
	c.servers.Push(s)
	return nil
} 

func (s *server) AssociateChunk(c *chunk) os.Error {
	s.chunks.Push(c)
	return nil
}

func AssociateChunkAndServer(c *chunk, s *server) os.Error {
	err := s.AssociateChunk(c)
	
	if err != nil {
		return err
	}
	
	err = c.AssociateServer(s)
	
	if err != nil {
		return err
	}
	
	return nil
}


func FindMissingChunkReplicas() (ret uint64) {
	for cID, _ := range chunks {
		if chunks[cID].servers.Len() < sfs.NREPLICAS {
			ret += 1
			replicateChunk(cID)
		}
	}

	return ret
}

func replicateChunk(cID uint64) (err os.Error) {
	//var m Master

	//m.WriteReplicationReq(chunks[cID])

	return nil
}

func sigHandler() {
	for {
		sig := <- signal.Incoming
		
		log.Printf("Signal received: %d!\n", sig)
		
		if sig.String() == "SIGTERM: termination" || sig.String() == "SIGINT: interrupt" {
			os.Exit(1337)
			DumpTrie()
		}

		for s := range servers {
			fmt.Printf("%+v\n", s)
		}
	}
}

func DumpTrie(){
	dump := t.Members()
	
	cnt := dump.Len()
	for i := 0; i < cnt; i++ {
		log.Printf("dumpTrie: %d: %s\n", i, dump.At(i))
	}
	log.Printf("dumpTrie: DOT file follows\n******\n%s\n******\n", t.GetDotString())
}

func init() {
	t = trie.NewTrie()
	sHeap = new(serverHeap)
	sHeap.vec = new(vector.Vector)
	chunks = make(map[uint64](*chunk))
	servers = make(map[uint64](*server))
	addrToServerMap = make(map[string](*server))
	heartbeatMonitors = make(map[uint64](chan int64))
	//	sMap = make(map[net.TCPAddr](*server))
	heap.Init(sHeap)
	sHeap.serverChan = make(chan *heapCommand)
	go sHeap.Handler()
	go sigHandler()

	//missingCh := make(chan uint64)

	//go FindMissingChunkReplicas(missingCh)
	//go QueueReplication(missingCh)
}
