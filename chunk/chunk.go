// Filesystem implementation of chunk server (lazzzzy)
package chunk

import (
	"../include/sfs" 
	"log"
	"os"
	"strconv"
)

type Server int

func (t *Server) Read(args *sfs.ReadArgs, chunk *sfs.Chunk) os.Error {

	fileName := strconv.Uitob64(args.ChunkID, 10)
	file, err := os.Open(fileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("os.Open error: ", err)
	}

	_, err = file.Read(chunk.Data[:])
	if err != nil {
		log.Fatal("file.Read error: ", err)
	}

	file.Close()

	return nil
}
