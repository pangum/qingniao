package qingniao

import (
	"github.com/goexl/qingniao"
	"github.com/pangum/http"
	"github.com/pangum/logging"
)

// Sender 发送者
type Sender struct {
	*qingniao.Sender
}

func newSender(http *http.Client, logger logging.Logger) *Sender {
	return &Sender{
		Sender: qingniao.New(
			qingniao.HttpClient(http.Client),
			qingniao.Logger(logger),
		),
	}
}
