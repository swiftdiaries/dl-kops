#!/bin/bash
# usage:
# ./trial.sh [username] [masterIP] [key file location]

if [ -z "$1" ]
then
	hostip="0.0.0.1"
else
	hostip="$3"
fi

echo "
hostip:$1
"