#!/bin/bash
################################################################################
# Spawns chunk servers on mumble-01 through mumble-N
################################################################################

pwd=`pwd`
num=0
ssh='ssh -o StrictHostKeyChecking=no '

if [ "$#" != "2" ]
then
    echo "Specify master and the number of computers to infect."
    echo "example: spawnChunks hostname numServers"
else
    if [ $2 -gt 40 ]
    then
	echo "There are only 40 mumbles"
	max=40
    else
	max=$2
    fi

    for (( c=1;c<=$max;c++ ))
    do
	if [ $c -lt 10 ]
	then
	    num="0$c"
	else
	    num="$c"
	fi
	$ssh mumble-$num "$pwd/serv $1"&
    done
fi
