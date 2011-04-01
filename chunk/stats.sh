#!/bin/bash
#who, mpstat, free
WHO=$(who | awk '{print $1;}'| uniq | wc -l |awk '{print $1;}')
echo $WHO
FREE=$(free -m | sed -n '2,2p' | awk '{print $4;}')
TOTAL=$(free -m | sed -n '2,2p' | awk '{print $2;}')
echo $FREE
echo $TOTAL
CPU=$(mpstat 1 3 | sed -n '7,7p' | awk '{print $11;}')
echo $CPU
