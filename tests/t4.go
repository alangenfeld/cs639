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

	fd := client.Open("newfile.txt", client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		panic("could not create new file")
	}

	s := "A"

	fmt.Printf("String: %s\n", s);
	buf := []byte(s)

	ret := client.Write(fd, buf)
	if(ret != 0) {
		panic("write failed")
	}

	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

