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
		for range time.Tick(10 * time.Second) {
			fmt.Println("Socket Job is called")

			var response ProcessStatusResponse
			response.Action = ACTION_SEND_STATUS
			// TODO:テストデータではなく実際の物にする

			status := ProcessStatus{
				CPU:     11.4514,
				Memory:  11.4514,
				Traffic: 114514,
				Name:    "馬車馬",
			}
			response.Status = status

			res, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
				continue
			}
			m := message{res, ROOM_ID}
			h.broadcast <- m

			//TODO 消せ
			// m = message{SampleResponse, ROOM_ID}
			// h.broadcast <- m
		}
	}()

}
