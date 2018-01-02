#!/bin/bash
#usage: ./setupkubernetes.sh
echo "######################### DOCKER ##########################################"
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
wget https://apt.dockerproject.org/repo/pool/main/d/docker-engine/docker-engine_1.12.6-0~ubuntu-xenial_amd64.deb
sudo dpkg -i docker-engine_1.12.6-0~ubuntu-xenial_amd64.deb
sudo apt install -y ~/docker-engine_1.12.6-0~ubuntu-xenial_amd64.deb
sudo groupadd docker
sudo usermod -aG docker $USER
echo 'You might need to reboot / relogin to make docker work correctly'
echo "######################### KUBERNETES ##########################################"
sudo bash -c 'apt-get update && apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update'
sudo apt-get install -y kubelet kubeadm kubectl kubernetes-cni
echo "######################### NVIDIA-DOCKER ##########################################"
wget https://github.com/NVIDIA/nvidia-docker/releases/download/v1.0.1/nvidia-docker_1.0.1-1_amd64.deb
sudo apt-get install -f
echo "######################### PATH ##########################################"
export PATH=/usr/local/cuda-8.0/bin${PATH:+:${PATH}}
export LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/local/cuda/lib64:/usr/local/cuda/extras/CUPTI/lib64"
echo "######################### KUBEADM RESET ##########################################"
sudo kubeadm reset
echo "######################### Clean-up ##########################################"
sudo rm -rf *.tgz *.deb
echo "######################### DOCKER-PULL ##########################################"
sudo docker pull swiftdiaries/bench