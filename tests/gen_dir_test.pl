#!/usr/bin/perl

use strict;

my @chars=('a'..'z','A'..'Z','0'..'9','_');

my %dirs=(
    "/" => "",
);




sub goMakeDir {
    $#_ == 0 or die;
    my ($path) = @_;

    print '
ret = client.MakeDir("'.$path.'")
if(ret != 0) {
  panic("MakeDir failed")
}
'
}



sub goReadDir {
    $#_ == 0 or die;
    my ($path) = @_;

    if($dirs{$path} eq '') {
	print '
lsExpected = []string{}
';
    } else {
	my @lsArr = split(':',$dirs{$path});
	for(my $i = 0; $i <= $#lsArr; $i++) {
		if($lsArr[$i] =~ /.txt$/){
			$lsArr[$i] = '"'.$lsArr[$i].'"';
		}
		else {
			$lsArr[$i] = '"'.$lsArr[$i].'/"';
		}
	    	
	}
	print '
lsExpected = []string{'.join(',',@lsArr).'}
';
    }

    print '
ls, ret = client.ReadDir("'.$path.'")
if(ret != 0) {
  panic("ReadDir failed")
}
sort.SortStrings(lsExpected)
sort.SortStrings(ls)
if(! ArrEquals(ls, lsExpected)) {
  fmt.Printf("For ReadDir we expected %v but got %v\n", lsExpected, ls);
  panic("Unexpected ReadDir results")
}
'
}




main();
sub main {
    my $pathCount = 3000;

    printPrefix();

    #randomly generate dirs
    for(my $i = 0; $i < $pathCount/3; $i++) {
	my $parent = (keys %dirs)[int(rand(keys %dirs))];
	my $newdir = randStr(6);
	my $dirpath = $parent.$newdir.'/';

	if($dirs{$parent} eq '') {
	    $dirs{$parent} = $newdir;
	} else {
	    $dirs{$parent} .= ':'.$newdir;
	}
	$dirs{$dirpath} = '';
    }


    #get array of all directories
    my @dirarr = sort {length($a) <=> length($b)} keys %dirs;
    print "/*\nTREE:\n";
    print join("\n", @dirarr)."\n";
    print "*/\n";



    #print go code that creates directories
    foreach my $d (@dirarr) {
	if($d ne "/") {
	    goMakeDir($d);
	}
    }



    #randomly generate files
    for(my $i = 0; $i < 2*$pathCount/3; $i++) {
	my $parent = (keys %dirs)[int(rand(keys %dirs))];
	my $newfile = randStr(6).".txt";
	my $filepath = $parent.$newfile;

	if($dirs{$parent} eq '') {
	    $dirs{$parent} = $newfile;
	} else {
	    $dirs{$parent} .= ':'.$newfile;
	}

	print '
createPath("'.$filepath.'")
';
    }


    #print go code that reads directories
    foreach my $d (@dirarr) {
	goReadDir($d);
    }

    #print code that recursively deletes everything
    print '
recursiveDelete("/")
ls, ret = client.ReadDir("/")
if(ret != 0) {
  panic("ReadDir failed")
}
if(len(ls) != 0) {
  panic("after a recursive delete, everything should be gone")
}
';

    printSuffix();
}

sub randStr {
    $#_ == 0 or die;
    my ($length) = @_;
    my $str = '';
    for(my $i = 0; $i < $length; $i++) {
	$str .= $chars[rand @chars];
    }
    return $str;
}




sub printPrefix {
    print '
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
';
}

sub printSuffix {
    print '
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
';
}

