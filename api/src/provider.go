package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// MasakariBackendとして全ユーザーの位置情報を定期的に返却する
func ProviderJob() {

	// WebSocketに良い感じに流すジョブ
	go func() {
		for range time.Tick(5 * time.Second) {
			fmt.Println("Socket Job is called")

			var response ProcessStatusResponse
			response.Action = ACTION_SEND_STATUS
			var status ServerStatus
			statusByteData, err := getServerStatus()
			if err != nil {
				log.Println(err)
				continue
			}
			json.Unmarshal(statusByteData, &status)

			response.Status = status

			res, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
				continue
			}
			m := message{res, ROOM_ID}
			h.broadcast <- m
		}
	}()

}
