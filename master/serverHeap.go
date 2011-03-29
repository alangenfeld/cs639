package master

import (
	"net"
	"container/vector"
	"container/heap"
)

type server struct {
	addr net.TCPAddr
	id uint64
	capacity uint64
	chunks *vector.Vector
}

type heapCommand struct {
	command uint64
	server interface {}
}
type serverHeap struct {
	serverChan chan * heapCommand
	vec * vector.Vector
}

func (s * serverHeap) Len() int { 
	return s.vec.Len() 
}

func (s * serverHeap) Less(i, j int) bool { 
	si := s.vec.At(i).(*server)
	sj := s.vec.At(j).(*server)
	if(si.chunks.Len() < 1 || sj.chunks.Len() < 1){
		return true;
	}else{
		return (si.capacity/uint64(si.chunks.Len())) < (sj.capacity/uint64(sj.chunks.Len()))
	}
	return false
}
func (s * serverHeap) Swap(i, j int)      {
	 s.vec.Swap(i,j)
}
func (s * serverHeap) Push(serv interface {}) {
	s.serverChan <- &heapCommand{0,serv}
	//s.vec.Push(serv)
}
func (s * serverHeap) Pop() interface {} {
	s.serverChan <- &heapCommand{1,nil}
	//command := new(heapCommand)
	command := <- s.serverChan
	return command.server
	//return s.vec.Pop()
}
func (s * serverHeap) Remove(serv interface{}) {
    s.serverChan <- &heapCommand{2,serv}
}
func (s * serverHeap) Handler() {

	for rec := range s.serverChan {
	
		if(rec.command == 0){
			s.vec.Push(rec.server)
		}
		if(rec.command == 1){
			s.serverChan <- &heapCommand{1,s.vec.Pop()}
		}
		if(rec.command == 2){
			deadServer := rec.server.(*server)
			//find the server to remove
			vecRange := s.vec.Len()
			for cnt := 0; cnt < vecRange; cnt++{
			
				testserv := s.vec.At(cnt).(*server)
				if(testserv.id == deadServer.id){
				
					//find each chunk to modify
					chunkRange := testserv.chunks.Len()
					for cnt1 := 0; cnt1 < chunkRange; cnt1++{
					
						tempchunk := testserv.chunks.At(cnt1).(*chunk)
						
						//find the server to remove from EACH CHUNK LIST
						searchRange := tempchunk.servers.Len()
						for cnt2 := 0; cnt2 < searchRange; cnt2++ {
						
							tempserv := tempchunk.servers.At(cnt2).(*server)
							if(tempserv.id == deadServer.id){
								tempchunk.servers.Delete(cnt2)
							}
						}	
					}
					s.vec.Delete(cnt)
				}
			}
			heap.Init(s)
		}
	}
}


