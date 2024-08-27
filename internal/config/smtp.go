package config

type Smtp struct {
	LabelBase `yaml:"inline"`

	Host     string `json:"host,omitempty" yaml:"host" xml:"host" toml:"host" validate:"required,hostname|ip"`
	Port     int    `json:"port,omitempty" yaml:"port" xml:"port" toml:"port" validate:"required,min=1,max=65535"`
	Username string `json:"username,omitempty" yaml:"username" xml:"username" toml:"username" validate:"required"`
	Password string `json:"password,omitempty" yaml:"password" xml:"password" toml:"password" validate:"required"`
	Identity string `json:"identity,omitempty" yaml:"identity" xml:"identity" toml:"identity" validate:"required"`
	Pool     int    `json:"pool,omitempty" yaml:"pool" xml:"pool" toml:"pool" validate:"required,min=1"`
}
