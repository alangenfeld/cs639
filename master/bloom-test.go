package main

import (
	"./bloom"
	"fmt"
)


func main(){
	tstring := "Hello world"

	b := bloom.New()
	
	if b.Test(tstring){
		fmt.Printf("success!\n")
	}
	
	b.Set(tstring)
	
	if b.Test(tstring){
		fmt.Printf("success!\n")
	}
	
	tstring = "Hello world!"
	b.Set(tstring)
	
	if b.Test(tstring){
		fmt.Printf("success!\n")
	}
}
