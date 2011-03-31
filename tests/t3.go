package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"../include/sfs"
)

func main(){


	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)

	fd := client.Open("newfile.txt", client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		panic("could not create new file")
	}

	s := ""
	for i := 0; i < sfs.CHUNK_SIZE; i++ {
		c := "A"
		if(i%3 == 1) {
			c = "B"
		} else if (i%3 == 2) {
			c = "C"
		}
		s = s + c
	}

	fmt.Printf("String: %s\n", s);
	buf := []byte(s)

	ret := client.Write(fd, buf)
	if(ret != 0) {
		fmt.Printf("Expected to write %d bytes, but actually wrote %d\n", len(buf), ret)
		panic("write wrong amount of data")
	}

	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

