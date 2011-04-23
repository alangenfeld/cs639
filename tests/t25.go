
package main

import (
	"../client/client"
	"fmt"
	"flag"
	"os"
	"sort"
	"strings"
)

func main(){
	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)
	var buf []byte
	buf = buf
	var ret int
	ret = ret
	var ls []string; ls = ls
	var lsExpected []string; lsExpected = lsExpected
/*
TREE:
/
/HKBFBe/
/YisZFh/
/HKBFBe/xA9o_c/
/HKBFBe/9O1cEa/
/YisZFh/TWciId/
/YisZFh/fmCxop/
/YisZFh/GRro57/
/YisZFh/J_PhjQ/
/YisZFh/_wIEZz/
/YisZFh/_wIEZz/lUGH0D/
*/

ret = client.MakeDir("/HKBFBe/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/YisZFh/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/HKBFBe/xA9o_c/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/HKBFBe/9O1cEa/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/YisZFh/TWciId/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/YisZFh/fmCxop/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/YisZFh/GRro57/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/YisZFh/J_PhjQ/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/YisZFh/_wIEZz/")
if(ret != 0) {
  panic("MakeDir failed")
}

ret = client.MakeDir("/YisZFh/_wIEZz/lUGH0D/")
if(ret != 0) {
  panic("MakeDir failed")
}

createPath("/YisZFh/_wIEZz/Gj46DI.txt")

createPath("/HKBFBe/9O1cEa/bTg0uN.txt")

createPath("/1mf6P2.txt")

createPath("/YisZFh/TWciId/6CBF7G.txt")

createPath("/YisZFh/J_PhjQ/0s4gx2.txt")

createPath("/HKBFBe/8aBbjb.txt")

createPath("/YisZFh/fmCxop/pHoiDi.txt")

createPath("/54EkC7.txt")

createPath("/HKBFBe/_S_mGz.txt")

createPath("/YisZFh/H2OC3g.txt")

createPath("/YisZFh/GRro57/ReMWv5.txt")

createPath("/HKBFBe/4eYB6n.txt")

createPath("/HKBFBe/xA9o_c/wz8XZV.txt")

createPath("/YisZFh/fmCxop/HzdYke.txt")

createPath("/HKBFBe/xA9o_c/hTgTyW.txt")

createPath("/YisZFh/TWciId/_wfEZr.txt")

createPath("/YisZFh/_wIEZz/Yomrr6.txt")

createPath("/HKBFBe/xA9o_c/cRxezL.txt")

createPath("/YisZFh/J_PhjQ/QI0UTJ.txt")

createPath("/YisZFh/fmCxop/wwSkkm.txt")

lsExpected = []string{"YisZFh/","HKBFBe/","1mf6P2.txt","54EkC7.txt"}

ls, ret = client.ReadDir("/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"xA9o_c/","9O1cEa/","8aBbjb.txt","_S_mGz.txt","4eYB6n.txt"}

ls, ret = client.ReadDir("/HKBFBe/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"_wIEZz/","GRro57/","fmCxop/","TWciId/","J_PhjQ/","H2OC3g.txt"}

ls, ret = client.ReadDir("/YisZFh/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"wz8XZV.txt","hTgTyW.txt","cRxezL.txt"}

ls, ret = client.ReadDir("/HKBFBe/xA9o_c/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"bTg0uN.txt"}

ls, ret = client.ReadDir("/HKBFBe/9O1cEa/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"6CBF7G.txt","_wfEZr.txt"}

ls, ret = client.ReadDir("/YisZFh/TWciId/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"pHoiDi.txt","HzdYke.txt","wwSkkm.txt"}

ls, ret = client.ReadDir("/YisZFh/fmCxop/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"ReMWv5.txt"}

ls, ret = client.ReadDir("/YisZFh/GRro57/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"0s4gx2.txt","QI0UTJ.txt"}

ls, ret = client.ReadDir("/YisZFh/J_PhjQ/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{"lUGH0D/","Gj46DI.txt","Yomrr6.txt"}

ls, ret = client.ReadDir("/YisZFh/_wIEZz/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

lsExpected = []string{}

ls, ret = client.ReadDir("/YisZFh/_wIEZz/lUGH0D/")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}

recursiveDelete("/")
ls, ret = client.ReadDir("/")
if(ret != 0) {
  panic("ReadDir failed")
}
if(len(ls) != 0) {
  panic("after a recursive delete, everything should be gone")
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



func recursiveDelete(path string) {
	fmt.Printf("Recursive delete on %s\n", path)

	ls, ret := client.ReadDir(path)
	if(ret != 0) {
		panic("ReadDir failed")
	}
	if(len(ls) == 0) {
		if(path != "/") {
			ret = client.RemoveDir(path);
			if(ret != 0) {
				panic("RemoveDir failed")
			}
		}
		return
	}

	//try deleting it non-empty
	if(client.RemoveDir(path) == 0) {
		panic("it let us delete a non-empty directory!")
	}

	//we have to recursive delete stuff
	for i:=0; i<len(ls); i++ {
		if(strings.HasSuffix(ls[i],"/")) {
			//its a dir, do recursive
			recursiveDelete(path+ls[i])
		} else {
			//its a file, delete it
			if(client.Delete(path+ls[i]) != 0) {
				fmt.Printf("Try to delete %s\n", path+ls[i])
				panic("Couldnt delete file");
			}
		}
	}

	//try deleting it empty
	if(path != "/") {
		if(client.RemoveDir(path) != 0) {
			panic("couldnt delete directory!")
		}
	}
}
