#!/bin/bash

# TODO replace this script to Bazel.

set -o nounset

WORKDIR="${WORKDIR:=./}"

TMP_DIR=$(mktemp -d)

rm -rf ./pkg/protos/*
rm -rf ./gateway/*
rm -rf ./api/*

mkdir -p ./{gateway,api}

function gen_server_protos {
  protoc \
    -I=./protos \
    -I=./protos/vendor/bufbuild/protovalidate/proto/protovalidate \
    -I=./protos/vendor/grpc-ecosystem/grpc-gateway \
    -I=./protos/vendor/protocolbuffers/protobuf/src \
    -I=./protos/vendor/googleapis/googleapis \
    --go_out=$TMP_DIR \
    --go-grpc_out=$TMP_DIR \
    --validate_out="lang=go:$TMP_DIR" \
    --openapiv2_out ./api/ \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt disable_default_errors=true \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt standalone=true \
    --grpc-gateway_out ./gateway/ \
    $(find protos/todo-list -type f -name "*.proto")
}

function gen_pkg_protos {
  protoc \
    -I=./pkg \
    -I=./protos/vendor/bufbuild/protovalidate/proto/protovalidate \
    -I=./protos/vendor/protocolbuffers/protobuf/src \
    -I=./protos/vendor/googleapis/googleapis \
    --go_out=$TMP_DIR \
    --go-grpc_out=$TMP_DIR \
    $(find pkg -type f -name "*.proto")
}

gen_server_protos
gen_pkg_protos

cp -R $TMP_DIR/github.com/kyfk/golang-todo-list-sample/* $WORKDIR
