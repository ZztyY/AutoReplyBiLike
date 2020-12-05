package main

import (
	"AutoReplyBiLike/api"
	"AutoReplyBiLike/config"
	"fmt"
)

func main() {
	likeCardIdList := api.GetLikeCards(config.SESSDATA)
	for _, v := range likeCardIdList {
		likeList := api.GetLikeUserIds(v, config.SESSDATA)
		fmt.Print(likeList)
	}
}
