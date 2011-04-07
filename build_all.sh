#!/bin/bash

CLEAN=""

if [ "$1" == "clean" ]; then
	CLEAN="clean-all"
fi

(cd include && make $CLEAN)
(cd master && make $CLEAN)&
(cd client && make $CLEAN)&
(cd logger && make $CLEAN && cd ../chunk && make $CLEAN)&

wait
