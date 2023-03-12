package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"notchman.tech/location-provider/redis"
)

// MasakariBackendとして全ユーザーの位置情報を定期的に返却する
func ProviderJob() {

	// WebSocketに良い感じに流すジョブ
	go func() {
		for range time.Tick(10 * time.Second) {
			fmt.Println("Socket Job is called")

			userList, err := redis.SMEMBERS(CONNECTION_PATH)
			if err != nil {
				log.Println(err)
				continue
			}
			var response UserPosReponse
			response.Action = "UPDATE_CHARACTER_POS"
			response.Characters = []User{}

			fmt.Println(userList)
			for _, userId := range userList {
				// 各ユーザーのスコアを取得
				userData, err := redis.GetValue(savePathBuilder(userId))
				fmt.Println(userData)
				if err != nil {
					log.Println(err)
					continue
				}
				var responseObj Request
				err = json.Unmarshal([]byte(userData), &responseObj)
				if err != nil {
					log.Println(err)
					continue
				}
				resData := User{
					UserId:    responseObj.UserId,
					UserImage: "https://exapmle.com",
					PosX:      responseObj.PosX,
					PosY:      responseObj.PosY,
				}
				response.Characters = append(response.Characters, resData)

			}

			res, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println("Current target num:", len(userList))
			m := message{res, ROOM_ID}
			h.broadcast <- m

			//TODO 消せ
			m = message{SampleResponse, ROOM_ID}
			h.broadcast <- m
		}
	}()

}
