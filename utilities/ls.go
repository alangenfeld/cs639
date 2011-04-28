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

	list, err := client.ReadDir(*file)
	if(err != 0) {
		fmt.Printf("Couldn't call ReadDir on file %s\n", *file)
	}

	fmt.Printf("BEGIN OF LIST\n")
	iter := len(list)
	for c := 0; c < iter; c++ {
	    fmt.Printf(list[c])
	}
	
	fmt.Printf("\nEND OF LIST\n")

	ret := client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
	os.Exit(0)
}

