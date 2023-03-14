package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"notchman.tech/masakari-backend/redis"
)

func buildGptMessages(messages []string) (res []GptMessage) {
	for _, msg := range messages {
		res = append(res, GptMessage{Role: "user", Content: msg})
	}
	return res
}

const ADJUST_SIZE = 40

func buildPresetMessages() ([]string, error) {
	//調教用データ
	presetMessages := []string{"あなたはベテランプログラマーです", "あなたはクソコードが嫌いです", "あなたはクソコードに対して厳しい意見を言わなければなりません", "ただし、30文字以内で話せ"}

	//キャッシュの作成
	messages, err := redis.LRange(KEY_GPT_WORD, 0, math.MaxInt)
	if err != nil {
		return nil, err
	}
	presetMessages = append(presetMessages, messages...)
	// fmt.Println("debug", presetMessages)
	return presetMessages, nil
}

func fetchGPTMessage(msg string) (gptMessage string, err error) {
	const (
		method      = "POST"
		url         = "https://api.openai.com/v1/chat/completions"
		contentType = "application/json"
	)

	// -------------------------------------------
	// Create request and set headers
	// -------------------------------------------

	presetMessages, err := buildPresetMessages()
	if err != nil {
		return
	}

	presetMessages = append(presetMessages, msg)

	// ここのエラーはもみ消す
	_ = redis.RPush(KEY_GPT_WORD, msg)

	body, err := json.Marshal(GptRequestBody{
		Messages: buildGptMessages(presetMessages),
		Model:    "gpt-3.5-turbo",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bytes.NewBuffer(body)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	accessToken := os.Getenv("GPT_API_KEY")
	if len(accessToken) == 0 {
		return "", fmt.Errorf("failed to get access token")
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	// -------------------------------------------
	// Send http POST request
	// -------------------------------------------

	var (
		client = &http.Client{}
	)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

		return
	}
	defer res.Body.Close()

	// // -------------------------------------------
	// // Check http status code
	// // -------------------------------------------

	// if res.StatusCode != http.StatusCreated {
	// 	return
	// }

	// -------------------------------------------
	// Decode response to JSON
	// -------------------------------------------
	// fmt.Println(res)
	resBody, err := ioutil.ReadAll(res.Body)
	// fmt.Println(resBody)
	if err != nil {
		return
	}
	var post GptResponseBody
	json.Unmarshal(resBody, &post)

	if err != nil {
		fmt.Println(err)

		return
	}

	// -------------------------------------------
	// Show results
	// -------------------------------------------
	// fmt.Println(post)
	gptMessage = post.Choices[0].Message.Content
	// if len(gptMessage) < ADJUST_SIZE {
	// 	for i := 0; i < ADJUST_SIZE-len(msg); i++ {
	// 		// iの値が0から9まで変化していく
	// 		gptMessage += "　"
	// 	}
	// }
	return gptMessage, nil
}
