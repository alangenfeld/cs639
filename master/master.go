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
	"strings"
	"rand"
	"math"
	"os/signal"
	"runtime"
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
var hashToChunkMap map[string](*chunk)

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
	hash	[]byte
	refCt	uint64
}

type Master int

func (m *Master) ReadOpen(args *sfs.OpenArgs, info *sfs.OpenReturn) os.Error {

	if sHeap.vec.Len() == 0 {
		err := os.NewError("No chunk servers!")
		return err
	}
	file, newFile, err := OpenFile(args.Name, args.NewFile)
	log.Println("CREATE? ", args.NewFile)
	
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
	
	for i := 0; i < file.chunks.Len(); i++ {
		thisChunk := file.chunks.At(i).(*chunk)
		info.Chunk[i].ChunkID = thisChunk.chunkID
		info.Chunk[i].Size = thisChunk.size
		info.Chunk[i].Hash = thisChunk.hash
		
		info.Chunk[i].Servers = make([]net.TCPAddr, thisChunk.servers.Len())
		
		for j := 0; j < thisChunk.servers.Len(); j++ {
			info.Chunk[i].Servers[j] = thisChunk.servers.At(j).(*server).addr
		}
	}

	return err
}

func (m *Master) MapChunkToFile(args *sfs.MapChunkToFileArgs, ret *sfs.MapChunkToFileReturn) os.Error {
	file, ok , error := QueryFile(args.Name)

	if !ok {
		log.Printf("master: MapChunkToFile: File %s does not exist\n", args.Name)
		return error
	}

	log.Printf("master: MapChunkToFile: ChunkID: %d  Offset: %d  nservers: %d Hash: %x\n", args.Chunk.ChunkID, args.Offset, len(args.Chunk.Servers), args.Chunk.Hash)
	
	thisChunk, ok := chunks[args.Chunk.ChunkID]
	
	if !ok {
		thisChunk = new(chunk)

		thisChunk.chunkID = args.Chunk.ChunkID
		thisChunk.size = args.Chunk.Size
		thisChunk.servers = new(vector.Vector)
		for i := 0; i < len(args.Chunk.Servers); i++ {
			thisChunk.AssociateServer(addrToServerMap[args.Chunk.Servers[i].String()])
		}
		thisChunk.hash = args.Chunk.Hash
	}
	
	_, err := file.MapChunk(args.Offset, thisChunk)

	if err != nil {
		return os.NewError("Could not add chunk! Ruh roh")
	}
	
	return nil
}

func (m *Master) GetNewChunk(args *sfs.GetNewChunkArgs, ret *sfs.GetNewChunkReturn) os.Error {
	ok := false
	var thisChunk *chunk
		
	if args.Hash != nil {
		thisChunk, ok = hashToChunkMap[string(args.Hash)]
	}
	
	if ok {
		log.Printf("GetNewChunk: duplicate hash found. Hash: %x ChunkID: %d\n", args.Hash, thisChunk.chunkID)
		ret.Info.ChunkID = thisChunk.chunkID
		ret.Info.Size = thisChunk.size
		ret.Info.Hash = thisChunk.hash
		ret.Info.Servers = make([]net.TCPAddr, thisChunk.servers.Len())
		
		for cnt1 := 0; cnt1 < thisChunk.servers.Len(); cnt1++ {
			ret.Info.Servers[cnt1] = thisChunk.servers.At(cnt1).(*server).addr
		}
		
		ret.NewChunk = false
	} else {
		log.Printf("GetNewChunk: Hash: %x ChunkID: %d\n", args.Hash, nextChunk)
		ret.Info.ChunkID = nextChunk
		ret.Info.Hash = args.Hash

		nextChunk++

		var nreps int

		if sHeap.vec.Len() < sfs.NREPLICAS {
			nreps = sHeap.vec.Len()
		} else {
			nreps = sfs.NREPLICAS
		}

		ret.Info.Servers = make([]net.TCPAddr, nreps)
		for i := 0; i < nreps; i++ {
			ret.Info.Servers[i] = sHeap.vec.At(i).(*server).addr
		}
		ret.NewChunk = true
	}
	
	return nil
}

func (m *Master) ReportWrite(args *sfs.ReportWriteArgs, ret *sfs.ReportWriteReturn) os.Error {

	return nil
}

