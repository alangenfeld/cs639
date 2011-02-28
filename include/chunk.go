package chunk

import (
	"os"
)

const CHUNK_SIZE = 1024*1024*32 // 32 MB

type Chunk struct {
	data [CHUNK_SIZE]byte
}

type ReadArgs struct {
	ChunkID uint64
	offset uint // bytes
	length uint // bytes
}

type SFS int

func (t *SFS) Read(args *ReadArgs, chunk *Chunk) os.Error
