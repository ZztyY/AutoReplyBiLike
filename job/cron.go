package job

import (
	"AutoReplyBiLike/api"
	"AutoReplyBiLike/config"
	"AutoReplyBiLike/util"
	"github.com/robfig/cron"
	"log"
	"time"
)

func InitCron() {
	// 秒、分、时、日、月、星期
	c := cron.New()

	_ = c.AddFunc("0 0-59/1 * * * *", func() {
		// todo
		log.Println(util.GetCurrentTimeStr())
	})
	// 粉丝定时消息
	_ = c.AddFunc("0 0 8 * * *", func() {
		midList := api.GetFollowers(216174037, config.SESSDATA)
		for _, k := range midList {
			time.Sleep(1000)
			err := api.SendMessage(k, api.GetAccountMid(config.SESSDATA), config.SESSDATA, "我是机器人：%……&%&……%*&￥*%……？")
			if err != nil {
				panic(err)
			}
			log.Print(util.GetCurrentTimeStr() + " Send to" + util.IntToStr(k) + "success")
		}
	})
	c.Start()
}
