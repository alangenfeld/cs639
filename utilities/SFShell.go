

package main

import (
	"fmt"
)

func main() {
	
	fmt.Printf("Hello World\n");


	for {
		buffer := new(string)
		i,err := fmt.Scan(buffer)

		if err != nil {
			fmt.Printf("%v\n", err)
		}
		
		if i > 0 {
			fmt.Printf("%s\n",* buffer)
		}
		

	}

}
