export ROOTDIR=$(shell pwd)

.EXPORT_ALL_VARIABLES:
ifeq (${ENV}, devel)
include .env.development
else
include .env.production
endif

.PHONY: all clean generate-go-pb test api 
all:

clean:
	rm -rf ${ROOTDIR}/core/pb/*.go	

generate-go-pb: $(wildcard ${ROOTDIR}/core/stubs/*.proto)
	${ROOTDIR}/helpers/generate-protobuf.sh

test:
	go clean -v -testcache && SERVER_ADDR=localhost:5000 go test -v ./...

api:
	go run ${ROOTDIR}/cmd/main.go

examples/fb_loginflow:
	${ROOTDIR}/examples/core/auth/fb-loginflow.sh
