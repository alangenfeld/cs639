package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
//	"../include/sfs"
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
	//close
	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}


	if(client.Delete("/newfile.txt") != 0) {
		panic("Delete failed")
	}

	fd = client.Open("/newfile.txt", client.O_RDWR)
	if(fd != -1) {
		panic("open succeeded on deleted file")
	}

	

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}

