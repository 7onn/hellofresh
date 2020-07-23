default: help

help:
	@echo 
	@echo "available commands"
	@echo "  - add-namespace"
	@echo "  - rm-namespace"
	@echo "  - add-database"
	@echo "  - rm-database"
	@echo "  - add-server"
	@echo "  - rm-server"
	@echo
	@echo "  - add-production (i.e: install production namespace and charts)"
	@echo "  - rm-production (i.e: remove production namespace and it's content)"
	@echo
	@echo
	@echo "note: you also might use add scripts to update charts"
	@echo

add-namespace:
	@kubectl create ns production

rm-namespace:
	@kubectl delete ns production

add-database:
	@helm upgrade -i hellofresh-database-production \
		--namespace production \
		--values ./values/hellofresh-database/production.yaml \
		./charts/hellofresh-database

rm-database:
	@helm del --purge hellofresh-database-production

add-server:
	@sleep 10
	@helm upgrade -i hellofresh-server-production \
		--namespace production \
		--values ./values/hellofresh-server/production.yaml \
		./charts/hellofresh-server

rm-server:
	@helm del --purge hellofresh-server-production

.PHONY: add-production
add-production: add-namespace add-database add-server
	@sleep 10
	@kubectl -n production get pods

.PHONY: rm-production
rm-production: rm-server rm-namespace
