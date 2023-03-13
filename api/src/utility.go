package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"notchman.tech/masakari-backend/redis"
)

func checkRoomId(roomId string) (bool, error) {
	// DBの初期化
	redisPath := os.Getenv("REDIS_PATH")
	log.Println(redisPath)
	client, err := redis.New(redisPath)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	defer client.Close()

	//値の存在チェック
	result := client.SIsMember(ID_PATH, roomId)
	if result.Err() != nil {
		//redisのエラー
		fmt.Println(err)
		return false, err
	} else if result.Val() {
		return true, nil
	} else {
		return false, nil
	}
}

func savePathBuilder(userId string) (path string) {
	return BASE_POS_PATH + userId
}
func convertJsonToStr(jsonObj interface{}) (jsonStr string, err error) {
	serializedObj, err := json.Marshal(jsonObj)
	if err != nil {
		log.Println(errors.WithStack(err))
		jsonStr = ""
		return

	}
	jsonStr = string(serializedObj)
	return
}

func convertStrToJson(jsonStr string, jsonObj *(interface{})) (err error) {
	err = json.Unmarshal([]byte(jsonStr), jsonObj)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func errorResponseFactory(name string, code int, msg string) []byte {
	errRes := ErrorObject{
		Action: ACTION_ERROR,
		Name:   name,
		Code:   code,
		Msg:    msg,
	}
	res, err := json.Marshal(errRes)
	if err != nil {
		return FatalErrorResponse
	}
	return res
}
