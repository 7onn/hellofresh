default: help

help:
	@echo 
	@echo "available commands"
	@echo "  - minikube-start"
	@echo "  - minikube-stop"
	@echo "  - minikube-addons-enabled"
	@echo "  - minikube-addons-disabled"
	@echo "  - artifacts"
	@echo

minikube-start:
	@minikube start --vm=true
	@kubectl config use-context minikube

minikube-stop:
	@minikube stop

PHONY: minikube-addons-enabled
minikube-addons-enabled: minikube-addons-disabled
	@minikube addons enable ingress
	@minikube addons enable helm-tiller
	@helm init --wait --upgrade --service-account tiller --override spec.selector.matchLabels.'name'='tiller',spec.selector.matchLabels.'app'='helm' --output yaml | sed 's@apiVersion: extensions/v1beta1@apiVersion: apps/v1@' | kubectl apply -f -

minikube-addons-disabled:
	@minikube addons disable ingress
	@minikube addons disable helm-tiller

artifacts:
	@docker-compose build --no-cache
