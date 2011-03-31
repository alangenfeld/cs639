#!/bin/bash

./master/master&
sleep 2
./chunk/serv localhost&
sleep 2
client/test localhost
killall serv
killall master

