package g

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
}

var RootPath string

var RBAC_CONFIG = Local{
	Path: "uploads/file",
}
