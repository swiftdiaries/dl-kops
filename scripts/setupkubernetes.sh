#!/bin/bash
#usage: ./setupkubernetes.sh
echo "######################### INSTALL DOCKER ##########################################"
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update
# Check Docker version
apt-cache policy docker-ce
# Install Docker
sudo apt-get install -y docker-ce=17.03.0~ce-0~ubuntu
# Run docker without sudo
sudo groupadd docker
sudo usermod -aG docker ${USER}
su - ${USER}
echo 'You might need to reboot / relogin to make docker work correctly'

echo "######################### INSTALL GOLANG ##########################################"
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
mkdir -p ~/go
mkdir -p ~/go/src
mkdir -p ~/go/bin
export GOPATH=~/go

echo "######################### KUBERNETES ##########################################"
KUBERNETES_VERSION=1.11.0
KUBERNETES_CNI=0.6.0
sudo bash -c 'apt-get update && apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get -y update'
sudo apt-get install -yf \
  socat \
  ebtables \
  apt-transport-https \
  kubelet=${KUBERNETES_VERSION}-00 \
  kubeadm=${KUBERNETES_VERSION}-00 \
  kubernetes-cni=${KUBERNETES_CNI}-00 \
  kubectl=${KUBERNETES_VERSION}-00

echo "######################### NVIDIA-DOCKER ##########################################"
# If you have nvidia-docker 1.0 installed: we need to remove it and all existing GPU containers
docker volume ls -q -f driver=nvidia-docker | xargs -r -I{} -n1 docker ps -q -a -f volume={} | xargs -r docker rm -f
sudo apt-get purge -y nvidia-docker
# Add the package repositories
curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | \
  sudo apt-key add -
distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | \
  sudo tee /etc/apt/sources.list.d/nvidia-docker.list
sudo apt-get update
# Install nvidia-docker2 and reload the Docker daemon configuration
sudo apt-get install -y nvidia-docker2
sudo pkill -SIGHUP dockerd
# NOTE: Check nvidia-docker runtime in /etc/docker/daemon.json
sudo systemctl daemon-reload
sudo systemctl restart kubelet

echo "######################### Clean-up ##########################################"
sudo rm -rf *.tgz *.deb
