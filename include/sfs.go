package sfs

import (
//	"os"
	"container/list"
	"net"
)

//const CHUNK_SIZE = 1024*1024*32 // 32 MB
const CHUNK_SIZE = 32 // 32 B

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

type PingArgs struct {
	ChunkServer net.TCPAddr
	Capacity uint64
}
type PingReturn struct {

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
    Error uint //reserved, but can't think of anything to put here yet
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

