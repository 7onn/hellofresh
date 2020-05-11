## HelloFresh Site Reliability Engineer Test

Hello and thank you for your time.

The goal of this test is to assess your coding, testing, automation, and documentation skills. You're given a simple problem so that you can focus on showcasing your techniques.

## Problem definition

We would like you to create a simple HTTP API service using Python or GoLang, that provides observability aspects, collects metrics from operations, and adds logs and tracing for HTTP requests.

Since we love automating things, the service should be automatically deployed to Kubernetes.

_Note: While we love open source here at HelloFresh, please do not create a public repo for your test! This challenge is only shared with people interviewing, and for obvious reasons, we'd like it to remain this way._

## Instructions

1. Clone this repository.
2. Create a new `dev` branch.
3. Solve the task and commit your code. Commit often; we like to see small commits that build-up to the result of your test, instead of one final commit with all the code.
4. Do a pull request from the `dev` branch to the `master` branch. More on that right below.
5. Reply to the thread you are having with our HR department so we can start reviewing your code.

In your pull request, make sure to write about your approach in the description. One or more of our engineers will then perform a code review.
We will ask questions which we expect you to be able to answer. Code review is an important part of our process; this gives you as well as us a better understanding of what working together might be like.

We believe it will take 4 to 8 hours to develop this task; however, feel free to invest as much time as you would like.

## Requirements

### Endpoints

Your application **MUST** conform to the following endpoint structure and return the HTTP status codes appropriate to each operation.

The following endpoints should be implemented:

| Name   | Method      | URL
| ---    | ---         | ---
| List   | `GET`       | `/configs`
| Create | `POST`      | `/configs`
| Get    | `GET`       | `/configs/{name}`
| Update | `PUT/PATCH` | `/configs/{name}`
| Delete | `DELETE`    | `/configs/{name}`

#### List

The list endpoint **MUST** return all configs.

List example:

```sh
curl http://config-service/configs
```

Response example:

```json
[
  {
    "name": "datacenter-1",
    "metadata": {
      "monitoring": {
        "enabled": "true"
      },
      "limits": {
        "cpu": {
          "enabled": "false",
          "value": "300m"
        }
      }
    }
  },
  {
    "name": "datacenter-2",
    "metadata": {
      "monitoring": {
        "enabled": "true"
      },
      "limits": {
        "cpu": {
          "enabled": "true",
          "value": "250m"
        }
      }
    }
  },
]
```


Get example:

```sh
curl http://config-service/configs/burger-nutrition
```

Response example:

```json
[
  {
    "name": "burger-nutrition",
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
        "nuts": "false",
        "seafood": "false",
        "eggs": "true"
      }
    }
  }
]
```

#### Schema

- **Config**
  - Name (string).
  - Metadata (nested key:value pairs where both key and value are strings of arbitrary length).


### Configuration

- Your application **MUST** serve the API on the port defined by the environment variable `SERVE_PORT`.
- The application **MUST** fail if the environment variable is not defined.

### Deployment

- The application **MUST** be deployable on a Kubernetes cluster. Please provide manifest files and a script that deploys the application on a Minikube cluster.
- The application **MUST** be accessible from outside the Minikube cluster.

## Rules
### Instrumentation

- Your application **MUST** expose metrics in Prometheus format.
- Your application **MUST** print structured logs in JSON format to stdout.
- Your application **MUST** publish distributed tracing data in Zipkin or Jaeger format.
- You **SHOULD** write testable code and demonstrate unit testing it.
- You **SHOULD** document your code and scripts.
- You **MAY** use any testing, mocking libraries provided that you state the reasoning and they are simple to install and run.
