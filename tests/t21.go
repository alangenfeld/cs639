package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"../include/sfs"
	"rand"
	"time"
)

func randString(n int) string {
	c := make([]byte, n)

	for i := 0; i < n; i++ {
		c[i] = uint8(65+rand.Intn(25))
	}

	return string(c[:])
}


func main(){
	var ret int
	rand.Seed(12345)

	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)

	start := time.Nanoseconds();
	now := time.Nanoseconds();
	for ; (now - start) < 60*1000*1000*1000; {
		//create file
		fd := client.Open("/file.txt", client.O_RDWR|client.O_CREATE)
		if(fd < 0) {
			panic("could not create new file")
		}

		s := randString(3*sfs.CHUNK_SIZE)
		ret = client.Write(fd, []byte(s))
		if(ret != 0) {
			panic("write failed")
		}

		ret = client.Close(fd)
		if(ret != 0) {
			panic("close failed")
		}


		//read file back
		fd = client.Open("/file.txt", client.O_RDWR|client.O_CREATE)
		if(fd < 0) {
			panic("could not open file")
		}

		buf2, ret2 := client.Read(fd, 3*sfs.CHUNK_SIZE)
		if(ret2 != 0) {
			panic("read failed")
		}
		if(string(buf2) != s) {
			panic("wrong data returned")
		}

		ret = client.Close(fd)
		if(ret != 0) {
			panic("close failed")
		}

		now = time.Nanoseconds();
	}

	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

