package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Like struct {
	Mid       int    // B站用户id
	TimeStamp string // 时间戳
}

// 获取获赞内容id列表
func GetLikeCards(sess string) []int {
	var cardIdList []int
	var body map[string]interface{}
	req, _ := http.NewRequest("GET", "https://api.bilibili.com/x/msgfeed/like", strings.NewReader(""))
	sessData := http.Cookie{Name: "SESSDATA", Value: sess, Domain: "api.bilibili.com", HttpOnly: true}
	req.AddCookie(&sessData)
	resp, _ := http.DefaultClient.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(res, &body)
	if err != nil {
		panic(err)
	}
	items := body["data"].(map[string]interface{})["total"].(map[string]interface{})["items"].([]interface{})
	for _, v := range items {
		cardId, _ := json.Marshal(v.(map[string]interface{})["id"])
		cardIdInt, _ := strconv.Atoi(string(cardId))
		cardIdList = append(cardIdList, cardIdInt)
	}
	return cardIdList
}

// 根据内容id列表获取点赞用户
func GetLikeUserIds(cardId int, sess string) []Like {
	var likeList []Like
	var body map[string]interface{}
	req, _ := http.NewRequest("GET", "https://api.bilibili.com/x/msgfeed/like_detail?card_id="+strconv.Itoa(cardId), strings.NewReader(""))
	sessData := http.Cookie{Name: "SESSDATA", Value: sess, Domain: "api.bilibili.com", HttpOnly: true}
	req.AddCookie(&sessData)
	resp, _ := http.DefaultClient.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(res, &body)
	if err != nil {
		panic(err)
	}
	list := body["data"].(map[string]interface{})["items"].([]interface{})
	for _, k := range list {
		var like Like
		user, _ := k.(map[string]interface{})["user"]
		mid, _ := json.Marshal(user.(map[string]interface{})["mid"])
		likeTime, _ := json.Marshal(k.(map[string]interface{})["like_time"])
		like.Mid, _ = strconv.Atoi(string(mid))
		like.TimeStamp = string(likeTime)
		likeList = append(likeList, like)
	}
	return likeList
}

// 发送私信
func SendMessage(receiverMid int, senderMid int, sess string, content string) error {
	var r http.Request
	err := r.ParseForm()
	if err != nil {
		return err
	}
	r.Form.Add("msg[receiver_id]", strconv.Itoa(receiverMid))
	r.Form.Add("msg[sender_uid]", strconv.Itoa(senderMid))
	r.Form.Add("msg[receiver_type]", "1")
	r.Form.Add("msg[msg[msg_type]]", "1")
	r.Form.Add("msg[msg[content]]", "{\"content\":"+content+"}")
	bodyStr := strings.TrimSpace(r.Form.Encode())
	req, _ := http.NewRequest("POST", "http://api.vc.bilibili.com/web_im/v1/web_im/send_msg", strings.NewReader(bodyStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	sessData := http.Cookie{Name: "SESSDATA", Value: sess, Domain: "api.bilibili.com", HttpOnly: true}
	req.AddCookie(&sessData)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
