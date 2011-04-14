#!/usr/bin/perl

my @chars=('a'..'z','A'..'Z','0'..'9','_');

#key = path name, val = file contents
my %files;
#key = fd variable name index, val = path name
my %fds;
#key = fd variable name index, val = offset
my %fdOffsets;

my $nextFD = 1;



sub doOpen {
    my $fileName;
    if(int(rand(2)) == 0 and (keys %files) > 0) {
	$fileName = (keys %files)[int(rand(keys %files))];
    } else {
	$fileName = "/".randStr(10).".txt";
	$files{$fileName} = '';
    }

    print '
fd'.$nextFD.' := client.Open("'.$fileName.'", client.O_RDWR|client.O_CREATE)
if(fd'.$nextFD.' < 0) {
  panic("open failed")
}

';

    $fds{$nextFD} = $fileName;
    $fdOffsets{$nextFD} = 0;
    $nextFD++;
}



sub doRead {
    my $fd = (keys %fds)[int(rand(keys %fds))];
    my $length = int(rand(100));

    print '
buf, ret = client.Read(fd'.$fd.', '.$length.')
if(ret != 0) {
  panic("read failed")
}
if(string(buf) != "'.substr($files{$fds{$fd}}, $fdOffsets{$fd}, $length).'") {
  panic("wrong data returned")
}

';

    $fdOffsets{$fd} += $length;
    if($fdOffsets{$fd} > length($files{$fds{$fd}})) {
	$fdOffsets{$fd} = length($files{$fds{$fd}})
    }
}


sub doWrite {
    my $fd = (keys %fds)[int(rand(keys %fds))];
    my $length = int(rand(100));
    my $str = randStr($length);

    print "//fd state: ($fdOffsets{$fd}) $files{$fds{$fd}}\n";

    print '
ret = client.Write(fd'.$fd.', []byte("'.$str.'"))
if(ret != 0) {
  panic("write failed")
}

';

    $files{$fds{$fd}} = 
	substr($files{$fds{$fd}}, 0, $fdOffsets{$fd}) . 
	$str . 
	substr($files{$fds{$fd}}, $fdOffsets{$fd} + $length);

    $fdOffsets{$fd} += $length;
    if($fdOffsets{$fd} > length($files{$fds{$fd}})) {
	$fdOffsets{$fd} = length($files{$fds{$fd}})
    }

    print "//fd state: ($fdOffsets{$fd}) $files{$fds{$fd}}\n";
}


sub doSeek {
    my $fd = (keys %fds)[int(rand(keys %fds))];
    my $offset = int(rand(length($files{$fds{$fd}})+1));

    print '
ret = client.Seek(fd'.$fd.', '.$offset.', client.SEEK_SET)
if(ret != '.$offset.') {
  fmt.Printf("Seek returned %d, but we expected %d\n", ret, '.$offset.')
  panic("seek failed")
}

';    

    $fdOffsets{$fd} = $offset;
}



sub doClose {
    my $fd = (keys %fds)[int(rand(keys %fds))];
    delete $fds{$fd};
    delete $fdOffset{$fd};

    print '
ret = client.Close(fd'.$fd.')
if(ret != 0) {
  panic("close failed")
}

';    
}


main();
sub main {
    my $instCount = 5000;

    printPrefix();

    for(my $i = 0; $i < $maxFDCount; $i++) {
	print "\tvar fd$i int\n"
    }

    for(my $i = 0; $i < $instCount; $i++) {
	my $type = int(rand(5));

	if($type == 0 or (keys %fds) == 0) {
	    #open
	    doOpen();
	} elsif ($type == 1) {
	    #read
	    doRead();
	} elsif ($type == 2) {
	    #write
	    doWrite();
	} elsif ($type == 3) {
	    #seek
	    doSeek();
	} elsif ($type == 4) {
	    doClose();
	}
    }

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
)

func main(){
	master := flag.String("m", "", "specify a master!")
	flag.Parse();

	client.Initialize(*master)
	var buf []byte
	buf = buf
	var ret int
	ret = ret
';
}

sub printSuffix {
    print '
	fmt.Printf("\n{{{{{pass}}}}}\n")
	os.Exit(0)
}
';
}

