package master

import (
	"net"
	"container/vector"
)

type server struct {
	addr net.TCPAddr
	capacity uint64
	chunks *vector.Vector
}

type serverHeap struct {
	vec * vector.Vector
}

func (s * serverHeap) Len() int           { 
	return s.vec.Len() 
}

func (s * serverHeap) Less(i, j int) bool { 
	si := s.vec.At(i).(*server)
	sj := s.vec.At(j).(*server)
	return (si.capacity/uint64(si.chunks.Len())) < (sj.capacity/uint64(sj.chunks.Len()))
}
func (s * serverHeap) Swap(i, j int)      {
	 s.vec.Swap(i,j)
}
func (s * serverHeap) Push(serv interface {}) {
	s.vec.Push(serv)
}
func (s * serverHeap) Pop() interface {} {
	return s.vec.Pop()
}

