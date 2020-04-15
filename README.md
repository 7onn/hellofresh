# WORK IN PROGRESS #

## HelloFresh Site Reliability Engineer Test

Hello and thanks for taking the time to try this out.

The goal of this test is to assess (to some degree) your coding, testing, automation, and documentation skills. You're given a simple problem, so you can focus on showcasing your techniques.

## Problem definition
The test aims to create a simple HTTP service that provides observability aspects, (e.g) you can store and returns random values, collect metrics from operations, add logs and tracing for requests.  

Since we love automating things; the service should be automatically deployed to Kubernetes.

_Note: While we love open source here at HelloFresh, please do not create a public repo with your test in! This challenge is only shared with people interviewing, and for obvious reasons, we'd like it to remain this way._

## Instructions

1. Clone this repository.
2. Create a new `dev` branch.
3. Solve the task and commit your code. Commit often, we like to see small commits that build-up to the result of your test, instead of one final commit with all the code.
4. Do a pull request from the `dev` branch to the `master` branch. More on that right below.
5. Reply to the thread you are having with our HR department so we can start reviewing your code.

In your pull request, make sure to write about your approach in the description. One or more of our engineers will then perform a code review.
We will ask questions which we expect you to be able to answer. Code review is an important part of our process;
this gives you as well as us a better understanding of what working together might be like.

We believe it will take 4 to 8 hours to develop this task, however, feel free to invest as much time as you want.

### Endpoints

Your application **MUST** conforms to the following endpoint structure and returns the HTTP status codes appropriate to each operation.

Following are the endpoints that should be implemented:

| Name   | Method      | URL
| ---    | ---         | ---
| List   | `GET`       | `/configs`
| Create | `POST`      | `/configs`
| Get    | `GET`       | `/configs/{name}`
| Collect| `GET`       | `/metrics`
| Update | `PUT/PATCH` | `/configs/{name}`
| Delete | `DELETE`    | `/configs/{name}`


### Configuration

Your application **MUST** serve the API on the port defined by the environment variable `SERVE_PORT`.
The application **MUST** fail if the environment variable is not defined.

### Deployment

The application **MUST** be deployable on a Kubernetes cluster. Please provide manifest files and a script that deploys the application on a Minikube cluster.
The application **MUST** be accessible from outside the Minikube cluster.


### Metrics expectation

We love see as much as indicators you can expose from the application, but please don't forget those metrics as we love see below
- Total incoming requests which are counted under 200. Similarly for 400 - 404 and 500
- Uptime of service which is rotated after every deployment time

### Logging expectation

- Ideally the HTTP status codes can be filtered from HTTP access logs
- For those manipulation requests like PUT/POST please give us a proper reason whether they are processed or not
- Would be perfect if each requests can has REQUEST ID in log-format and searchable

We currently dealing with JSON Log format everyday but feel free to choose the one your most familiar

### Tracing expectation

## Rules

-  Applicants for SRE roles are requested to use either Python or GoLang.
- You **SHOULD** write testable code and demonstrate unit testing it.
- You can use any testing, mocking libraries provided that you state the reasoning and it's simple to install and run.
- You **SHOULD** document your code and scripts.
- YOU **MUST** have metrics exposed to Prometheus. (Please don't use MUST here. I think we can ask them feel free to use any indicator format)
- YOU **MUST** instrument your code, publishing distributed tracing to Jaeger. (Similarly with Prometheus expose, and not MUST)
