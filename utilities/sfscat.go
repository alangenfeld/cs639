package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
)

func main(){


	master := flag.String("m", "", "specify a master!")
	file := flag.String("f", "", "specify a file (-f)")

	flag.Parse();

	client.Initialize(*master)

	fd := client.Open(*file, client.O_RDONLY)
	if(fd < 0) {
		fmt.Printf("could not open file %s", *file)
		os.Exit(1)
	}

	fmt.Printf("BEGIN OF FILE\n")
	for ;true; {
		data, err := client.Read(fd, 1024)
		if(err != 0) {
			panic("read failed")
		}
		fmt.Printf("%s", string(data))
		if len(string(data)) == 0 {
			break
		}
	}
	fmt.Printf("END OF FILE\n")

	ret := client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
	os.Exit(0)
}

