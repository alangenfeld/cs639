package master

import(
	"net"
	"./trie"
	"fmt"
	"os"
)

type inode struct {
	name string
	permissions uint64
	size uint64
	nchunks uint64
}

type chunk struct {
	chunkId		uint64
	pFileChunk	*chunk
	nFileChunk	*chunk
	pServChunk	*chunk
	nServChunk	*chunk
	server		*server
}

type server struct {
	addr net.IPAddr
	sChunkHead *chunk
}

func AddFile(name string) err os.Error {
	i := new(inode)
	
	t.AddValue(name, inode)
}

func GetFile(name string) (i *inode, status bool) {
	
}

func init() {
	t := trie.NewTrie()
}