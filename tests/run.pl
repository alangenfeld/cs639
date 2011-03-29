#!/usr/bin/perl
use strict;
use Cwd 'abs_path';

my $master = 'mumble-02.cs.wisc.edu';
my $ssh = 'ssh -o StrictHostKeyChecking=no ';
my $chunkCount = 1;
my $testdir;


sub doLaunch {
    sys("ssh $master -o StrictHostKeyChecking=no $testdir/../master/master &");

    for(my $i = 3; $i < 3+$chunkCount; $i++) {
	sys("$ssh mumble-1$i.cs.wisc.edu $testdir/../chunk/serv $master &");
    }
}

sub doKill {
    sys("$ssh $master 'killall master'");
    for(my $i = 3; $i < 3+$chunkCount; $i++) {
	sys("$ssh mumble-1$i.cs.wisc.edu 'killall serv'");
    }
}

main();
sub main {
    $testdir = abs_path($0);
    $testdir =~ s/\/[^\/]*$//;
    print "Test dir: ".$testdir."\n";
    doLaunch();
    sleep(2);

    system("$testdir/t1");

    sleep(2);
    doKill();
}



sub sys {
    my ($cmd) = @_;
    print("Running $cmd\n");
    my $ret = system($cmd);
    return $ret;
}
