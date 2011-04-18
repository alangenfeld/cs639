package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
//	"../include/sfs"
)

func main(){
	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)

	fd := client.Open("/newfile.txt", client.O_WRONLY)
	if(fd != client.FAIL) {
		panic("no fail on open even though file doesn't exists")
	}

	fd = client.Open("/newfile.txt", client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		panic("could not create new file")
	}

	ret := client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}

	fd = client.Open("/newfile.txt", client.O_WRONLY)
	if(fd < 0) {
		panic("couldn't open file that was created earlier")
	}
	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

