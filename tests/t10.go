package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
//	"../include/sfs"
//	"rand"
)

func createPath(path string) {
	fd := client.Open(path, client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		panic("open failed")
	}
	ret := client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
}

func main(){
	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)

	if(client.MakeDir("/a") != 0) {
		panic("make dir failed")
	}
	if(client.MakeDir("/a/b") != 0) {
		panic("make dir failed")
	}

	createPath("/a/b/c1")
	createPath("/a/b/c2")
	createPath("/a/b/c3")

	children, err := client.ReadDir("/a")
	if(err != 0) {
		panic("readDir failed")
	}

	if(len(children) != 1) {
		panic("wrong number of children from ReadDir")
	}

	val := children[0]
	if(val != "b") {
		panic("wrong child from ReadDir")
	}


	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

