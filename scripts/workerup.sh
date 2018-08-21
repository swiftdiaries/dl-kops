#!/bin/bash
#Usage: ./slavejoin.sh [join-token] [masterIPaddress] [discovery certs]
#Assumes Kubernetes is already-installed using setupkubernetes.sh
if [ -z "$1" ]
then
	echo "Token needed to join worker node to master"
else
	token="$1"
fi
if [ -z "$2" ]
then
	echo "Master IP address needed to join worker node to master"
else
	masterIPaddress="$2"
fi
if [ -z "$3" ]
then
	echo "Discovery certs needed to join worker node to master"
else
	certs="$3"
fi
echo $1 $2 $3
sudo kubeadm reset
sudo systemctl enable docker
sudo systemctl start docker
sudo systemctl enable kubelet
sudo systemctl start kubelet
for file in /etc/systemd/system/kubelet.service.d/*-kubeadm.conf
do
    echo "Found ${file}"
    FILE_NAME=$file
done
echo "Chosen ${FILE_NAME} as kubeadm.conf"
sudo sed -i -e 's/Environment="KUBELET_EXTRA_ARGS=--feature-gates=DevicePlugins=true" /g' $FILE_NAME
sudo systemctl daemon-reload
sudo systemctl restart kubelet
sudo kubeadm join --token $token $masterIPaddress:6443 --discovery-token-ca-cert-hash $certs

# support for NodeAffinity
#NVIDIA_GPU_NAME=$(nvidia-smi --query-gpu=gpu_name --format=csv,noheader --id=0)
#source /etc/default/kubelet
#KUBELET_OPTS="$KUBELET_OPTS --node-labels='alpha.kubernetes.io/nvidia-gpu-name=$NVIDIA_GPU_NAME'"
#echo "KUBELET_OPTS=$KUBELET_OPTS"
#echo "KUBELET_OPTS=$KUBELET_OPTS" > /etc/default/kubelet
#KUBELET_OPTS=--node-labels='alpha.kubernetes.io/nvidia-gpu-name=Tesla M40'
#sudo systemctl daemon-reload
#sudo systemctl restart kubelet
