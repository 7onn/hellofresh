# Description
this project is meant to accomplish the HelloFresh's SRE challenge; failed btw hahahaha but the code is there

#### my setup
- Darwin OS 10.15.5
- [docker](https://www.docker.com/get-started) 19.03.8 (afacb8b)
- [docker-compose](https://docs.docker.com/compose/install/) 1.25.5 (8a1c60f6)
- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) v1.12.0 
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) (server) v1.18.3
- [helm](https://github.com/helm/helm/releases/tag/v2.14.2) v2.14.2

#### get started
firstly, check the helper:
```bash
make
```

#### minikube
```bash
make minikube-start
make minikube-addons-enabled
eval $(minikube docker-env)
# eval $(minikube docker-env -u) if you want to switch back from minikube docker context
make artifacts
```

append the following command's output in your `/etc/hosts`
```bash
echo $(minikube ip) minikube
```
e.g:
```
192.168.64.2 minikube
```

once every artifact is successfully built and the minikube is identified as a HOST in your computer, you might follow [this](https://github.com/hellofreshdevtests/devbytom-sre-test/tree/dev/challenge/kubernetes) in order to install the application charts
```bash
cd ./challenge/kubernetes
```
