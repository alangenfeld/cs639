

package main

import (
	"../client/client"
	"../include/sfs"
	"fmt"
	//"flag"
	"os"
	"strings"
	"flag"
	"io/ioutil"
)

const BUFSIZE = 256

func main() {
	
	fmt.Printf("Hello World\n")
	master := flag.String("m", "mumble-02.cs.wisc.edu", "specify the master")
	flag.Parse();
	
	if *master == "" {
		fmt.Printf("Error, must specify a master.\n");
		os.Exit(1);
	}else{
		fmt.Printf("%s\n", *master);
	}
		
	
	client.Initialize(*master)

	for {
		buffer := make([]byte,BUFSIZE)
		i,err := os.Stdin.Read(buffer[0:BUFSIZE])
		
		line := fmt.Sprintf("%s",buffer[0:i])

		if (err != nil) {
			fmt.Printf("%v", err);
			os.Exit(1);
		}
		
		if i > 0 {
			//fmt.Printf("************************\n")
			//fmt.Printf("%s\n",line)
			//fmt.Printf("1st: %s\n", strings.TrimRight(line," "))
			//fmt.Printf("1st: %d\n", strings.Index(line," "))
			//fmt.Printf("1st: %+v\n", strings.Split(line," ",-1))
			//fmt.Printf("1st: %+v\n", strings.HasPrefix(line,"put"))
		}
		
		if strings.HasPrefix(line,"put") {
			err := put(line)
			if err == false {
				fmt.Printf("Fatal error in put.\n")
			}
			
		}else if strings.HasPrefix(line,"get") {
			err := get(line)
			if err == false {
				fmt.Printf("Fatal error in get.\n")
			}
			
		
		}else if strings.HasPrefix(line,"ls") {
			err := ls(line)
			if err == false {
				fmt.Printf("Fatal error in ls.\n")
			}
		}else if strings.HasPrefix(line,"rmdirg") {
			err := rmdir(line)
			if err == false {
				fmt.Printf("Fatal error in rmdir.\n")
			}
			
		}else if strings.HasPrefix(line,"rm") {
			err := rm(line)
			if err == false {
				fmt.Printf("Fatal error in rm.\n")
			}
		
	
		}else if strings.HasPrefix(line,"pwd") {
			fmt.Printf("pwd not currently supported")
		
		}else if strings.HasPrefix(line,"mkdir") {
			err := mkdir(line)
			if err == false {
				fmt.Printf("Fatal error in mkdir.\n")
			}
		}else if strings.HasPrefix(line,"help") {
			help()
		
		}else if strings.HasPrefix(line,"cd") {
			fmt.Printf("cd not currently supported")	
		
		}else if strings.HasPrefix(line,"logout") {
		
			fmt.Printf("exiting..\n")
			os.Exit(1)
		}else{
		
			fmt.Printf("Command does not exist\n")
		}
		
		
		

	}

}

func help() {

	fmt.Printf("get <src> <dst>\n")
	fmt.Printf("put <src> <dst>\n")
	fmt.Printf("ls <dir>\n")
	fmt.Printf("mkdir <dir>")
	fmt.Printf("rm <file>")
	fmt.Printf("rmdir <file>")
	
	
}

func rm(line string) bool {

	file, err1 := parse1args(line)

	if err1 == false {
		fmt.Printf("Usage: put <src> <dst>\n")
		return false
	}
	
	err := client.Delete(file)

	
	if err == sfs.FAIL {
		return false
	}
	
	return true

}

func rmdir(line string) bool {

	file, err1 := parse1args(line)

	if err1 == false {
		fmt.Printf("Usage: put <src> <dst>\n")
		return false
	}
	
	err := client.RemoveDir(file)

	
	if err == sfs.FAIL {
		return false
	}
	
	return true

}


func mkdir(line string) bool {

	dir, err1 := parse1args(line)
	
	if err1 == false {
		fmt.Printf("Usage: mkdir <dir>\n")
		return false
	}
	fmt.Printf("%s\n", dir);
	err := client.MakeDir(dir)
	
	if err == sfs.FAIL {
		return false
	}
	
	return true

}
func put(line string) bool {
	source, dest, err1 := parse2args(line)
	
	if err1 == false {
		fmt.Printf("Usage: put <src> <dst>\n")
		return false
	}
	fmt.Printf("%s, %s\n", source, dest)
	
	fmt.Printf("end\n")
	f, err := ioutil.ReadFile(source)
	if err != nil {	
		fmt.Printf("Open %s locally failed\n", source)
		fmt.Printf("%v\n",err)
		return false
	}

	fd := client.Open(dest, client.O_WRONLY|client.O_CREATE)
	if(fd < 0) {
		fmt.Printf("Open %s in SFS failed\n", dest)
		return false
	} else {
		fmt.Printf("Open worked!\n")
	}

	fmt.Printf("Writing...\n")
	ret := client.Write(fd, f)
	if(ret != 0) {
		fmt.Printf("Write failed\n")
		return false
	} else {
		fmt.Printf("Write worked!\n")
	}

	ret = client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
			
	return true

}

func get(line string) bool {
	source, dest, err1 := parse2args(line)
	
	if err1 == false {
		fmt.Printf("Usage: get <src> <dst>\n")
		return false
	}
	
	fd := client.Open(source, client.O_RDONLY)
	if(fd < 0) {
		fmt.Printf("could not open file %s", source)
		return false
	}

	destFile, err := os.Open(dest, os.O_CREAT|os.O_WRONLY, 0666)
	if(err != nil) {
		fmt.Printf("Couldn't open %s\n", dest)
	}

	fmt.Printf("BEGIN OF FILE\n")
	for ;true; {
		data, err := client.Read(fd, sfs.CHUNK_SIZE)
		if(err != 0) {
			panic("read failed")
		}

		_, err2 := destFile.Write(data)
		if(err2 != nil) {
			fmt.Printf("Write failed\n")
			return false
		}

		if len(data) == 0 {
			break
		}
	}
	fmt.Printf("END OF FILE\n")

	destFile.Close()

	ret := client.Close(fd)
	if(ret != client.WIN) {
		panic("close failed")
	}
	return true
}

func ls(line string) bool {


	file, err1 := parse1args(line)
	
	if err1 == false {
		fmt.Printf("Usage: ls <file>\n")
		return false
	}

	list, err := client.ReadDir(file)
	if(err != 0) {
		fmt.Printf("\nCouldn't call ReadDir on file %s*\n", file)
	}

	iter := len(list)
	for c := 0; c < iter; c++ {
		fmt.Printf(list[c])
		fmt.Printf("\n")
	}
	return true
}


func parse1args(line string) (string,bool) {

	slice := strings.Index(line," ")
	if slice == -1 {
		return "", false
	}
	
	arg  := fmt.Sprintf("%s",line[slice+1:len(line)-1])
	
	return arg, true

}


func parse2args(line string) (string, string, bool) {

	slice := strings.Index(line," ")
	if slice == -1 {
		return "", "", false
	}
	
	both_args  := fmt.Sprintf("%s",line[slice+1:])
	
	split := strings.Index(both_args," ")
	if split == -1 {
		return "", "", false
	}
	
	arg2 := fmt.Sprintf("%s",both_args[split+1:len(both_args)-1])
	arg1 := fmt.Sprintf("%s",both_args[:split])
	
	return arg1, arg2, true

}


