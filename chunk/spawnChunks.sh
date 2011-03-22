#!/bin/bash
################################################################################
# Spawns chunk servers on mumble-01 through mumble-N
#
# If you havent already, make sure that you have sshed in to all mumble machines
# so the RSA fingerprinting has already happened.
################################################################################

pwd=`pwd`
num=0

if [ "$#" != "2" ]
then
    echo "Specify master and the number of computers to infect."
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
	ssh mumble-$num "$pwd/serv $1"&
    done
fi
