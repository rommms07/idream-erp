export ROOTDIR=$(shell pwd)

.EXPORT_ALL_VARIABLES:
ifeq (${ENV}, devel)
include .env.development
else
include .env.production
endif

.PHONY: all clean generate-go-pb test 
all:

clean:
	echo ${FB_CLIENT_ID}
	rm -rf ${ROOTDIR}/core/pb/*.go	

generate-go-pb: $(wildcard ${ROOTDIR}/core/stubs/*.proto)
	${ROOTDIR}/helpers/generate-protobuf.sh

test:
	go test -v ./...

examples/fb_loginflow:
	${ROOTDIR}/examples/core/auth/fb-loginflow.sh
