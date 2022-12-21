#!/bin/bash -

protoc -I=${ROOTDIR}/core/stubs --go_opt=paths=source_relative --go_out=${ROOTDIR}/core/pb ${ROOTDIR}/core/stubs/*.proto