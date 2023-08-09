package json

import (
	"errors"
	"fmt"
	"strings"
)

// Config is the configuration for the json plugin.
type Config struct {
	// 输出的文件名后缀, default: .json.go
	FileNameSuffix string
	// encode json method name
	EncodeMethodName string

	// import writer
	ImportWriter string
	// new writer
	NewWriter string
	// write bytes
	WriteBytes string

	// debug logging
	Debug bool
}

func (c *Config) String() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf(
		"FileNameSuffix=%s,EncodeMethodName=%s,ImportWriter=%s,NewWriter=%s, WriteBytes=%s, Debug=%t",
		c.FileNameSuffix, c.EncodeMethodName, c.ImportWriter, c.NewWriter, c.WriteBytes, c.Debug)
}

func (c *Config) Usage() string {
	return "config args, format: key=val, " +
		"support keys: [FileNameSuffix,EncodeMethodName,ImportWriter,NewWriter,WriteBytes,Debug]" +
		"example: FileNameSuffix=.json.go,EncodeMethodName=MarshalJSON,ImportWriter=bytes,NewWriter=Buffer,WriteBytes=.Bytes(),Debug=true"
}

func (c *Config) Set(s string) error {
	if c == nil {
		return fmt.Errorf("config is nil")
	}
	if len(s) == 0 {
		return nil
	}

	return c.parseStr(s)
}

// SetDefaults 参数配置设置默认值
func (c *Config) SetDefaults() Config {
	cfg := Config{}
	if c != nil {
		cfg = *c
	}
	if len(cfg.FileNameSuffix) == 0 {
		cfg.FileNameSuffix = ".json.go"
	}
	if len(cfg.EncodeMethodName) == 0 {
		cfg.EncodeMethodName = "MarshalJSON"
	}
	if len(cfg.ImportWriter) == 0 {
		cfg.ImportWriter = "bytes"
		cfg.NewWriter = "Buffer"
		cfg.WriteBytes = ".Bytes()"
	}

	return cfg
}

func (c *Config) parseStr(s string) error {
	split := strings.Split(s, ",")
	if len(split) == 0 {
		return nil
	}
	for _, val := range split {
		list := strings.Split(val, "=")
		if len(list) != 2 {
			return errors.New("expect format key=val, actual format %s" + val)
		}
		switch list[0] {
		case "FileNameSuffix":
			c.FileNameSuffix = list[1]
		case "EncodeMethodName":
			c.EncodeMethodName = list[1]
		case "ImportWriter":
			c.ImportWriter = list[1]
		case "NewWriter":
			c.NewWriter = list[1]
		case "WriteBytes":
			c.WriteBytes = list[1]
		case "Debug":
			c.Debug = list[1] == "true" || list[1] == "True"
		default:
			return errors.New("not support config key " + list[0])
		}
	}

	return nil
}

const (
	// Instance 实例缩写
	Instance = "x"

	Buf = "buf"

	WriteString = ".WriteString"
	WriteBytes  = ".Write"
	WriteByte   = ".WriteByte"
	// CommaVarName 逗号变量名
	CommaVarName = "writeComma"
	CommaValue   = "(',')"
)
