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

gen_go_pb: $(wildcard ${ROOTDIR}/core/stubs/*.proto)
	${ROOTDIR}/helpers/generate-protobuf.sh

db/create:
	${ROOTDIR}/helpers/db.sh create

test:
	go clean -v -testcache && SERVER_ADDR=localhost:5000 go test -v ./...

examples/gorm_mysql_connect:
	go run ${ROOTDIR}/examples/core/source/mysql/connect.go

examples/fb_api:
	go run ${ROOTDIR}/examples/core/auth/main.go

examples/fb_loginflow:
	${ROOTDIR}/examples/core/auth/fb-loginflow.sh

build/cmd:
	go run cmd/main.go 1 2 3 4 5
