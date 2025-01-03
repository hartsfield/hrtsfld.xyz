#!/bin/bash
pkill -9 $1
go build -o $1
servicePort=5659 logFilePath=./logfile.txt ./$1 & 
