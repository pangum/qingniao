package qingniao

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependency(
		newSender,
	)
}
