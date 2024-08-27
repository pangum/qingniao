package config

import (
	"github.com/goexl/config"
)

type Sms struct {
	// nolint:lll
	Chuangcache config.Slice[Chuangcache] `json:"chuangcache,omitempty" yaml:"chuangcache" xml:"chuangcache,innerxml" toml:"chuangcache"`
}
