#!/usr/bin/perl
use strict;
use Cwd 'abs_path';

my $master = 'mumble-20.cs.wisc.edu';
my $ssh = 'ssh -o StrictHostKeyChecking=no ';
my $chunkCount = 1;
my $timeout = 10;
my $testdir;
my $verbose = 1;


sub doLaunch {
    sys("ssh $master ".
	"'$testdir/../master/master ".
	"&> $testdir/output/master.log' ".($verbose != 1 ? " &> /dev/null":"")." &");

    sleep(1);

    for(my $i = 3; $i < 3+$chunkCount; $i++) {
	sys("$ssh mumble-1$i.cs.wisc.edu ".
	    "'$testdir/../chunk/serv $master ".
	    "&> $testdir/output/chunk$i.log' ".($verbose != 1 ? " &> /dev/null":"")." &");
    }
}

sub doKill {
    sys("$ssh $master 'killall master'".($verbose != 1 ? " &> /dev/null":""));
    for(my $i = 3; $i < 3+$chunkCount; $i++) {
	killOne($i)
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

    print "\n*** $test... ***\n";

    print "Compiling...\n";
    sys("rm -f $test; 6g $test.go; 6l -o $test $test.6");
    if(! -e "$test") {
	print "Compile failed!\n";
	return 0;
    }

    print "Running...\n";

    #setup output directory
    sys("rm -rf $testdir/output; mkdir $testdir/output");

    #start servers
    doLaunch();
    sleep(1);

    #run client with timeout
    my $pid = fork();
    if($pid == 0) {
	sleep($timeout);
	sys("killall $test".($verbose != 1 ? " &> /dev/null":""));
	exit(0);
    }
    sys("$testdir/$test -m $master &> $testdir/output/$test.out");
    print "Done running\n";

    #kill the child that was supposed to kill us! (for the timeout)
    kill 9, $pid;

    #stop servers
    sleep(1);
    doKill();

    #check if output is success
    my $results = '';
    open RESULT, "$testdir/output/$test.out" or die "Could not read result: $!";
    while(<RESULT>) {
	$results .= $_;
    }
    close RESULT;
    if($results =~ "{{{{{pass}}}}}") {
	print "Result = passed\n";
	return 1;
    } else {
	print "Result = failed\n";
	return 0;
    }
}


sub usage {
    die "Usage:\n./run.pl \n\t[-k (all|mumbleIndex)] \n\t[-t (all|testName)] \n\t[-c chunkCount] \n\t[-v (1|0)] \n\t[-s timeoutSeconds]\n ";
}


main();
sub main {
    $testdir = abs_path($0);
    $testdir =~ s/\/[^\/]*$//;

    my $killNum = '';
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
	} else {
	    usage();
	}
    }

    #does this do nothing
    if($testName eq '' and $killNum eq '') {
	die ("No test num or kill num\n");
    }


    #administrative job, not testing
    if($killNum ne '') {
	if($killNum eq 'all') {
	    superKill();
	} else {
	    print "Kill $killNum\n";
	    killOne($killNum);
	}
    }


    #run the test
    if($testName ne '') {
	my @passCases;
	my @failCases;

	if($testName eq 'all') {
	    opendir(my $dh, $testdir) || 
		die "can't opendir $testdir: $!";
	    while(my $file = readdir($dh)) {
		if($file =~ /^(t\d+)\.go$/) {
		    my $name = $1;
		    if(runTest($name)) {
			push(@passCases, $name);
		    } else {
			push(@failCases, $name);
		    }
		}
	    }
	    closedir $dh;

	    my $passCount = $#passCases + 1;
	    my $failCount = $#failCases + 1;

	    if($passCount + $failCount > 0) {
		printFailSummary(@failCases);

		my $rate = int(100 * $passCount / ($passCount + $failCount)) . '%';
		print "\n\nPass=$passCount, Fail=$failCount, Rate=$rate\n";
		if($failCount == 0) {
		    print "T-shirt time!";
		} else {
		    print "Close, but no t-shirt!";
		}
	    } else {
		print "No tests to run...\n";
	    }
	} else {
	    if(runTest($testName)) {
		print "\n\nPASS\n";
	    } else {
		print "\n\nFAIL\n";
	    }
	}
	print "\n\n\n";
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
