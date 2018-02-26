#!/bin/bash
#usage: ./masterkubeup.sh [masterIPaddress]
if [ -z "$1" ]
then
	ipaddress="0.0.0.1"
else
	ipaddress="$1"
fi
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
sudo swapoff -a
sudo sed -i -e 's/ExecStart=\/usr\/bin\/kubelet /ExecStart=\/usr\/bin\/kubelet --feature-gates="Accelerators=true" /g' $FILE_NAME
sudo systemctl daemon-reload
sudo systemctl restart kubelet
sudo kubeadm reset
sudo kubeadm init --apiserver-advertise-address=$ipaddress --pod-network-cidr=192.168.0.0/16
mkdir -p $HOME/.kube
sudo cp -f /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
export KUBECONFIG=~/.kube/config
sudo kubectl apply -f http://docs.projectcalico.org/v2.3/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml
cd
mkdir -p config
sudo cp -f /etc/kubernetes/admin.conf config/admin.conf
sudo chmod 777 config/admin.conf
