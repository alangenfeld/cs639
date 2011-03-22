#!/bin/bash

for (( c=1;c<=40;c++ ))
do
    if [ $c -lt 10 ]
    then
	num="0$c"
    else
	num="$c"
    fi
    ssh mumble-$num "exit"
done
