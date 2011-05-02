#!/usr/bin/perl
use strict;
use Cwd 'abs_path';

my $mumbleBase = 20;
my $ssh = 'ssh -o StrictHostKeyChecking=no ';
my $chunkCount = 3;
my $timeout = 120;
my $testdir;
my $verbose = 0;
my $master;
my $divide = 1;
my @servers;
my $outputDir = 'output';

sub doLaunch {
    sys("ssh $master ".
	"'$testdir/../master/master ".
	"&> $testdir/$outputDir/master.log' ".($verbose != 1 ? " &> /dev/null":"")." &");

    sleep(1);

    for(my $i = 1; $i <= $chunkCount; $i++) {
	sys("$ssh $servers[$i] ".
	    "'$testdir/../chunk/serv $master ".
	    "&> $testdir/$outputDir/chunk$i.log' ".($verbose != 1 ? " &> /dev/null":"")." &");
    }
}

sub doKill {
    sys("$ssh $master 'killall master'".($verbose != 1 ? " &> /dev/null":""));
    for(my $i = 1; $i <= $chunkCount; $i++) {
	sys("$ssh $servers[$i] 'killall serv'".($verbose != 1 ? " &> /dev/null":""));
	sys("$ssh $servers[$i] 'killall master'".($verbose != 1 ? " &> /dev/null":""));
    }
}


sub killOne {
    $#_ == 0 or die;
    my ($num) = @_;
    if($num < 10) {
	$num = '0'.$num;
    }
    sys("$ssh mumble-$num.cs.wisc.edu 'killall serv'".($verbose != 1 ? " &> /dev/null":""));
    sys("$ssh mumble-$num.cs.wisc.edu 'killall master'".($verbose != 1 ? " &> /dev/null":""));
}



sub superKill {
    for(my $i = 1; $i < 40; $i++) {
	killOne($i);
    }
}


sub printFailSummary {
    my (@failList) = @_;
    my %descriptions;

    if($#failList == -1) {
	return;
    }

    open SUMMARY, "$testdir/summary.txt" or die;
    while(<SUMMARY>) {
	if($_ =~ /(.*):(.*)/) {
	    $descriptions{$1} = $2;
	}
    }
    close SUMMARY;

    print "\n\nFailures:\n";
    foreach my $name (@failList) {
	print "$name: $descriptions{$name}\n";
    }
}


sub runTest {
    $#_ == 0 or die "Bad args to runTest";
    my ($test) = @_;

    print "Compiling $test...\n";
    sys("rm -f $test; 6g $test.go; 6l -o $test $test.6");
    if(! -e "$test") {
	print "Compile $test failed!\n";
	return 0;
    }

    print "Running $test...\n";

    #setup output directory
    sys("rm -rf $testdir/$outputDir; mkdir $testdir/$outputDir");

    #start servers
    doLaunch();
    sleep(1);

    #start timeout killer
    my $pid1 = fork();
    if($pid1 == 0) {
	sleep($timeout);
	print "TIMEOUT!!!";
	sys("killall $test".($verbose != 1 ? " &> /dev/null":""));
	exit(0);
    }


    #start chaos proc (chunk killer)
    my $pid2 = -1;
    if(killTest1($test)) {
	$pid2 = fork();
	if($pid2 == 0) {
	    while(1) {
		sleep(10);
		sys("$ssh $servers[1] 'killall serv'".($verbose != 1 ? " &> /dev/null":""));
		sleep(10);
		sys("$ssh $servers[1] ".
		    "'$testdir/../chunk/serv $master ".
		    "&> $testdir/$outputDir/chunk1.log' ".($verbose != 1 ? " &> /dev/null":"")." &");
	    }
	    exit(0);
	}
    }




    #start chaos proc (other chunk killer)
    my $pid3 = -1;
    if(killTest2($test)) {
	$pid3 = fork();
	if($pid3 == 0) {
	    my $i = 0;
	    while(1) {
		$i++;
		if($i == 0 or $i > $#servers) {
		    $i == 1;
		}
		sleep(2);
		print "Killing $i\n";
		sys("$ssh $servers[$i] 'killall serv'".($verbose != 1 ? " &> /dev/null":""));
		sleep(8);
		print "Restarting $i\n";
		sys("$ssh $servers[$i] ".
		    "'$testdir/../chunk/serv $master ".
		    "&> $testdir/$outputDir/chunk$i.log' ".($verbose != 1 ? " &> /dev/null":"")." &");
	    }
	    exit(0);
	}
    }




    sys("$testdir/$test -m $master &> $testdir/$outputDir/$test.out");
    print "Done running $test\n";

    #kill the child that was supposed to kill us! (for the timeout)
    kill 9, $pid1;
    if($pid2 > 0) {
	kill 9, $pid2;
    }
    if($pid3 > 0) {
	kill 9, $pid3;
    }

    #stop servers
    sleep(1);
    doKill();

    #check if output is success
    my $results = '';
    open RESULT, "$testdir/$outputDir/$test.out" or die "Could not read result: $!";
    while(<RESULT>) {
	$results .= $_;
    }
    close RESULT;
    if($results =~ "{{{{{pass}}}}}") {
	print "$test result = passed\n";
	return 1;
    } else {
	print "$test result = failed\n";
	return 0;
    }
}


sub killTest1 {
    $#_ == 0 or die;
    my ($test) = @_;
    my $ret = 0;
    $test =~ s/t//;

    open KT1, "killtest1.txt" or die "Could not open file: $!\n";
    while(my $line = <KT1>) {
	chomp($line);
	$line =~ s/t//;
	if($line eq $test) {
	    $ret = 1;
	}
    }
    close KT1;
    return $ret;
}



