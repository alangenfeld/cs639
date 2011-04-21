#!/bin/bash


ssh mumble-40 "killall cproxy"
ssh mumble-39 "killall cproxy"
ssh mumble-38 "killall cproxy"
ssh mumble-37 "killall cproxy"
ssh mumble-01 "killall proxy"
