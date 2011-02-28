package bloom

import (
	"./bob"
	"io"
	//"fmt"
)

type Bloom struct{
	filter [(1<<16) - 1]uint8
	enable bool
}

func New() *Bloom{
	b := new(Bloom)
	b.enable = true
	return b
}

func (b *Bloom) Test(key string) bool {
	if !b.enable {
		return true
	}

	head, tail := genhash(&key)
	
	//fmt.Printf("***TEST*** key: %s, head: %d, tail: %d\n", key, head, tail)

	if b.filter[head] > 0 && b.filter[tail] > 0 {
		return true
	}
	
	return false
}

func (b *Bloom) Set(key string) {
	if !b.enable {
		return
	}
	
	head, tail := genhash(&key)

	if b.filter[head] > 100 && b.filter[tail] > 100 {
		//fmt.Printf("shit is full!\n")
		b.enable = false
	}

	b.filter[head] += 1
	b.filter[tail] += 1
	
	//fmt.Printf("***SET*** key: %s, head: %d, tail: %d, hcnt: %d, tcnt: %d\n",
	//	key, head, tail, b.filter[head], b.filter[tail])
	
}

func genhash(key *string) (head uint16, tail uint16){
	h := bob.New()
	io.WriteString(h, *key)
	s := h.Sum32()
	
	head = uint16(s >> 16)
	tail = uint16(s & 0x0000FFFF)
	
	return head, tail	
}