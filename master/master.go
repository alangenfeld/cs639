package main

type inode struct (
	name string
	permissions uint64
	size uint64
	nchunks uint64
)

import (
	"fmt"
)

func main(){
	
	fmt.println("the master is working\n")
	
}
