#!/bin/bash


cd include;
make clean;
make;
cd ../logger;
make clean;
make;
cd ../master;
make clean;
make;
cd ../chunk;
make clean;
make;
cd ../client;
make clean;
make;
