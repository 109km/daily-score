package autotask

import (
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

var cr *cron.Cron

func Start() {
	cr = cron.New()
	cr.AddFunc("00 00 08 * * *", StartDailySentence)
	logs.Info("Auto tasks are started.")
	cr.Start()
}

func Stop() {
	cr.Stop()
}
