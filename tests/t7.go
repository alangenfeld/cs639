package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"../include/sfs"
	"rand"
)

func main(){
	chunkCount := 3
	rand.Seed(12345)

	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)

	fd := client.Open("/newfile.txt", client.O_RDWR|client.O_CREATE)
	if(fd < 0) {
		panic("could not create new file")
	}

	s := ""
	for i := 0; i < chunkCount*sfs.CHUNK_SIZE; i++ {
		var c [1]byte
		c[0] = uint8(65+rand.Intn(25))
		s = s + string(c[:])
	}

	fmt.Printf("String: %s\n", s);
	buf := []byte(s)

	ret := client.Write(fd, buf)
	if(ret != 0) {
		panic("write failed")
	}

	//read data
	ret = client.Seek(fd, 0, client.SEEK_SET)
	if(ret != 0) {
		panic("seek failed")
	}

	data, err := client.Read(fd, chunkCount*sfs.CHUNK_SIZE)
	if(err != 0) {
		panic("read failed")
	}

	s_ret := string(data)
	fmt.Printf("Got back: '%s'\n", s_ret)

	if(s_ret != s) {
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

