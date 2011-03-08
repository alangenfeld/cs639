package sfs

import (
//	"os"
	"container/list"
	"net"
  "container/vector"
)

//const CHUNK_SIZE = 1024*1024*32 // 32 MB
const CHUNK_SIZE = 32 // 32 B
const HEARTBEAT_WAIT = 15 * 1000000000 // 15 seconds

type Chunk struct {
	Data [CHUNK_SIZE]byte
}

type ReadArgs struct {
	ChunkID uint64
	Offset uint // bytes
	Length uint // bytes
}

type ReadReturn struct {
	Data Chunk
	Status int
}

type ChunkBirthArgs struct {
	ChunkServer net.TCPAddr
	Capacity uint64
}
type ChunkBirthReturn struct {
}


type WriteArgs struct {
	ChunkID uint64
	Data Chunk
	Offset uint // bytes
	Length uint // bytes
}

type WriteReturn struct{
	Status int
}

type HeartbeatArgs struct {
    capacity uint64
    addedChunks vector.Vector
}

type HeartbeatReturn struct {
  chunksToGet vector.Vector
}

type Status struct {
    ChunkCount uint
    ChunkIDs list.List
}

type OpenArgs struct{
	Name string
}

type OpenReturn struct{
	New bool
	Size uint64
	ServerLocation net.TCPAddr // size ??? 
	Chunk uint64 //size???	
}


type Handle int

//func (t *Handle) Read(args *ReadArgs, chunk *Chunk) os.Error

