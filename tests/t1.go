package main

import (
	"../client/client"
	"log"
	"os"
	"fmt"
)

func main(){
	if len(os.Args) < 1 {
		log.Printf("Error: Run as \"client [masters ip]\"");
	}
	
	fmt.Printf("Hello")

	ret := client.Open("newfile.txt", true, "mumble-02", 0)
	fmt.Print("Return is %d\n",ret);
}

