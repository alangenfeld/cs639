#!/bin/bash

ssh='ssh -o StrictHostKeyChecking=no '

echo "##########################################################"
echo "#           PLEASE RUN CLEANUP.SH WHEN YOURE DONE        #"
echo "##########################################################"
sleep 2

pwd=`pwd`
logout=`mktemp $pwd/stealth-XXXXXX`
$ssh mumble-01 "$pwd/master/master | tee $logout"&
sleep 2
$ssh mumble-40 "$pwd/chunk/serv mumble-01"&
$ssh mumble-39 "$pwd/chunk/serv mumble-01"&
$ssh mumble-38 "$pwd/chunk/serv mumble-01"&
sleep 2
client/test mumble-01


