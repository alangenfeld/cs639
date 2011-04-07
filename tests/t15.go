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


	s := randString(1.5*sfs.CHUNK_SIZE)


	//test O_RDWR
	fd := client.Open("newfile.txt", client.O_RDWR|client.O_CREATE)
	if(fd < 0) {
		panic("could not create new file")
	}
	

	if(client.Seek(fd, 0, client.SEEK_CURR) != 0) {
		panic("bad seek")
	}
	if(client.Seek(fd, 100, client.SEEK_CURR) != 0) {
		panic("bad seek")
	}
	if(client.Seek(fd, -100, client.SEEK_CURR) != 0) {
		panic("bad seek")
	}


	if(client.Seek(fd, 0, client.SEEK_SET) != 0) {
		panic("bad seek")
	}
	if(client.Seek(fd, 100, client.SEEK_SET) != 0) {
		panic("bad seek")
	}
	if(client.Seek(fd, -100, client.SEEK_SET) != 0) {
		panic("bad seek")
	}


	if(client.Seek(fd, 0, client.SEEK_END) != 0) {
		panic("bad seek")
	}
	if(client.Seek(fd, 100, client.SEEK_END) != 0) {
		panic("bad seek")
	}
	if(client.Seek(fd, -100, client.SEEK_END) != 0) {
		panic("bad seek")
	}



	//DO WRITE
	ret = client.Write(fd, []byte(s))
	if(ret != 0) {
		panic("write failed")
	}



	if(client.Seek(fd, 0, client.SEEK_CURR) != 1.5*sfs.CHUNK_SIZE) {
		panic("bad seek")
	}
	if(client.Seek(fd, 2*sfs.CHUNK_SIZE, client.SEEK_CURR) != 1.5*sfs.CHUNK_SIZE) {
		panic("bad seek")
	}
	if(client.Seek(fd, -2*sfs.CHUNK_SIZE, client.SEEK_CURR) != 0) {
		panic("bad seek")
	}


	if(client.Seek(fd, 0, client.SEEK_SET) != 0) {
		panic("bad seek")
	}
	if(client.Seek(fd, 2*sfs.CHUNK_SIZE, client.SEEK_SET) != 1.5*sfs.CHUNK_SIZE) {
		panic("bad seek")
	}
	if(client.Seek(fd, -2*sfs.CHUNK_SIZE, client.SEEK_SET) != 0) {
		panic("bad seek")
	}


	if(client.Seek(fd, 0, client.SEEK_END) != 1.5*sfs.CHUNK_SIZE) {
		panic("bad seek")
	}
	if(client.Seek(fd, 2*sfs.CHUNK_SIZE, client.SEEK_END) != 1.5*sfs.CHUNK_SIZE) {
		panic("bad seek")
	}
	if(client.Seek(fd, -2*sfs.CHUNK_SIZE, client.SEEK_END) != 0) {
		panic("bad seek")
	}
	if(client.Seek(fd, -10, client.SEEK_END) != 1.5*sfs.CHUNK_SIZE-10) {
		panic("bad seek")
	}




	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}

	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

