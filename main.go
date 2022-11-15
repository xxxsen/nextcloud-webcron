package main

import (
	"nextcloud-webcron/crontask"

	"github.com/xxxsen/common/logger"
	flag "github.com/xxxsen/envflag"
	"go.uber.org/zap"
)

var expression = flag.String("expression", "*/5 * * * *", "cron expression")
var url = flag.String("url", "http://127.0.0.1/cron.php", "nextcloud url")
var runOnStart = flag.Bool("run_on_start", true, "run cron task when service start")
var loglv = flag.String("log_level", "debug", "log level")

func main() {
	flag.Parse()

	log := logger.Init("", *loglv, 1, 1, 1, true)
	log.With(zap.String("url", *url), zap.String("expression", *expression), zap.Bool("run_on_start", *runOnStart)).
		Info("crontask init")
	tk, err := crontask.New(
		crontask.WithCronExpression(*expression),
		crontask.WithRunOnStart(*runOnStart),
		crontask.WithURL(*url),
	)
	if err != nil {
		log.With(zap.Error(err)).Fatal("create task fail")
	}

	if err := tk.Run(); err != nil {
		log.With(zap.Error(err)).Error("run task fail")
	}
}
