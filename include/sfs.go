package sfs

import (
	//	"os"
	"container/list"
	"net"
	"container/vector"
)

//const CHUNK_SIZE = 1024*1024*32 // 32 MB
const CHUNK_SIZE = 32                  // 32 B
const HEARTBEAT_WAIT = 15 * 1000000000 // 15 seconds

type Chunk struct {
	Data [CHUNK_SIZE]byte
}

type ReadArgs struct {
	ChunkIDs uint64
	Offsets  uint // bytes
	Lengths  uint // bytes
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
	AddedChunks   vector.Vector
}

type HeartbeatReturn struct {
	ChunksToGet vector.Vector
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
	Chunk vector.Vector // bytes
}

type ReplicateChunkArgs struct {
	ChunkID uint64
	Servers vector.Vector
}

type ReplicateChunkReturn struct {
	Status int
}

type AddChunkArgs struct {
	Name string
	Count uint64
} // returns a ChunkInfo

type Handle int

type ChunkInfo struct {
	ChunkID uint64
	Servers vector.Vector
}
