package main

import (
	"encoding/json"
	"log"

	"github.com/pkg/errors"
	"notchman.tech/masakari-backend/redis"
)

func count(target string) (value int, err error) {
	err = redis.AddValue(target)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to add connection")
	}

	value, err = redis.DeclValue(target)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to decl connection")
	}
	return
}

func handler(s []byte) []byte {
	var requestObject Request
	if err := json.Unmarshal(s, &requestObject); err != nil {
		log.Println(err)
		return errorResponseFactory("faile to parse json", 503, err.Error())

	}

	// 各アクションケースに応じて処理を行う
	switch {
	case requestObject.Action == ACTION_CHAT_MESSAGE:
		res, err := json.Marshal(ChatResponse{
			UserId:  requestObject.UserId,
			Action:  requestObject.Action,
			Message: requestObject.Message,
			Name:    requestObject.Name,
		})
		if err != nil {
			log.Println(err)

			return errorResponseFactory("faile to send message", 503, "data is not json object")

		}
		return res
	case requestObject.Action == ACTION_RECV_GPT:
		r := []byte(`{"action":"ACTION_RECV_GPT","user_id":"examper-user-id","message":"hogehoge"}`)
		return r
	case requestObject.Action == ACTION_SEND_STATUS:
		r := []byte(`{"name":"デスマTV","cpu":11.4514,"memory":11.514,"traffic":114514}`)
		return r
	case requestObject.Action == ACTION_RECV_MASAKARI:

		//メッセージのスコアを計算してサーバーを攻撃する（こちらは特に待つ必要がないので別で処理）
		attackServer(requestObject.Message)

		//メッセージをGPTへ投げる
		msg, err := fetchGPTMessage(requestObject.Message)

		if err != nil {
			return errorResponseFactory("faile to send message to gpt server", 503, err.Error())
		}
		//メッセージを返却する
		return messageResponseFactory(msg)

	default:
		return errorResponseFactory("faile to execute action", 503, "no such action type")
	}

}
