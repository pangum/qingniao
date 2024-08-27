package config

type LabelBase struct {
	Label string `default:"default" json:"label,omitempty" yaml:"label" xml:"label" toml:"label"`
}
