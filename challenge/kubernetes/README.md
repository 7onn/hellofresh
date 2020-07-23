# kubernetes
this directory contains essential schema and values files to run every service expected by [devbytom-sre-test](https://github.com/hellofreshdevtests/devbytom-sre-test/tree/master)


## TL;DR;
```bash
make add-production
# make rm-production
```

#### namespace
first we need somewhere to install our charts in minikube 
```bash
make
make add-namespace
#make rm-namespace
#"sudo rm -rf /" equivalent commented above
```

#### database
then setup mysql internal service (this is not available from outside the cluster context)
```bash
make add-database
#make rm-database
```

#### server
once database container is healthy, install hellofresh-server chart
```bash
make add-server
# make rm-server
```

#### after install
check the application pods with
```bash
kubectl -n production get po
```

once everything is healthy, you might test the endpoints mentioned at the [_README.md](https://github.com/hellofreshdevtests/devbytom-sre-test/tree/dev/_README.md) on the repository's root directory
```bash
curl -X GET http://minikube:32080/metrics
```
