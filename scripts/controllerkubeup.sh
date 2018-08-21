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
sudo sed -i -e 's/Environment=/Environment="KUBELET_EXTRA_ARGS=--feature-gates=DevicePlugins=true"/' $FILE_NAME
sudo systemctl daemon-reload
sudo systemctl restart kubelet
sudo kubeadm init --apiserver-advertise-address=$ipaddress --pod-network-cidr=192.168.0.0/16
mkdir -p $HOME/.kube
sudo cp -f /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
sudo kubectl apply -f http://docs.projectcalico.org/v2.3/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml
cd
mkdir -p $HOME/.kube
sudo cp -f /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
export KUBECONFIG=$HOME/.kube/config
kubectl apply -f http://docs.projectcalico.org/v2.3/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml
# Uncomment next line for single node cluster
# kubectl taint nodes --all node-role.kubernetes.io/master-