package plugin

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/qingniao"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(http *http.Client, logger log.Logger) *qingniao.Sender {
	return qingniao.New(
		qingniao.HttpClient(http),
		qingniao.Logger(logger),
	)
}
