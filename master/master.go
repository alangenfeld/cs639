package master

import (
	"net"
	"./trie"
	"log"
	"os"
	"fmt"
	"container/vector"
	"container/heap"
	"../include/sfs"
)

var t 			*trie.Trie
var nextChunk	uint64 = 0
var serverIndex	uint64 = 0
var sHeap		*serverHeap
var chunks		map[uint64](*chunk)

var nReplicas	int = 1

type inode struct {
	name        string
	permissions uint64
	size        uint64
	chunks      *vector.Vector
}

type chunk struct {
	chunkID		uint64
	servers		*vector.Vector
}

type Master int

func (m *Master) ReadOpen(args *sfs.OpenArgs, info *sfs.OpenReturn) os.Error {
	i, newFile, err := OpenFile(args.Name)

	info.New = newFile
	info.Size = i.size
	info.Chunk = (i.chunks.At(0).(*chunk)).chunkID
	info.ServerLocation = ((i.chunks.At(0).(*chunk)).servers.At(0).(*server)).addr

	return err
}

func (m *Master) BirthChunk(args *sfs.ChunkBirthArgs, info *sfs.ChunkBirthReturn) os.Error {
	AddServer(args.ChunkServer, args.Capacity)

	return nil
}

func (m *Master) BeatHeart(args *sfs.HeartbeatArgs, info *sfs.HeartbeatReturn) os.Error {

	return nil

}
/*func (m *Master) WriteReplicationReq(c *chunk) (err os.Error){
	return nil
}*/

func OpenFile(name string) (i *inode, newFile bool, err os.Error) {
	err = nil

	i, newFile = QueryFile(name)

	newFile = !newFile

	if newFile {
		log.Printf("OpenFile: file %s does not exist\n", name)
		i, err = AddFile(name)
	}

	return i, newFile, err
}

func AddFile(name string) (i *inode, err os.Error) {
	i = new(inode)

	log.Printf("AddFile: nextChunk %d, len(servers) %d\n", nextChunk, sHeap.Len())

	i.size = 1
	//i.addr = *(servers.At(int(nextChunk) % servers.Len()).(*net.TCPAddr))
	//i.addr = servers[0]
	
	i.chunks = new(vector.Vector)

	i.AddChunk()

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

func (i *inode) AddChunk() (chunkID uint64) {
	var serv *server = heap.Pop(sHeap).(*server)
	thisChunk := new(chunk)
	thisChunk.chunkID = nextChunk
	nextChunk += 1

	thisChunk.servers = new(vector.Vector)

	thisChunk.servers.Push(serv)
	serv.chunks.Push(thisChunk)
	i.chunks.Push(thisChunk)

	heap.Push(sHeap, serv)
	
	chunks[thisChunk.chunkID] = thisChunk

	return thisChunk.chunkID
}

func AddServer(servAddr net.TCPAddr, capacity uint64) os.Error {
	str := fmt.Sprintf("%s:%d", servAddr.IP.String(), servAddr.Port)
	log.Printf("AddServer: adding %s\n", str)

	s := new(server)

	s.addr = servAddr
	s.capacity = capacity
	s.chunks = new(vector.Vector)

	heap.Push(sHeap, s)

	return nil
}

func FindMissingChunkReplicas() (ret uint64) {
	for cID, _ := range chunks{
		if chunks[cID].servers.Len() < nReplicas {
			ret += 1
			replicateChunk(cID)
		}
	}
	
	return ret
}

func replicateChunk(cID uint64) (err os.Error){
	//var m Master
	
	//m.WriteReplicationReq(chunks[cID])
	
	return nil
}

func init() {
	t = trie.NewTrie()
	sHeap = new(serverHeap)
	sHeap.vec = new(vector.Vector)
	chunks = make(map[uint64](*chunk))
	heap.Init(sHeap)
	
	//missingCh := make(chan uint64)
	
	//go FindMissingChunkReplicas(missingCh)
	//go QueueReplication(missingCh)
}
