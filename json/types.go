package json

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
