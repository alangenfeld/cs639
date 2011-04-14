package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
)

func main(){
	master := flag.String("m", "", "specify a master (-m)")
	source := flag.String("s", "", "specify a source file (-s)")
	dest := flag.String("d", "", "specify a destination file (-d)")
	flag.Parse();

	client.Initialize(*master)


	f, err := ioutil.ReadFile(*source)
	if err != nil {	
		fmt.Printf("Open %s locally failed\n", *source)
		os.Exit(1)
	}

	fd := client.Open(*dest, client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		fmt.Printf("Open %s in SFS failed\n", *dest)
		os.Exit(1)
	} else {
		fmt.Printf("Open worked!\n")
	}

	fmt.Printf("Writing...\n")
	ret := client.Write(fd, f)
	if(ret != 0) {
		fmt.Printf("Write failed\n")
		os.Exit(1)
	} else {
		fmt.Printf("Write worked!\n")
	}

	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
}

