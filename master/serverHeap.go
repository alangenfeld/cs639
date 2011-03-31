package master

import (
	"net"
	"log"
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
	server interface{}
	retChan chan * heapCommand
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
func (s * serverHeap) Push(serv interface{}) {
	ch := make(chan *heapCommand)
	s.serverChan <- &heapCommand{0,serv,ch}
	<-ch
	//s.vec.Push(serv)
}
func (s * serverHeap) Pop() interface {} {
	ch := make(chan *heapCommand)
	s.serverChan <- &heapCommand{1,nil,ch}
	//command := new(heapCommand)
	command := <- ch
	return command.server
	//return s.vec.Pop()
}
func (s * serverHeap) Remove(serv interface{}) {
	ch := make(chan *heapCommand)
    s.serverChan <- &heapCommand{2,serv,ch}
	<-ch
}
func (s * serverHeap) Handler() {

	for rec := range s.serverChan {
	
		if(rec.command == 0){
			s.vec.Push(rec.server)
		}
		if(rec.command == 1){
			rec.retChan <- &heapCommand{1,s.vec.Pop(),nil}
		}
		if(rec.command == 2){
			deadServer := rec.server.(*server)
			//find the server to remove
			for cnt := 0; cnt < s.vec.Len(); cnt++ {
			
				testserv := s.vec.At(cnt).(*server)
				if(testserv.id == deadServer.id){
					log.Printf("master: heap Handler: found server %d to delete\n", deadServer.id)
				
					//find each chunk to modify
					for cnt1 := 0; cnt1 < testserv.chunks.Len(); cnt1++{
					
						tempchunk := testserv.chunks.At(cnt1).(*chunk)
						
						//find the server to remove from EACH CHUNK LIST
						for cnt2 := 0; cnt2 < tempchunk.servers.Len(); cnt2++ {
						
							tempserv := tempchunk.servers.At(cnt2).(*server)
							if(tempserv.id == deadServer.id){
								tempchunk.servers.Delete(cnt2)
							}
						}	
					}
					prevCnt := s.vec.Len()
					s.vec.Delete(cnt)
					log.Printf("master: heap Handler: prev vec count %d, now %d\n", prevCnt, s.vec.Len())
					break
				}
			}
			heap.Init(s)
			ch <- heapCommand{}
		}
	}
}


