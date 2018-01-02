#!/bin/bash
# usage:
# ./kubernetes_cluster.sh [username] [masterIP] [key file location]
# this script has to be executed on the master node.
# copy the pem/key file to the master node before running this script.
# also manually add hostname to /etc/hosts
#for server in $serverList; do
#		$SSH_CMD $username@$server "sudo sed -i -e 's/127.0.0.1 localhost/127.0.0.1 localhost \n127.0.0.1 $HOSTNAME/g' /etc/hosts" &
#done	
#wait

master="SETCloud";
masterIP="129.110.96.19"
keyfile="~/cloud.key"
#slaves="p100";
#slavesIP="129.114.108.146";
#servers="$master $slaves";
#serversIP="$masterIP $slavesIP";
: '
if [ -z "$1" ]
then
	username="cc"
else
	username="$1"
fi
if [ -z "$3" ]
then
	keyfile=~/cloud.key
else
	keyfile="$3"
fi
if [ -z "$2" ]
then
	masterIP="129.110.96.19"
else
	masterIP="$2"
fi'

SSH_CMD="ssh -i $keyfile"

# setup kubernetes
$SSH_CMD $master@$masterIP 'bash -s' < ./setupkubernetes.sh 
#for server in $slavesIP; do
#		$SSH_CMD $username@$server 'bash -s' < ./setupkubernetes.sh &
#done	
#wait

# configure kubernetes master
#$SSH_CMD $username@$master 'bash -s' < ./masterkubeup.sh $masterIP
./masterkubeup.sh $masterIP
#echo "Enter Token :"
#read token
# configure kubernetes slave
#for server in $slavesIP; do
#		$SSH_CMD $username@$server 'bash -s' < ./slavejoin.sh $token $masterIP
#done
