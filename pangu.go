package qingniao

import (
	"github.com/pangum/pangu"
	"github.com/pangum/qingniao/internal/plugin"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		ctor.New,
	).Build().Build().Apply()
}
