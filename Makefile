# Makefile

.PHONY: mk_tools clean

mk_tools:
	if [ ! -d "./tools" ]; then \
	    mkdir ./tools; \
		echo "Directory 'tools' created."; \
	else \
		echo "Directory 'tools' already exists."; \
	fi


export GOPATH=${PWD}/tools
export PATH=$(shell printenv PATH):${PWD}/tools/bin
install_protoc: mk_tools
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	cd ./tools && \
	npm init -y && \
	npm i --save-dev @grpc/grpc-js

gen: install_protoc
	protoc -I ${PWD} items.proto --go_out=./backend/domain --go-grpc_out=./backend/domain
	./tools/node_modules/@grpc/proto-loader/build/bin/proto-loader-gen-types.js ./items.proto -O ./frontend/src --grpcLib=@grpc/grpc-js
	cp ./items.proto ./frontend/src