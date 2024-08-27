package config

type Chuangcache struct {
	LabelBase `yaml:"inline"`

	Ak string `json:"ak,omitempty" yaml:"ak" xml:"ak" toml:"ak" validate:"required"`
	Sk string `json:"sk,omitempty" yaml:"sk" xml:"sk" toml:"sk" validate:"required"`
}
