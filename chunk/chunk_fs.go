package chunk

import (
	"../include/sfs" 
	"log"
	"os"
	"strconv"
)

var CurrentStatus* sfs.Status
var currentFD int

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

func (t *Server) Write(args *sfs.WriteArgs, chunk *sfs.Chunk) os.Error {
	
	// this assumes that master tells client what chunk number it is allowed to write to
	// but this could be an incorrect assumption
	// if opening chunk doesn't work, assume it doesn't exist and create one	
	fileName := strconv.Uitob64(args.ChunkID, 10)
	file, err := os.Open(fileName, os.O_RDONLY, 0666)
	if err != nil {
		file = os.NewFile(currentFD, fileName)
		currentFD++
	}
	_, err = file.Write(chunk.Data[:])
	if err != nil {
		log.Fatal("file.Write error: ", err);
	}
	file.Close()

	return nil
}

func (t *Server) Heartbeat(args *sfs.HeartbeatArgs, status *sfs.Status) os.Error {
	status = CurrentStatus
        return nil
}

func Init() {
        CurrentStatus := new(sfs.Status)
	CurrentStatus.ChunkCount = 0
	currentFD = 100
        return
}
