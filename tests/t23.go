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


	if(client.Open("/a/b/c", client.O_WRONLY|client.O_CREATE) != -1) {
		panic("open w/ create shouldn't work")
	}

	if(client.MakeDir("/a/b") == 0) {
		panic("makedir shouldn't work")
	}

	if(client.MakeDir("") == 0) {
		panic("makedir shouldn't work")
	}

	ls, err := client.ReadDir("");
	if(err == 0) {
		panic("readdir shouldn't work")
	}

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

