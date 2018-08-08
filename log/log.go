package log

import (
	"microservice/service"
	"time"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		//TODO 后续添加
		_ = mw.logger.Log()
	}(time.Now())
	return
}
