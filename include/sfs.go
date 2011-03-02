package sfs

import (
//	"os"
	"container/list"
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

type WriteArgs struct {
	ChunkID uint64
	offset uint // bytes
	length uint // bytes
}

type HeartbeatArgs struct {
        error uint //reserved, but can't think of anything to put here yet
}

type Status struct {
        ChunkCount uint
        ChunkIDs list.List
}

type Handle int

//func (t *Handle) Read(args *ReadArgs, chunk *Chunk) os.Error
