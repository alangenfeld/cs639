package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"../include/sfs"
	"rand"
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

	//create file
	fd := client.Open("/newfile.txt", client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		panic("could not create new file")
	}
	
	s := randString(10.5*sfs.CHUNK_SIZE)
	ret = client.Write(fd, []byte(s))
	if(ret != 0) {
		panic("write failed")
	}

	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}



	//READ THE FILE

	//open file
	fd = client.Open("/newfile.txt", client.O_RDONLY)
	if(fd < 0) {
		panic("could not open file")
	}
	
	buf, ret := client.Read(fd, 12*sfs.CHUNK_SIZE)
	if(ret != 0) {
		panic("write failed")
	}

	if(string(buf) != s) {
		fmt.Printf("Got %s, expected %s", string(buf), s)
		panic("Strings differ")
	}

	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}


	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

