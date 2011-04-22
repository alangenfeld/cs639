package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"sort"
)

func createPath(path string) {
	fd := client.Open(path, client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		panic("open failed")
	}
	ret := client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
}

func main(){
	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)


	ls, err := client.ReadDir("/");

	if(err != 0) {
		panic("readdir should work")
	}
	if(len(ls) != 0) {
		panic("readdir returned wrong number of results")
	}


	if(client.MakeDir("/a") != 0) {
		panic("makedir should work")
	}

	fd := client.Open("/b", client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		panic("open failed")
	}

	//do read dir
	ls, err = client.ReadDir("/")
	if(err != 0) {
		panic("readdir failed")
	}
	lsExpected := []string{"a/", "b"}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if(! ArrEquals(ls, lsExpected)) {
		fmt.Printf("Actual:\t%+v\nExpected:\t%+v\n", ls, lsExpected)
		panic("readir results differ")
	}



	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}




func ArrEquals(a []string, b []string) bool { 
	if(len(a) != len(b)) {
		return false
	}
	for i:=0; i < len(a); i++ {
		if(a[i] != b[i]) {
			return false
		}
	}
	return true
}