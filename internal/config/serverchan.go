package config

type Serverchan struct {
	LabelBase `yaml:"inline"`

	Key string `json:"key,omitempty" yaml:"key" xml:"key" toml:"key" validate:"required"`
}
