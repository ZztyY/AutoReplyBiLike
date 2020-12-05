package job

import (
	"github.com/robfig/cron"
)

func InitCron() {
	// 秒、分、时、日、月、星期
	c := cron.New()

	_ = c.AddFunc("0 0-59/1 6-22 * * *", func() {

	})
}
