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

func setUserPos(request Request) (err error) {
	// convert object to json string
	objStr, err := convertJsonToStr(request)
	if err != nil {
		return err
	}
	err = redis.Update(savePathBuilder(request.UserId), objStr)
	if err != nil {
		return err
	}
	return nil
}

func handler(s []byte) []byte {
	var requestObject Request
	if err := json.Unmarshal(s, &requestObject); err != nil {
		log.Println(err)
		return ErrorSocketResponse
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

			return ErrorSocketResponse
		}
		return res
	case requestObject.Action == ACTION_RECV_GPT:
		r := []byte(`{"action":"ACTION_RECV_GPT","user_id":"examper-user-id","message":"hogehoge"}`)
		return r
	case requestObject.Action == ACTION_SEND_STATUS:
		r := []byte(`{"name":"デスマTV","cpu":11.4514,"memory":11.514,"traffic":114514}`)
		return r
	default:
		return ErrorSocketResponse
	}

}
