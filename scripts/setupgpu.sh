#!/bin/bash
#usage: ./setupgpu.sh
DRIVER_VERSION=384
echo "################################# Install nvidia driver ######################################"
DRIVER_VERSION=384
sudo apt-get install -y linux-headers-$(uname -r)
sudo apt-get install -y software-properties-common
sudo add-apt-repository -y ppa:graphics-drivers
sudo apt-get update
sudo apt install -y nvidia-$DRIVER_VERSION


echo "################################# Install CUDA ######################################"
wget https://developer.nvidia.com/compute/cuda/9.2/Prod2/local_installers/cuda-repo-ubuntu1604-9-2-local_9.2.148-1_amd64.deb
sudo apt-key add /var/cuda-repo-9-2-local/7fa2af80.pub
sudo apt-get update
sudo apt-get install -y cuda

# Cleanup
sudo rm *.deb
