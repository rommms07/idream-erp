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

access_token_demo:
	curl -i -X GET \
		"https://www.facebook.com/v15.0/dialog/oauth?client_id=${FB_CLIENT_ID}&redirect_uri=\{"https://www.google.com"\}"