sub killTest2 {
    $#_ == 0 or die;
    my ($test) = @_;
    my $ret = 0;
    $test =~ s/t//;

    open KT2, "killtest2.txt" or die "Could not open file: $!\n";
    while(my $line = <KT2>) {
	chomp($line);
	$line =~ s/t//;
	if($line eq $test) {
	    $ret = 2;
	}
    }
    close KT2;
    return $ret;
}


sub usage {
    die "Usage:\n./run.pl \n\t[-k (all|mumbleIndex)] \n\t[-t (all|testName)] \n\t[-c chunkCount] \n\t[-v (1|0)] \n\t[-s timeoutSeconds]\n\t[-b mumbleBase]\n\t[-d divide]\n ";
}


main();
sub main {
    $testdir = abs_path($0);
    $testdir =~ s/\/[^\/]*$//;

    my $killNum = '';
    my @testArr;
    my $testName = '';

    #grab test parameters
    (($#ARGV+1) % 2 == 0 and $#ARGV >= 0) or usage();
    for(my $i = 0; $i <= $#ARGV; $i+=2) {
	if($ARGV[$i] eq '-k') {
	    $killNum = $ARGV[$i+1];
	} elsif($ARGV[$i] eq '-t') {
	    $testName = $ARGV[$i+1];
	} elsif($ARGV[$i] eq '-c') {
	    $chunkCount = $ARGV[$i+1];
	} elsif($ARGV[$i] eq '-v') {
	    $verbose = $ARGV[$i+1];
	} elsif($ARGV[$i] eq '-s') {
	    $timeout = $ARGV[$i+1];
	} elsif($ARGV[$i] eq '-b') {
	    $mumbleBase = $ARGV[$i+1];
	} elsif($ARGV[$i] eq '-d') {
	    $divide = $ARGV[$i+1];
	} else {
	    usage();
	}
    }


    #does this do nothing?
    if($testName eq '' and $killNum eq '') {
	die ("No test num or kill num\n");
    }


    #kill specified servers
    if($killNum ne '') {
	if($killNum eq 'all') {
	    superKill();
	} else {
	    print "Kill $killNum\n";
	    killOne($killNum);
	}
    }
    if($testName eq '') {
	exit(0);
    }


    $testName =~ s/t//gi;
    my $testLowBound = 0;
    my $testHighBound = 1000000;

    if($testName =~ /^(\d+)$/) {
	$testLowBound = $testHighBound = $1;
    } elsif($testName =~ /^(\d+)\-(\d+)$/) {
	$testLowBound = $1;
	$testHighBound = $2;
    } elsif($testName eq '') {
	$testHighBound = -1;
    }

    #get list of tests
    opendir(my $dh, $testdir) || 
	die "can't opendir $testdir: $!";
    while(my $file = readdir($dh)) {
	if($file =~ /^t(\d+)\.go$/) {
	    my $index = $1;
	    if($index >= $testLowBound && $index <= $testHighBound) {
		push(@testArr, "t$index");
	    }
	}
    }
    closedir $dh;



    #fork to divide tests among processes
    sys("rm -rf $testdir/resultDir; mkdir $testdir/resultDir");
    my @children;
    my $divSize = ($#testArr+1)/$divide;
    for(my $i = 0; $i < $divide; $i++) {
	my @testGroup;
	for(my $j = 0; $j < $divSize and $#testArr>=0; $j++) {
	    push(@testGroup, shift(@testArr));
	}

	my $pid = fork();
	if($pid == 0) {
	    $outputDir = "output$i";
	    @testArr = @testGroup;
	    $mumbleBase += ($chunkCount+1)*$i;
	    last; #run this division
	} #else

	push(@children, $pid);

	if ($i == $divide-1) {
	    foreach my $child (@children) {
		waitpid($child,0);
	    }

	    #all children are done running!  aggregate
	    #over their results

	    my @passCases;
	    my @failCases;
	    opendir(my $resultDir, "$testdir/resultDir") || 
		die "can't opendir $testdir/resultDir: $!";
	    while(my $file = readdir($resultDir)) {
		if($file =~ /(t\d+)pass/) {
		    push(@passCases, $1);
		} elsif($file =~ /(t\d+)fail/) {
		    push(@failCases, $1);
		}
	    }

	    my $passCount = $#passCases + 1;
	    my $failCount = $#failCases + 1;

	    if($passCount + $failCount > 0) {
		printFailSummary(@failCases);

		my $rate = int(100 * $passCount / ($passCount + $failCount)) . '%';
		print "\n\nPass=$passCount, Fail=$failCount, Rate=$rate\n";
		if($passCount+$failCount >= 1) {
		    if($failCount == 0) {
			print "T-shirt time!";
		    } else {
			print "Close, but no t-shirt!";
		    }
		} else {
		    print "No tests to run...\n";
		}
	    }
	    
	    print "\n\n\n";

	    exit(0);
	}
    }



    #generate list of servers
    for(my $i = 0; $i < $chunkCount+1; $i++) {
	my $index = $i + $mumbleBase;
	$index = ($index < 10) ? "0$index" : $index;
	push(@servers, "mumble-$index.cs.wisc.edu");
    }
    $master = $servers[0];



    #run the test
    if($testName ne '') {
	my @passCases;
	my @failCases;

	foreach my $name (@testArr) {
	    if(runTest($name)) {
		sys("touch $testdir/resultDir/$name".'pass');
		push(@passCases, $name);
	    } else {
		sys("touch $testdir/resultDir/$name".'fail');
		push(@failCases, $name);
	    }
	}
    }
}


sub sys {
    my ($cmd) = @_;
    if($verbose == 1) {
	print("Running:{\n$cmd\n}\n\n");
    }
    my $ret = system($cmd);
    return $ret;
}
