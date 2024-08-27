package plugin

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/qingniao"
	"github.com/pangum/pangu"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(config *pangu.Config, http *http.Client, logger log.Logger) (sender qingniao.Sender, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		sender = c.new(wrapper.Qingniao, http, logger)
	}

	return
}

func (c *Constructor) new(config *Qingniao, http *http.Client, logger log.Logger) qingniao.Sender {
	builder := qingniao.New().Http(http).Logger(logger)
	if nil != config.Sms {
		sms := config.Sms
		sb := builder.Sms()
		for _, cache := range sms.Chuangcache {
			sb = sb.Label(cache.Label)
			sb.Chuangcache(cache.Ak, cache.Sk).Build()
		}

		builder = sb.Build()
	}

	if nil != config.Email {
		email := config.Email
		eb := builder.Email()
		for _, smtp := range email.Smtp {
			eb = eb.Label(smtp.Label)
			sb := eb.Smtp(smtp.Host, smtp.Port).Pool(smtp.Pool)
			sb = sb.Identity(smtp.Identity).Username(smtp.Username).Password(smtp.Password)
			eb = sb.Build()
		}

		builder = eb.Build()
	}

	if nil != config.Wechat {
		wechat := config.Wechat
		wb := builder.Wechat()
		for _, server := range wechat.Serverchan {
			wb = wb.Label(server.Label)
			wb.Serverchan(server.Key)
		}

		builder = wb.Build()
	}

	return builder.Build()
}
