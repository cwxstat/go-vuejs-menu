.PHONY: default help

default: help
help: ## display make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(word 1, $(MAKEFILE_LIST)) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m make %-20s -> %s\n\033[0m", $$1, $$2}'

.PHONY: dist
dist: ## create dist folder under Go
	@bash -c "cd vue-project;npm install; npm run build"
	@bash -c "rm -rf ./Go/dist"
	@bash -c "mv ./vue-project/dist ./Go/dist"


.PHONY: dev
dev: ## npm run dev
	@bash -c "cd vue-project;npm install; npm run dev"

.PHONY: run
run: dist ## dist and run Go
	@bash -c "cd Go; go run server.go"


.PHONY: docker
docker: ## todo ... build docker image
	@bash -c "echo 'not implemented yet'"

