package plugin

type Wrapper struct {
	Qingniao *Qingniao `json:"qingniao,omitempty" yaml:"qingniao" xml:"qingniao" toml:"qingniao" validate:"required"`
}
