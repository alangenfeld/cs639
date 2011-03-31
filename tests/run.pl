#!/usr/bin/perl
use strict;
use Cwd 'abs_path';

my $master = 'mumble-20.cs.wisc.edu';
my $ssh = 'ssh -o StrictHostKeyChecking=no ';
my $chunkCount = 1;
my $testdir;


sub doLaunch {
    sys("ssh $master ".
	"'$testdir/../master/master ".
	"&> $testdir/output/master.log' &");

    sleep(1);

    for(my $i = 3; $i < 3+$chunkCount; $i++) {
	sys("$ssh mumble-1$i.cs.wisc.edu ".
	    "'$testdir/../chunk/serv $master ".
	    "&> $testdir/output/chunk$i.log' &");
    }
}

sub doKill {
    sys("$ssh $master 'killall master'");
    for(my $i = 3; $i < 3+$chunkCount; $i++) {
	sys("$ssh mumble-1$i.cs.wisc.edu 'killall serv'");
    }
}


sub killOne {
    $#_ == 0 or die;
    my ($num) = @_;
    if($num < 10) {
	$num = '0'.$num;
    }
    sys("$ssh mumble-$num.cs.wisc.edu 'killall serv'");
    sys("$ssh mumble-$num.cs.wisc.edu 'killall master'");
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

    print "\n*** Running test $test... ***\n\n";

    #setup output directory
    sys("rm -rf $testdir/output; mkdir $testdir/output");

    #start servers
    doLaunch();
    sleep(1);

    #run client with timeout
    print "Running test $test\n";
    if(fork() == 0) {
	sys("sleep 10; killall $test");
	exit(0);
    }
    sys("$testdir/$test -m $master &> $testdir/output/$test.out");
    print "Done running test $test\n";

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
	return 1;
    } else {
	return 0;
    }
}


main();
sub main {
    $testdir = abs_path($0);
    $testdir =~ s/\/[^\/]*$//;

    my $killNum = '';
    my $testName = '';

    ($#ARGV+1) % 2 == 0 or die "Bad usage";

    for(my $i = 0; $i <= $#ARGV; $i+=2) {
	if($ARGV[$i] eq '-k') {
	    $killNum = $ARGV[$i+1];
	} elsif($ARGV[$i] eq '-t') {
	    $testName = $ARGV[$i+1];
	} else {
	    die "Bad usage";
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
		if($file =~ /^t\d+$/) {
		    if(runTest($file)) {
			push(@passCases, $file);
		    } else {
			push(@failCases, $file);
		    }
		}
	    }
	    closedir $dh;

	    my $passCount = $#passCases + 1;
	    my $failCount = $#failCases + 1;

	    if($passCount + $failCount > 0) {
		printFailSummary(@failCases);

		my $rate = (100 * $passCount / ($passCount + $failCount)) . '%';
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
    print("Running:{\n$cmd\n}\n\n");
    my $ret = system($cmd);
    return $ret;
}
