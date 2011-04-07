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


	s := randString(1*sfs.CHUNK_SIZE)
	s2 := randString(1*sfs.CHUNK_SIZE)


	//test O_RDWR
	fd := client.Open("newfile.txt", client.O_RDWR|client.O_CREATE)
	if(fd < 0) {
		panic("could not create new file")
	}
	
	ret = client.Write(fd, []byte(s))
	if(ret != 0) {
		panic("write failed")
	}
	client.Seek(fd, 0, client.SEEK_SET)

	buf, ret := client.Read(fd, 1*sfs.CHUNK_SIZE)
	if(ret != 0) {
		panic("read failed")
	}

	if(string(buf) != s) {
		fmt.Printf("%s != %s\n", string(buf), s)
		panic("read returned wrong data")
	}

	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}









	//test O_RDONLY
	fd = client.Open("newfile.txt", client.O_RDONLY)
	if(fd < 0) {
		panic("could not open file")
	}

	buf, ret = client.Read(fd, 1*sfs.CHUNK_SIZE)
	if(ret != 0) {
		panic("read failed")
	}
	if(string(buf) != s) {
		fmt.Printf("%s != %s\n", string(buf), s)
		panic("read returned wrong data")
	}

	client.Seek(fd, 0, client.SEEK_SET)
	ret = client.Write(fd, []byte(s2))
	if(ret != -1) {
		panic("write didn't fail even though it should have")
	}

	client.Seek(fd, 0, client.SEEK_SET)
	buf, ret = client.Read(fd, 1*sfs.CHUNK_SIZE)
	if(ret != 0) {
		panic("read failed")
	}
	if(string(buf) != s) {
		panic("data was updated even though we had O_WRONLY")
	}

	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}









	//test O_WRONLY
	fd = client.Open("newfile.txt", client.O_WRONLY)
	if(fd < 0) {
		panic("could not open file")
	}

	buf, ret = client.Read(fd, 1*sfs.CHUNK_SIZE)
	if(ret != -1) {
		panic("read didn't fail, even though we have O_WRONLY")
	}

	ret = client.Seek(fd, 0, client.SEEK_CURR)
	if(ret != 0) {
		panic("file pointer was changed improperly")
	}

	ret = client.Write(fd, []byte(s2))
	if(ret != 0) {
		panic("write failed")
	}

	ret = client.Seek(fd, 0, client.SEEK_CURR)
	if(ret != sfs.CHUNK_SIZE) {
		panic("read failed")
	}

	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}







	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

