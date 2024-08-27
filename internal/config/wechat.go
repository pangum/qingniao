package config

import (
	"github.com/goexl/config"
)

type Wechat struct {
	// nolint:lll
	Serverchan config.Slice[Serverchan] `json:"serverchan,omitempty" yaml:"serverchan" xml:"serverchan,innerxml" toml:"serverchan"`
}
