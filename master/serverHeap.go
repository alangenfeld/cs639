package master

import (
	"net"
	"container/vector"
)

type server struct {
	addr net.TCPAddr
	capacity uint
	chunks *vector.Vector
}

type ServerHeap * vector.Vector

func (s ServerHeap) Len() int           { 
	return Len(s) 
}

func (s ServerHeap) Less(i, j int) bool { 
	si := s.at(i).(*server)
	sj := s.at(j).(*server)
	return (si.capacity/Len(si.chunks)) < (sj.capacity/Len(sj.chunks))
}
func (s ServerHeap) Swap(i, j int)      {
	 s.Swap(i,j)
}
func (s ServerHeap) Push(serv * server) {
	s.Push(server)
}
func (s ServerHeap) Pop() * server {
	return s.Pop().(*server)
}

