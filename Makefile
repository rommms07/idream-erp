export ROOTDIR=$(shell pwd)

all:

generate-go-pb: $(wildcard ${ROOTDIR}/core/stubs/*.proto)
	${ROOTDIR}/helpers/generate-protobuf.sh

test:
	go test -v ./...