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
		panic("write failed")
	}

	//read first character
	ret = client.Seek(fd, 0, client.SEEK_SET)
	if(ret != 0) {
		panic("seek failed")
	}

	data, err := client.Read(fd, 1)
	if(err != 0) {
		panic("read failed")
	}

	s_ret := string(data)
	fmt.Printf("Got back: '%s'\n", s_ret)

	if(s_ret != "A") {
		panic("strings differ")
	}

	//read second character
	ret = client.Seek(fd, 2, client.SEEK_SET)
	if(ret != 0) {
		panic("seek failed")
	}

	data, err = client.Read(fd, 1)
	if(err != 0) {
		panic("read failed")
	}

	s_ret = string(data)
	fmt.Printf("Got back: '%s'\n", s_ret)

	if(s_ret != "C") {
		panic("strings differ")
	}

	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

