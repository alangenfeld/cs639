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
	fd := client.Open("/newfile.txt", client.O_RDWR|client.O_CREATE)
	if(fd < 0) {
		panic("could not create new file")
	}

	
	s := randString(10.5*sfs.CHUNK_SIZE)
	ret = client.Write(fd, []byte(s))
	if(ret != 0) {
		panic("write failed")
	}
	
	offset := int(0.5*sfs.CHUNK_SIZE)

	seekResult := client.Seek(fd, 0.5*sfs.CHUNK_SIZE, client.SEEK_SET)
	if(seekResult != offset) {
		fmt.Printf("Expected %d but got %d\n", offset, seekResult)
		panic("seek failed")
	}


	for i:=0; i < 10; i++ {
		val, err := client.Read(fd, sfs.CHUNK_SIZE)
		sret := string(val)

		if(err != 0) {
			panic("read failed")
		}

		fmt.Printf("Got %d: %s\n", i, sret)
		if(sret != s[offset:offset+sfs.CHUNK_SIZE]){
			fmt.Printf("Strings differ: %s, %s", sret, s[offset:offset+sfs.CHUNK_SIZE])
			panic("Strings differ")
		}

		offset += sfs.CHUNK_SIZE
	}


	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

