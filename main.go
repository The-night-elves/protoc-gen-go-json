package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"protoc-gen-go-json/json"
)

func main() {
	var cfg json.Config

	flagSet := flag.FlagSet{
		Usage: func() {
			fmt.Fprintf(os.Stderr, "Usage of  protoc-gen-go-json:\n")
			fmt.Fprintf(os.Stderr, "\t"+cfg.Usage()+"\n")
			fmt.Fprintf(os.Stderr, "Flags:\n")
			flag.PrintDefaults()
		},
	}

	flagSet.Var(&cfg, "config", "set config args")
	protogen.Options{ParamFunc: flagSet.Set}.Run(func(plugin *protogen.Plugin) error {
		cfg = cfg.SetDefaults()
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return json.Generate(plugin, &cfg)
	})
}
