package crontask

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/robfig/cron/v3"
	"github.com/xxxsen/common/errs"
	"github.com/xxxsen/common/logutil"
	"go.uber.org/zap"
)

type Crontask struct {
	c *config
}

func New(opts ...Option) (*Crontask, error) {
	c := &config{}
	for _, opt := range opts {
		opt(c)
	}
	return &Crontask{c: c}, nil
}

func (c *Crontask) Run() error {
	if c.c.runOnStart {
		c.run()
	}
	cr := cron.New()
	if _, err := cr.AddFunc(c.c.cronexp, c.run); err != nil {
		return errs.Wrap(errs.ErrServiceInternal, "add cron func fail", err)
	}
	cr.Run()
	return nil
}

func (c *Crontask) run() {
	logger := logutil.GetLogger(context.Background()).With(zap.String("url", c.c.url))
	defer func() {
		if err := recover(); err != nil {
			logger.With(zap.Any("err", err), zap.String("stack", string(debug.Stack()))).
				Error("task panic")
		}
	}()

	req, err := http.NewRequest(http.MethodGet, c.c.url, nil)
	if err != nil {
		logger.With(zap.Error(err)).Error("create http request fail")
		return
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.With(zap.Error(err)).Error("do http request fail")
		return
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		logger.With(zap.Int("status_code", rsp.StatusCode)).Error("call response fail")
		return
	}
	raw, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		logger.With(zap.Error(err)).Error("read response fail")
		return
	}
	msg := &CronMessage{}
	if err := json.Unmarshal(raw, msg); err != nil {
		logger.With(zap.Error(err), zap.String("rspdata", string(raw))).Error("decode nextcloud cron response fail")
		return
	}
	if !strings.EqualFold(msg.Status, msgCronSucc) {
		logger.With(zap.Any("msg", *msg)).Error("exec cron task fail")
		return
	}
	logger.Debug("exec cron task succ")
}
