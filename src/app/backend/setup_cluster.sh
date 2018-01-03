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
masterIP="128.110.96.19"
keyfile="~/cloud.key"
SSH_CMD="ssh -i $keyfile"
SCP_CMD="scp"

# setup kubernetes on master
$SSH_CMD $master@$masterIP 'bash -s' < ./src/app/backend/setupkubernetes.sh 
$SCP_CMD ./src/app/backend/masterkubeup.sh $master@$masterIP:~/ 
$SSH_CMD $master@$masterIP chmod +x masterkubeup.sh 
SSH_CMD $master@$masterIP ./masterkubeup.sh