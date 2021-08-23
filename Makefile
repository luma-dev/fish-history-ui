GO_SWAGGER_VER=v0.27.0
GO_SWAGGER=docker run --rm --user $$(id -u):$$(id -g) -e GOPATH=$$HOME/go:/go -v $$HOME:$$HOME -w $$(pwd) quay.io/goswagger/swagger:$(GO_SWAGGER_VER)

.PHONY: default
default: help

.PHONY: help
help:
	@echo Usage: make <target...>

.PHONY: generate
generate: generate/aspida generate/go_swagger

.PHONY: build
build: generate build/web-ui build/cli

.PHONY: generate/aspida
generate/aspida: swagger/swagger.yml
	rm -rf ./projects/web-ui/src/api/
	cd ./projects/web-ui/src && ../../../node_modules/.bin/openapi2aspida -i ../../../swagger/swagger.yml

.PHONY: generate/go_swagger
generate/go_swagger: swagger/swagger.yml
	rm -rf ./projects/cli/gen/
	mkdir -p ./projects/cli/gen/
	$(GO_SWAGGER) generate server --exclude-main -t ./projects/cli/gen -f ./swagger/swagger.yml

.PHONY: build/web-ui
build/web-ui:
	rm -rf ./projects/web-ui/dist
	pnpm --filter=./projects/web-ui build
	rm -rf ./projects/cli/web-ui-dist
	cp -r ./projects/web-ui/dist ./projects/cli/web-ui-dist

.PHONY: build/cli
build/cli:
	cd ./projects/cli && go build
