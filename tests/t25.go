package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"sort"
	"strings"
)

func main() {
	master := flag.String("m", "", "specify a master!")
	flag.Parse()

	client.Initialize(*master)
	var buf []byte
	buf = buf
	var ret int
	ret = ret
	var ls []string
	ls = ls
	var lsExpected []string
	lsExpected = lsExpected
	/*
	TREE:
	/
	/La3LFT/
	/w9FouU/
	/w9FouU/67iAM_/
	/La3LFT/Z4CuPz/
	/La3LFT/pQZOdo/
	/La3LFT/pQZOdo/kL3PUG/
	/La3LFT/pQZOdo/njHo3G/
	/La3LFT/pQZOdo/r4Z4Hb/
	/La3LFT/pQZOdo/kL3PUG/JInFKn/
	/La3LFT/pQZOdo/njHo3G/NSSZkg/
	*/

	ret = client.MakeDir("/La3LFT/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/w9FouU/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/w9FouU/67iAM_/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/La3LFT/Z4CuPz/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/La3LFT/pQZOdo/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/La3LFT/pQZOdo/kL3PUG/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/La3LFT/pQZOdo/njHo3G/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/La3LFT/pQZOdo/r4Z4Hb/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/La3LFT/pQZOdo/kL3PUG/JInFKn/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	ret = client.MakeDir("/La3LFT/pQZOdo/njHo3G/NSSZkg/")
	if ret != 0 {
		panic("MakeDir failed")
	}

	createPath("/rkE42a.txt")

	createPath("/La3LFT/pQZOdo/kL3PUG/gzOmTW.txt")

	createPath("/La3LFT/pQZOdo/kL3PUG/RDA6ux.txt")

	createPath("/La3LFT/jgJiEV.txt")

	createPath("/Mcn_0P.txt")

	createPath("/La3LFT/pQZOdo/njHo3G/NSSZkg/p9LNpw.txt")

	createPath("/La3LFT/pQZOdo/x0fRwt.txt")

	createPath("/La3LFT/Z4CuPz/VdAcKu.txt")

	createPath("/bUl52u.txt")

	createPath("/w9FouU/67iAM_/GxhNt0.txt")

	createPath("/La3LFT/pQZOdo/njHo3G/NSSZkg/bDgIwU.txt")

	createPath("/La3LFT/Z4CuPz/tj5hfp.txt")

	createPath("/La3LFT/pQZOdo/kL3PUG/JInFKn/m4Nm2S.txt")

	createPath("/w9FouU/67iAM_/cR_oUB.txt")

	createPath("/La3LFT/pQZOdo/njHo3G/xoj_h9.txt")

	createPath("/La3LFT/pQZOdo/kL3PUG/JInFKn/xkHUZh.txt")

	createPath("/La3LFT/pQZOdo/njHo3G/w3m6rq.txt")

	createPath("/La3LFT/pQZOdo/kL3PUG/JInFKn/jthx9V.txt")

	createPath("/w9FouU/t0Rc1T.txt")

	createPath("/La3LFT/pQZOdo/r4Z4Hb/gQ_aYr.txt")

	lsExpected = []string{"La3LFT", "w9FouU", "rkE42a.txt", "Mcn_0P.txt", "bUl52u.txt"}

	ls, ret = client.ReadDir("/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"pQZOdo", "Z4CuPz", "jgJiEV.txt"}

	ls, ret = client.ReadDir("/La3LFT/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"67iAM_", "t0Rc1T.txt"}

	ls, ret = client.ReadDir("/w9FouU/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"GxhNt0.txt", "cR_oUB.txt"}

	ls, ret = client.ReadDir("/w9FouU/67iAM_/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"VdAcKu.txt", "tj5hfp.txt"}

	ls, ret = client.ReadDir("/La3LFT/Z4CuPz/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"r4Z4Hb", "kL3PUG", "njHo3G", "x0fRwt.txt"}

	ls, ret = client.ReadDir("/La3LFT/pQZOdo/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"JInFKn", "gzOmTW.txt", "RDA6ux.txt"}

	ls, ret = client.ReadDir("/La3LFT/pQZOdo/kL3PUG/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"NSSZkg", "xoj_h9.txt", "w3m6rq.txt"}

	ls, ret = client.ReadDir("/La3LFT/pQZOdo/njHo3G/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"gQ_aYr.txt"}

	ls, ret = client.ReadDir("/La3LFT/pQZOdo/r4Z4Hb/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"m4Nm2S.txt", "xkHUZh.txt", "jthx9V.txt"}

	ls, ret = client.ReadDir("/La3LFT/pQZOdo/kL3PUG/JInFKn/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	lsExpected = []string{"p9LNpw.txt", "bDgIwU.txt"}

	ls, ret = client.ReadDir("/La3LFT/pQZOdo/njHo3G/NSSZkg/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	sort.SortStrings(lsExpected)
	sort.SortStrings(ls)
	if !ArrEquals(ls, lsExpected) {
		fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls)
		panic("Unexpected ReadDir results")
	}

	recursiveDelete("/")
	ls, ret = client.ReadDir("/")
	if ret != 0 {
		panic("ReadDir failed")
	}
	if len(ls) != 0 {
		panic("after a recursive delete, everything should be gone")
	}

	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}


func ArrEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}


func createPath(path string) {
	fd := client.Open(path, client.O_WRONLY|client.O_CREATE)
	if fd < 0 {
		panic("open failed")
	}
	ret := client.Close(fd)
	if ret != client.WIN {
		panic("close failed")
	}
}


func recursiveDelete(path string) {
	fmt.Printf("Recursive delete on %s\n", path)

	ls, ret := client.ReadDir(path)
	if ret != 0 {
		panic("ReadDir failed")
	}
	if len(ls) == 0 {
		if path != "/" {
			ret = client.RemoveDir(path)
			if ret != 0 {
				panic("RemoveDir failed")
			}
		}
		return
	}

	//try deleting it non-empty
	if client.RemoveDir(path) == 0 {
		panic("it let us delete a non-empty directory!")
	}

	//we have to recursive delete stuff
	for i := 0; i < len(ls); i++ {
		if strings.HasSuffix(ls[i], "/") {
			//its a dir, do recursive
			recursiveDelete(path + ls[i])
		} else {
			//its a file, delete it
			if client.Delete(path+ls[i]) != 0 {
				fmt.Printf("Try to delete %s\n", path+ls[i])
				panic("Couldnt delete file")
			}
		}
	}

	//try deleting it empty
	if path != "/" {
		if client.RemoveDir(path) != 0 {
			panic("couldnt delete directory!")
		}
	}
}
