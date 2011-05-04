package main

import (
	"../client/client"
	"../include/sfs"
	"fmt"
	"flag"
	"os"
)

func main(){


	master := flag.String("m", "", "specify a master!")
	source := flag.String("s", "", "specify a source (-s)")
	dest := flag.String("d", "", "specify a dest (-d)")
	dest = dest

	flag.Parse();

	client.Initialize(*master)

	fd := client.Open(*source, client.O_RDONLY)
	if(fd < 0) {
		fmt.Printf("could not open file %s", *source)
		os.Exit(1)
	}

	destFile, err := os.Open(*dest, os.O_CREAT|os.O_WRONLY, 0666)
	if(err != nil) {
		fmt.Printf("Couldn't open %s\n", *dest)
	}

	fmt.Printf("BEGIN OF FILE\n")
	for ;true; {
		data, err := client.Read(fd, sfs.CHUNK_SIZE)
		if(err != 0) {
			panic("read failed")
		}

		_, err2 := destFile.Write(data)
		if(err2 != nil) {
			fmt.Printf("Write failed\n")
			os.Exit(1)
		}

		if len(data) == 0 {
			break
		}
	}
	fmt.Printf("END OF FILE\n")

	destFile.Close()

	ret := client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
	os.Exit(0)
}

