package main

import (
	"../client/client"
	"fmt"
	"flag"
)

func main(){
	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)

	fd := client.Open("newfile.txt", client.O_WRONLY|client.O_CREATE)

	fmt.Print("Return is %d\n",fd);
}

