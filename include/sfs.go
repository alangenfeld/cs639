package sfs

import (
	//	"os"
	"container/list"
	"net"
)

//const CHUNK_SIZE = 1024*1024*32 // 32 MB
const CHUNK_SIZE = 32                  // 32 B
const HEARTBEAT_WAIT = 15 * 1000000000 // 15 seconds
const NREPLICAS = 3

type Chunk struct {
	Data [CHUNK_SIZE]byte
}

type ReadArgs struct {
	ChunkID uint64
}

type ReadReturn struct {
	Data   Chunk
	Status int
}

type ChunkBirthArgs struct {
	ChunkServerIP net.TCPAddr
	Capacity      uint64
}
type ChunkBirthReturn struct {
	ChunkServerID uint64
}

type WriteArgs struct {
	Info   ChunkInfo
	Data   Chunk
	Offset uint // bytes
	Length uint // bytes
}

type WriteReturn struct {
	Status int
}

type HeartbeatArgs struct {
	ChunkServerIP net.TCPAddr
	ChunkServerID uint64
	Capacity      uint64
	AddedChunks   []ChunkInfo
}

type HeartbeatReturn struct {
	ChunksToGet []ChunkInfo
}

type Status struct {
	ChunkCount uint
	ChunkIDs   list.List
}

type OpenArgs struct {
	Name string
	Size uint64
}

type OpenReturn struct {
	New   bool
	Size  uint64        // bytes
	Chunk []ChunkInfo// bytes
}

type ReplicateChunkArgs struct {
	ChunkID uint64
	Servers []net.TCPAddr
}

type ReplicateChunkReturn struct {
	Status int
}

type AddChunkArgs struct {
	Name string
	Count uint64
}

type AddChunkReturn struct {
	Info ChunkInfo
}

type RemoveArgs struct {
	Name string
}

type RemoveReturn struct {
	Success bool
}

type Handle int

type ChunkInfo struct {
	ChunkID uint64
	Servers []net.TCPAddr
}
