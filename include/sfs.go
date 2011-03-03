package sfs

import (
//	"os"
	"container/list"
	"net"
)

const CHUNK_SIZE = 1024*1024*32 // 32 MB

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

type WriteArgs struct {
	ChunkID uint64
	Offset uint // bytes
	Length uint // bytes
}


type HeartbeatArgs struct {
    Error uint //reserved, but can't think of anything to put here yet
}

type Status struct {
    ChunkCount uint
    ChunkIDs list.List
}

type CreateArgs struct{
	Name string
}

type CreateReturn struct{
	Confirmation int  //0 on success 1 on failure?
	ServerLocation net.IP 
	Chunk int
/// anything else?
}

type OpenArgs struct{
	Name string
	Write bool
}

type OpenReturn struct{
	Confirmation int
	Size int
	ServerLocations [5]net.IP // size ??? 
	Chunks [5]int //size???	
}


type Handle int

//func (t *Handle) Read(args *ReadArgs, chunk *Chunk) os.Error

