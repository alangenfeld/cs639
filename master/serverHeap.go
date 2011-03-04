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

type serverHeap * vector.Vector

func (s serverHeap) Len() int           { 
	return Len(s) 
}

func (s serverHeap) Less(i, j int) bool { 
	si := s.at(i).(*server)
	sj := s.at(j).(*server)
	return (si.capacity/Len(si.chunks)) < (sj.capacity/Len(sj.chunks))
}
func (s serverHeap) Swap(i, j int)      {
	 s.Swap(i,j)
}
func (s serverHeap) Push(serv * server) {
	s.Push(server)
}
func (s serverHeap) Pop() * server {
	return s.Pop().(*server)
}

