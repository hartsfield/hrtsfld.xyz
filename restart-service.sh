#!/bin/bash
export servicePort=5659
export logFilePath=./logfile.txt
trap -- '' SIGTERM
git pull
go build -o ComfortStarr
pkill -f ComfortStarr
nohup ./ComfortStarr > /dev/null & disown
sleep 2
