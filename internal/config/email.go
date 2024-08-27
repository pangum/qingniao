package config

import (
	"github.com/goexl/config"
)

type Email struct {
	Smtp config.Slice[Smtp] `json:"smtp,omitempty" yaml:"smtp" xml:"smtp,innerxml" toml:"smtp"`
}
