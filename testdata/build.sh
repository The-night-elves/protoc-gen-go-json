#!/usr/bin/env bash


pluginName="protoc-gen-go-json"
pluginOutName="--go-json_out"


protoc -I proto proto/* --go_out=. --go-vtproto_out=. --plugin=$pluginName=../protoc-gen-go-json $pluginOutName=.


