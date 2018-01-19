#!/bin/bash
# usage:
# ./trial.sh [username] [masterIP] [key file location]

if [ -z "$1" ]
then
	hostname="cc"
else
	hostname="$1"
fi
if [ -z "$2" ]
then
	keyfile=~/keyfile.pem
else
	keyfile="$2"
fi
if [ -z "$3" ]
then
	hostip="0.0.0.1"
else
	hostip="$3"
fi

echo "
hostname:$1
hostip:$3
keyfile:$2
"