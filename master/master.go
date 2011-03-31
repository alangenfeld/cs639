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
	chunks      *vector.Vector
}

type chunk struct {
	chunkID uint64
	servers *vector.Vector
}

type Master int

func (m *Master) ReadOpen(args *sfs.OpenArgs, info *sfs.OpenReturn) os.Error {

	if sHeap.vec.Len() == 0 {
		err := os.NewError("No chunk servers!")
		return err
	}
	file, newFile, err := OpenFile(args.Name)

	info.New = newFile
	info.Size = file.size

	info.Chunk = make([]sfs.ChunkInfo, file.chunks.Len())

	//for each chunk in the server, make a replication call.
	for i := 0; i < file.chunks.Len(); i++ {
		var chunkInfo sfs.ChunkInfo
		chunk := file.chunks.At(i).(*chunk)

		//populate chunk location vector
		chunkInfo.Servers = make([]net.TCPAddr, chunk.servers.Len())

		for j := 0; j < chunk.servers.Len(); j++ {
			chunkInfo.Servers[j] = chunk.servers.At(j).(*server).addr
		}

		chunkInfo.ChunkID = chunk.chunkID

		info.Chunk[i] = chunkInfo
	}

	return err
}

func (m *Master) MapChunkToFile(args *sfs.MapChunkToFileArgs, ret *sfs.MapChunkToFileReturn) os.Error {
	file, ok := QueryFile(args.Name)

	if !ok {
		return os.NewError("File does not exist! Herp derp")
	}

	var newChunk chunk
	
	newChunk.chunkID = args.Chunk.ChunkID
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

	info.ChunkServerID = s.id

	log.Println("Birthed a Chunk Server!\n")

	return nil
}

func (m *Master) BeatHeart(args *sfs.HeartbeatArgs, info *sfs.HeartbeatReturn) os.Error {
	str := fmt.Sprintf("%s:%d", args.ChunkServerIP.IP.String(), args.ChunkServerIP.Port)
	log.Printf("BeatHeart: %s's HEART IS BEATING\n", str)

	//find the server who's heart is beating
	server, servOK := servers[args.ChunkServerID]
	if servOK == false {
		log.Printf("BeatHeart: Error server (%s) not in server map\n", str)
	}

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

	//////////////////////////////////////////////////////////////////////
	////////////      THIS IS A PROBLEM -- NEED TO SPREAD REPLICATION REQS
	//////////////////////////////////////////////////////////////////////
	otherserver := sHeap.vec.At(0).(*server)

	str := fmt.Sprintf("%s:%d", otherserver.addr.IP.String(), otherserver.addr.Port)
	
	log.Printf("master: RemoveServer: dialing %s\n", str)

	client, err := rpc.Dial("tcp", str)
	
	if err != nil {
		log.Printf("master: RemoveServer: unable to dial %s\n", str)
	} else {
		log.Printf("master: RemoveServer: dial %s succeeded\n", str)
	}

	//for each chunk in the server, make a replication call.
	for cnt := 0; cnt < serv.chunks.Len(); cnt++ {

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

func OpenFile(name string) (i *inode, newFile bool, err os.Error) {
	err = nil

	i, oldFile := QueryFile(name)

	if !oldFile {
		log.Printf("OpenFile: file %s does not exist\n", name)
		i, err = AddFile(name)
	}
	
	newFile = !oldFile

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

	t.AddValue(name, i) // trie insert

	return i, nil
}

func QueryFile(name string) (i *inode, fileExists bool) {
	inter, exists := t.GetValue(name)

	if !exists {
		log.Printf("QueryFile: file %s does not exist\n", name)
		return nil, exists
	}

	return inter.(*inode), exists
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
	var oldID uint64 = i.chunks.At(offset).(*chunk).chunkID
	
	chunks[oldID] = &chunk{}, false
	
	i.chunks.Set(offset, newChunk)
	
	for i := 0; i < newChunk.servers.Len(); i++ {
		newChunk.servers.At(i).(*server).chunks.Push(newChunk)
	}

	return newChunk.chunkID, nil
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

	//missingCh := make(chan uint64)

	//go FindMissingChunkReplicas(missingCh)
	//go QueueReplication(missingCh)
}
