SWAGGER_CLI_VER=v0.27.0
SWAGGER_CLI=docker run --rm -it --user $$(id -u):$$(id -g) -e GOPATH=$$HOME/go:/go -v $$HOME:$$HOME -w $$(pwd) quay.io/goswagger/swagger:$(SWAGGER_CLI_VER)

.PHONY: default
default: help

.PHONY: help
help:
	@echo Usage: make <target...>

.PHONY: generate
generate: generate__aspida generate__go_swagger

.PHONY: build
build: build__web-ui

.PHONY: generate__aspida
generate__aspida: swagger/swagger.yml
	rm -rf ./projects/web-ui/src/api/
	cd ./projects/web-ui/src && node ~/ghq/github.com/aspida/openapi2aspida/bin/index.js -i ../../../swagger/swagger.yml -c ../aspida.config.cjsa

.PHONY: generate__go_swagger
generate__go_swagger: swagger/swagger.yml
	rm -rf ./projects/cli/gen/
	mkdir -p ./projects/cli/gen/
	$(SWAGGER_CLI) generate server --exclude-main -t ./projects/cli/gen -f ./swagger/swagger.yml

.PHONY: build__web-ui
build__web-ui:
	rm -rf ./projects/web-ui/dist
	pnpm --filter=./projects/web-ui build
	rm -rf ./projects/cli/web-ui-dist
	cp -r ./projects/web-ui/dist ./projects/cli/web-ui-dist
