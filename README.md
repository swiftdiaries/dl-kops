dl-kops
=======

<p align="left">
<img src="https://goreportcard.com/badge/github.com/swiftdiaries/dl-kops" alt="Go Report">
</p>

About
-----------
This app is to simplify machine learning on Kubernetes for beginners. **NOTE:** This is a work in progress. 

* MVP of an end-to-end deployment app for machine learning jobs on Kubernetes clusters.
* Enables single click installation of CUDA, CuDNN, Kubernetes components.
* Single click bootstrapping of clusters uses kubeadm.
* Single click activation of alpha features to enable discovery of GPUs.
* Runs on Ubuntu machines.
* No guarantee that everything works, use at your own risk.
* Written in Go.


Future Work
-----------
* Prometheus, Grafana dashboard for monitoring.
* Tensorboard for visualizing training.
* Support for custom schedulers.


Setup
-----
* Requires kubectl on the local machine

```bash
    $ brew install kubectl
```

```bash
    $ cd $GOPATH/src
    $ mkdir -p github.com/swiftdiaries
    $ cd github.com/swiftdiaries/
    $ git clone github.com/swiftdiaries/dl-kops
    $ cd dl-kops/src/app/frontend
    $ export PORT_1=8081
    $ go run dashboard.go
``` 

Go to http://localhost:8081/ on your browser.

Get Involved
------------

### Contributing
* If you have a cool new feature in mind, please consider creating an Issue or even better, a Pull Request.
* If you are annoyed by the loading GIF, please do consider contributing.
* Please feel free take a look at the Issues and contribute to any of them. 

I would love feedback and comments,
* [Twitter](https://twitter.com/adhitadselvaraj)
* Kubernetes and Golang slack (adhita). 

Landing Page
------------
![Landing Page](landingpage.png) <br>     
