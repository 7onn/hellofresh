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

#### jaeger
It was quite troublesome to implement it locally on a dual core with 8Gb of RAM =(
then I gave up this step; holefully the whole documentation and scripts will be enough for this challenge assessment


#### after install
check the application pods with
```bash
kubectl -n production get po
```

once everything is healthy, you might test the endpoints mentioned at the [_README.md](https://github.com/hellofreshdevtests/devbytom-sre-test/tree/dev/_README.md) on the repository's root directory
```bash
# to scrape prometheus metrics
curl -X GET http://minikube:32080/metrics

# to add a data center
curl -X POST http://minikube:32080/configs \
    -H "Content-Type: application/json" \
    -d \
    '{ 
        "name": "datacenter_devbytom",
        "metadata": {
            "monitoring": {
                "enabled": true
            },
            "limits": {
                "cpu": {
                "enabled": false,
                "value": "300m"
                }
            }
        }
    }'

# to list data centers
curl -X GET http://minikube:32080/configs

# to add a meal
curl -X PUT http://minikube:32080/configs/meal_devbytom \
    -H "Content-Type: application/json" \
    -d \
    '{
        "name": "meal_devbytom",
        "metadata": {
            "calories": 230,
                "fats": {
                "saturated-fat": "0g",
                "trans-fat": "1g"
            },
            "carbohydrates": {
                "dietary-fiber": "4g",
                "sugars": "1g"
            },
            "allergens": {
                "nuts": false,
                "seafood": false,
                "eggs": true
            }
        }
    }'

# to get an existing meal
curl -X GET http://minikube:32080/configs/meal_devbytom

# to update an existing meal
curl -X PATCH http://minikube:32080/configs/meal_devbytom \
    -H "Content-Type: application/json" \
    -d \
    '{
        "name": "updt_meal_devbytom",
        "metadata": {
            "calories": 231,
                "fats": {
                "saturated-fat": "0g",
                "trans-fat": "1g"
            },
            "carbohydrates": {
                "dietary-fiber": "4g",
                "sugars": "2g"
            },
            "allergens": {
                "nuts": false,
                "seafood": false,
                "eggs": true
            }
        }
    }'

# to delete an existing meal
curl -X DELETE http://minikube:32080/configs/updt_meal_devbytom
```
