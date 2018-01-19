#!/bin/bash
# usage:
# ./kubernetes_cluster.sh [username] [key file location] [masterIP]
# this script has to be executed on the master node.
# copy the pem/key file to the master node before running this script.
# also manually add hostname to /etc/hosts
#for server in $serverList; do
#		$SSH_CMD $username@$server "sudo sed -i -e 's/127.0.0.1 localhost/127.0.0.1 localhost \n127.0.0.1 $HOSTNAME/g' /etc/hosts" &
#done	
#wait
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
PROJ_DIR=$GOPATH"/src/github.com/swiftdiaries/dl-kops/"
SSH_CMD="ssh -i $keyfile"
SCP_CMD="scp"

echo $1, $2, $3, $PROJ_DIR, $SSH_CMD, $hostname@$hostip
# setup kubernetes on master
$SSH_CMD $hostname@$hostip 'bash -s' < $PROJ_DIR/scripts/setupkubernetes.sh 
$SCP_CMD ./scripts/controllerkubeup.sh $hostname@$hostip:~/ 
$SSH_CMD $hostname@$hostip chmod +x controllerkubeup.sh 
$SCP_CMD $hostname@$hostip:~/config/admin.conf ~/Desktop/admin.conf
export KUBECONFIG=~/Desktop/admin.conf