#!/bin/bash -

if [ ! $ROOTDIR ]; then
  echo "\$ROOTDIR environment variable is not set. Must be executed with `Makefile`"
  exit;
fi

rm -rf $ROOTDIR/core/pb/*

for i in $(ls $ROOTDIR/core/stubs/*.proto); do
  protofile=$(basename -s .proto $i)
  schemapath=${protofile}_schema

  # create a directory if it does not exists.
  [ ! -d ${ROOTDIR}/core/pb/$schemapath ] && mkdir ${ROOTDIR}/core/pb/$schemapath

  protoc -I=${ROOTDIR}/core/stubs --go_opt=paths=source_relative --go_out=${ROOTDIR}/core/pb/$schemapath ${ROOTDIR}/core/stubs/$protofile.proto
done

