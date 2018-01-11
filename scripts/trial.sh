#!/bin/bash
if [ -z "$1" ]
then
	hostip="128.110.96.19"
else
	hostip="$1"
fi
echo "hostip: "$hostip