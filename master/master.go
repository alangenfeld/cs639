package master

import(
	"net"
	"./trie"
	"fmt"
	"os"
	"container/vector"
	"../include/sfs"
)

var t *trie.Trie

var nextChunk uint64 = 0

var servers map[int]net.TCPAddr

type inode struct {
	name string
	permissions uint64
	size uint64
	chunk uint64
	addr net.TCPAddr
}

type chunk struct {
	chunkId		uint64
	pFileChunk	*chunk
	nFileChunk	*chunk
	pServChunk	*chunk
	nServChunk	*chunk
	server		*server
}

type Master int

func (m *Master) ReadOpen(args *sfs.OpenArgs, info *sfs.OpenReturn) os.Error {
	i, newFile, err := OpenFile(args.Name)
	
	info.New = newFile
	info.Size = i.size
	info.Chunk = i.Chunk
	info.ServerLocation = inode.addr
	
	return nil
}

func OpenFile(name string) (i *inode, newFile bool, err os.Error){
	err = nil
	
	i, newFile = QueryFile(args.Name)
	
	newFile = !newFile
	
	if newFile {
		i, _, err = AddFile(args.Name)
	}
	
	return i, newFile, err
}

func AddFile(name string) (i *inode, err os.Error) {
	i := new(inode)
	
	i.size = 1
	i.addr = servers[nextChunk % len(servers)]
	i.chunk = nextChunk
	nextChunk += 1
	
	
	t.AddValue(name, inode)
	
	return i, nil
}

func QueryFile(name string) (i *inode, fileExists bool) {
	return t.GetValue(name)
}

func init() {
	t = trie.NewTrie()
}