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


	list, err := client.ReadDir(*file)
	if(err != 0) {
		fmt.Printf("\nCouldn't call ReadDir on file %s\n", *file)
	}

	iter := len(list)
	for c := 0; c < iter; c++ {
	    fmt.Printf(list[c])
	    fmt.Printf("\n")
	}
	


	os.Exit(0)
}

