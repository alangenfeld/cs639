#!/bin/bash
master/master&
sleep 5
chunk/chunk localhost&
sleep 5
client/client c;
client/client r;
