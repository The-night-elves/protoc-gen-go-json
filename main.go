package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"protoc-gen-go-json/json"
)

func main() {
	var cfg json.Config
	cfg = cfg.SetDefaults()
	cfg.Debug = true

	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return json.Generate(plugin, &cfg)
	})
}
