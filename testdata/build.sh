#!/usr/bin/env bash


pluginName="protoc-gen-go-json"
pluginOutName="--go-json_out"
pluginConfigName="--go-json_opt"


protoc -I proto proto/* --go_out=. \
 --plugin=$pluginName=../protoc-gen-go-json $pluginOutName=. \
$pluginConfigName=config=FileNameSuffix=.json.go,config=EncodeMethodName=MarshalJSON


