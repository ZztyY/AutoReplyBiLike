package main

import (
	"AutoReplyBiLike/api"
	"AutoReplyBiLike/config"
)

func main() {
	//likeCardIdList := api.GetLikeCards(config.SESSDATA)
	//for _, v := range likeCardIdList {
	//	likeList := api.GetLikeUserIds(v, config.SESSDATA)
	//	fmt.Print(likeList)
	//}
	api.SendMessage(320170411, 216174037, config.SESSDATA, "test")
}
