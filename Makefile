SHELL := /bin/bash

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate-mocks: gen-mocks-docs ## generates mock files for services
	@echo "Generating go server mocks"
	@go generate -v ./...

.SILENT: gen-mocks-docs
gen-mocks-docs: ## generates doc.go file used for generating server mocks
	for service in `ls ./services`;											\
	do																		\
	  	rm -rf ./services/"$$service"/mocks;								\
	  	if  [[ -f ./services/"$$service"/"$$service"_grpc.pb.go ]]; then	\
		  	mkdir ./services/"$$service"/mocks;								\
	  		echo "// Package mocks ..." >> ./services/"$$service"/mocks/doc.go;	\
	  		echo "package mocks" >> ./services/"$$service"/mocks/doc.go;	\
	  		echo "" >> ./services/"$$service"/mocks/doc.go;	\
	  		echo "//go:generate mockgen -source ../"$$service"_grpc.pb.go -destination=./"$$service"_grpc_mock.go -package mocks" >> ./services/"$$service"/mocks/doc.go;	\
	  	fi;		\
	done;		\


test:
	@go test -v ./...

# All targets are phony
.PHONY: $(MAKECMDGOALS)
