package plugin

import (
	"github.com/pangum/qingniao/internal/config"
)

type Qingniao struct {
	Sms    *config.Sms    `json:"sms,omitempty" yaml:"sms" xml:"sms" toml:"sms"`
	Email  *config.Email  `json:"email,omitempty" yaml:"email" xml:"email" toml:"email"`
	Wechat *config.Wechat `json:"wechat,omitempty" yaml:"wechat" xml:"wechat" toml:"wechat"`
}
