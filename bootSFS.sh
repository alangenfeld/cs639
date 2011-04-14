#!/bin/bash
################################################################################
# Spawns chunk servers on mumble-01 through mumble-N
################################################################################

pwd=`pwd`
num=0
ssh='ssh -o StrictHostKeyChecking=no '

trap cleanup 2

cleanup()
{
    $ssh mumble-01 "killall master" &
    for (( c=1;c<=40;c++ ))
    do
	if [ $c -lt 10 ]
	then
	    num="0$c"
	else
	    num="$c"
	fi
	$ssh mumble-$num "killall serv" &
    done
    exit
}

if [ "$#" != "1" ]
then
    echo "Specify the number of computers to infect."
    echo "example: spawnChunks hostname numServers"
else
    if [ $1 -gt 40 ]
    then
	echo "There are only 40 mumbles"
	max=40
    else
	max=$1
    fi

    $ssh mumble-01 "$pwd/master/master"&

    for (( c=1;c<=$max;c++ ))
    do
	if [ $c -lt 10 ]
	then
	    num="0$c"
	else
	    num="$c"
	fi
	$ssh mumble-$num "$pwd/chunk/serv mumble-01" &
    done

    while :
    do
	sleep 1
    done
fi    
