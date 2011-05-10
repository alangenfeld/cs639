package master

import (
	"net"
	"log"
	"container/vector"
	"container/heap"
	"strconv"
	"fmt"
	"time"
)

type server struct {
	addr net.TCPAddr
	id uint64
	capacity uint64
	chunks *vector.Vector
	evictedChunks *vector.Vector //uint64s
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
		//return (si.capacity/uint64(si.chunks.Len())) > (sj.capacity/uint64(sj.chunks.Len()))
		return uint64(si.chunks.Len()) > uint64(sj.chunks.Len())
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
	var rec *heapCommand

	ticker := time.Tick(30 * 1000000000)
	for{
		select{
		case rec = <- s.serverChan:
			if(rec.command == 0){
				s.vec.Push(rec.server)
				rec.retChan <- &heapCommand{}
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
				rec.retChan <- &heapCommand{}
			}
		case <- ticker :
			log.Printf("master: scheduled reheap\n")
			log.Printf("BEFORE \n %s \n", s.printPresent())
			heap.Init(s)
			log.Printf("AFTER  \n %s \n", s.printPresent())
		}
	}
}

func (sh *serverHeap) printPresent() string {
	var out string = ""
	cnt := sh.vec.Len()
	for i := 0; i < cnt; i++ {
		s := sh.vec.At(i).(*server)
		out += "\ts.id: " + strconv.Uitoa64(s.id) + " addr: " + fmt.Sprintf("%s:%d", s.addr.IP.String(), s.addr.Port) + " nchunks: " + strconv.Itoa(s.chunks.Len()) + " chunks: "
		
		cnt2 := s.chunks.Len()
		for j := 0; j < cnt2; j++ {
			out += strconv.Uitoa64(s.chunks.At(j).(*chunk).chunkID) + ", "
		}
		
		out += "\n"
	}
	
	//log.Printf("master: server heap state follows:\n%s", out)
	return out
}