func (m *Master) ReadDir(args *sfs.ReadDirArgs, ret *sfs.ReadDirReturn) os.Error {
	var lookupPath string
	
	if args.Prefix == "/" {
		lookupPath = "/"
	} else {
		lookupPath = strings.TrimRight(args.Prefix, "/")
	}
	
	log.Printf("ReadDir: prefix %s, trimmed: %s\n", args.Prefix, lookupPath)
	
	dirs, files, err := t.ReadDir(lookupPath)
	
	if err != nil {
		log.Printf("ReadDir: err: %+v\n", err)
		return err
	}
	
	var cnt int
	if dirs != nil {
		cnt = dirs.Len()		
	} else {
		cnt = 0
	}
	
	retSlice := make([]string, cnt + len(files))

	var i int
	for i = 0; i < cnt; i++ {
		retSlice[i] = dirs.At(i) + "/"
	}
	
	log.Printf("ReadDir: retSlice dirs -- %+v\n", retSlice)
	log.Printf("ReadDir: retSlice file -- %+v\n", files)
	
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
func (m *Master) RemoveDir(args *sfs.RemoveDirArgs, ret *sfs.RemoveDirReturn) os.Error {
	err := t.RemoveDir(args.DirName)
	
	DumpTrie()
	
	return err
}

func (m *Master) RemoveFile(args *sfs.RemoveArgs, result *sfs.RemoveReturn) os.Error {
	result.Success = true
	name := args.Name

	i, exists , _:= QueryFile(name)
	if !exists {
		log.Printf("RemoveFile: file %s does not exist\n", name)
		err := os.NewError("You are trying to delete a file that doesn't exist.")
		result.Success = false
		return err
	} else {
		for j := 0; j < i.chunks.Len(); j++ {
			i.chunks.At(j).(*chunk).unmapChunk()
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
	}else{
		err := populateServer(s)
		if err != nil {
			log.Fatal("error populating server %v\n", s);
		}
	}

	info.ChunkServerID = s.id

	log.Println("Birthed a Chunk Server!\n")

	return nil
}

func (m *Master) ReleaseLock(args *sfs.LockReleaseArgs, ret *sfs.LockReleaseReturn) os.Error {
	file, exists ,_:= QueryFile(args.Name)
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
	
	DumpTrie()
	 
	return err
}

func (m *Master) BeatHeart(args *sfs.HeartbeatArgs, info *sfs.HeartbeatReturn) os.Error {
	str := fmt.Sprintf("%s:%d", args.ChunkServerIP.IP.String(), args.ChunkServerIP.Port)
	//log.Printf("BeatHeart: %s's HEART IS BEATING\n", str)

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
			chunk, chunkOK := chunks[args.AddedChunks[cnt].ChunkID]
			//log.Printf("Herp dDerp %d\n", args.AddedChunks[0].ChunkID)
			if chunkOK == true {
				//server.chunks.Push(chunk)
				//chunk.servers.Push(server)
				AssociateChunkAndServer(chunk, server)
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
	
	old_server,check := addrToServerMap[servAddr.String()];
	if check{
		err := RemoveServer(old_server)
		if err != nil {
			log.Fatal(err.String())
			return nil
		}
	}
	
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

func populateServer(serv *server) os.Error {
	str := fmt.Sprintf("%s:%d", serv.addr.IP.String(), serv.addr.Port)
	log.Printf("master: PopulateServer: populating %s\n", str)
	log.Printf("master: PopulateServer: server heap state:\n%s\n", sHeap.printPresent())

	if len(chunks) == 0 {
		return nil
	}
	
	for i := 0; i < 3; i++{
		client, err := rpc.Dial("tcp", str)
		if(client == nil){
			log.Printf("master: PopulateServer: dialing client %s, nil\n", str)
			client.Close()
			continue
		}
		
		for _, chunk := range chunks {
			log.Printf("master: PopulateServer: examining chunk %+v, nservers %d\n", *chunk, chunk.servers.Len())
			if chunk.servers.Len() < sfs.NREPLICAS {

				//populate chunk location list
				chunklist := make([]net.TCPAddr, chunk.servers.Len())
				for cnt1 := 0; cnt1 < chunk.servers.Len(); cnt1++ {
					chunklist[cnt1] = chunk.servers.At(cnt1).(*server).addr
				}

				//send rpc call off
				args := &sfs.ReplicateChunkArgs{chunk.chunkID, chunklist}
				reply := new(sfs.ReplicateChunkReturn)
				log.Printf("master: PopulateServer: issuing replication req to %s\n", str)
				err = client.Call("Server.ReplicateChunk", args, reply)
				if err != nil {
					log.Printf("master: PopulateServer: unable to call %s\n", str)
				}
				log.Printf("%s", reply)
			}
		}
		client.Close()
	}
	
	return nil
}

func RemoveServer(serv *server) os.Error {
	log.Printf("master: RemoveServer: server heap pre removal:\n%s\n", sHeap.printPresent())

	//Remove the Server
	sHeap.Remove(serv)
	log.Printf("master: RemoveServer: server heap post removal:\n%s\n", sHeap.printPresent())

	for i := 0; i < sHeap.vec.Len(); i++ {
		if sHeap.vec.At(i).(*server).id == serv.id {
			log.Printf("master: RemoveServer: remove didn't actually remove server %d! Busto Rhymes\n", serv.id)
			return nil
		}
	}

	servers[serv.id] = &server{}, false
	addrToServerMap[serv.addr.String()] = &server{}, false

	str1 := fmt.Sprintf("removing server %s:%d", serv.addr.IP.String(), serv.addr.Port)
	log.Printf("master: RemoveServer: begin %s\n", str1)
	
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

	rep_size := int(math.Fmin(math.Ceil(rep_factor * network_size), network_size))

	//for each chunk in the server, make a replication call
	sanity_threshhold1 := 4
	sanity_threshhold2 := 10
	sanity_count := 0
	
	targetServerMap := make(map[string](bool))
	
ChunkReplicate:	for cnt := 0; cnt < serv.chunks.Len(); {
		
		var index int
		if sanity_count > sanity_threshhold1 {
		    index = rand.Intn(rep_size)
		}else{
		    index = rand.Intn(sHeap.vec.Len())
		}

		otherserver := sHeap.vec.At(index).(*server)
		chunk := serv.chunks.At(cnt).(*chunk)
		
		str := fmt.Sprintf("%s:%d", otherserver.addr.IP.String(), otherserver.addr.Port)
		
		sCnt := chunk.servers.Len()
		
		if sCnt >= len(servers) {
			log.Printf("master: RemoveServer: abort replication req for chunk %d; all active servers already hold replicas\n", chunk.chunkID)
			cnt++
			sanity_count = 0
			targetServerMap = make(map[string](bool))
		}
		
		if targetServerMap[str] == true {
			continue
		}
		
		for ijk := 0; ijk < sCnt; ijk++ {
			if chunk.servers.At(ijk).(*server) == otherserver {
				targetServerMap[str] = true
				
				if len(targetServerMap) >= len(servers){
					log.Printf("master: RemoveServer: abort replication req for chunk %d; all active servers already hold replicas\n", chunk.chunkID)
					cnt++
					sanity_count = 0
					targetServerMap = make(map[string](bool))
				}
				
				continue ChunkReplicate
			}
		}
		
		log.Printf("master: RemoveServer: attempting to replicate chunk %d to server %s\n", chunk.chunkID, str)

		//log.Printf("master: RemoveServer: dialing %s to replicate\n", str)

		client, err := rpc.Dial("tcp", str)

		if err != nil {
			log.Printf("master: RemoveServer: unable to dial %s\n", str)
			sanity_count = sanity_count + 1
			continue
		} else {
			log.Printf("master: RemoveServer: dial %s succeeded\n", str)
		}

		if sanity_count > sanity_threshhold2 {
			log.Printf("master: RemoveServer: tried %d times to dial servers to replicate, and gave up!!!!\n", sanity_threshhold2)
			break
			
		}

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
			client.Close()
			continue
		}
		//log.Printf("%s", reply)
		client.Close()
		cnt++
		sanity_count = 0
		targetServerMap = make(map[string](bool))
	}

	log.Printf("master: RemoveServer: finished %s\n", str1)
	log.Printf("master: RemoveServer: server heap post replication:\n%s\n", sHeap.printPresent())
	return nil
}

func OpenFile(name string, create bool) (i *inode, newFile bool, err os.Error) {
	err = nil

	i, exists, error := QueryFile(name)

	if !exists && create {
		log.Printf("OpenFile: file %s does not exist\n", name)
		i, err = AddFile(name)
	} else if !exists && !create {
		return nil, true, error
	}

	newFile = !exists

	return i, newFile, nil
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

func QueryFile(name string) (i *inode, fileExists bool, err os.Error) {
	inter, exists, error := t.QueryFile(name)

	if !exists {
		log.Printf("QueryFile: file %s does not exist\n", name)
		return nil, exists, error
	}

	log.Printf("QueryFile: %d nodes in trie\n", t.Size())

	return inter.(*inode), exists, nil
}

func DeleteFile(name string) (err os.Error) {
	inode, exists ,_:= QueryFile(name)
	if exists {
		err := t.DeleteFile(name)
		
		if err != nil {
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
		newChunk.refCt++		
		i.chunks.Set(offset, newChunk)
	} else if offset == i.chunks.Len() {
		newChunk.refCt++
		i.chunks.Push(newChunk)
	} else {
		return 0, os.NewError("Fucking A.")
	}

	for j := 0; j < newChunk.servers.Len(); j++ {
		s := newChunk.servers.At(j).(*server)
		
		//if the server dies after right replication but before dying..
		if s != nil {
			s.AssociateChunk(newChunk)
		} else {
			newChunk.servers.Delete(j)
		}
	}

	chunks[newChunk.chunkID] = newChunk
	
	if newChunk.hash != nil {
		hashToChunkMap[string(newChunk.hash)] = newChunk
	}

	return newChunk.chunkID, nil
}

func (c *chunk) unmapChunk() (err os.Error){
	c.refCt--
	
	if c.refCt == 0 {
		return nil
	}
	
	cnt := c.servers.Len()	
	for j := 0; j < cnt; j++ {
		c.servers.At(j).(*server).evictedChunks.Push(c.chunkID)
		
		cnt2 := c.servers.At(j).(*server).chunks.Len()
		for k := 0; k < cnt2; k++ {
			if c.servers.At(j).(*server).chunks.At(k).(*chunk) == c {
				c.servers.At(j).(*server).chunks.Delete(k)
				cnt2--
			}
		}
	}
	
	chunks[c.chunkID] = &chunk{}, false
	hashToChunkMap[string(c.hash)] = &chunk{}, false
	
	return nil
}

func (c *chunk) AssociateServer(s *server) os.Error {
	cnt := c.servers.Len()
	for i := 0; i < cnt; i++ {
		if c.servers.At(i).(*server) == s {
			log.Printf("master: AssociateServer: dupe found\n")
			return nil
		}
	}
	
	c.servers.Push(s)
	return nil
} 

func (s *server) AssociateChunk(c *chunk) os.Error {
	cnt := s.chunks.Len()
	for i := 0; i < cnt; i++ {
		if s.chunks.At(i).(*chunk) == c {
			log.Printf("master: AssociateChunk: dupe found\n")
			return nil
		}
	}
	
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
			//log.Printf("\n\nchunk map: %+v \n\nhashToChunks map: %+v\n\n", chunks, hashToChunkMap)
			log.Printf("\n\nchunk map len: %d \n\nhashToChunks map len: %d\n\n", len(chunks), len(hashToChunkMap))
			DumpTrie()
			os.Exit(1337)
		} else if sig.String() == "SIGHUP: terminal line hangup" {
			runtime.GC()
			log.Printf("Garbage collection complete!")
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
	//log.Printf("dumpTrie: DOT file follows\n******\n%s\n******\n", t.GetDotString())
}

func init() {
	t = trie.NewTrie()

	sHeap = new(serverHeap)
	sHeap.vec = new(vector.Vector)
	chunks = make(map[uint64](*chunk))
	servers = make(map[uint64](*server))
	addrToServerMap = make(map[string](*server))
	heartbeatMonitors = make(map[uint64](chan int64))
	hashToChunkMap = make(map[string](*chunk))
	//	sMap = make(map[net.TCPAddr](*server))
	heap.Init(sHeap)
	sHeap.serverChan = make(chan *heapCommand)
	go sHeap.Handler()
	go sigHandler()

	t.AddDir("/")

	//missingCh := make(chan uint64)

	//go FindMissingChunkReplicas(missingCh)
	//go QueueReplication(missingCh)
}
