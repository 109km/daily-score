package autotask

import (
	"fmt"

	"github.com/robfig/cron"
)

func Start() {
	c := cron.New()
	c.AddFunc("08 08 * * *", StartDailySentence)
	fmt.Println("startAutoTasks")
	c.Start()
}
