package main

import (
	"encoding/json"

	"github.com/pkg/errors"
	"notchman.tech/location-provider/redis"
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
		return ErrorSocketResponse
	}

	// 各アクションケースに応じて処理を行う
	switch {
	case requestObject.Action == "UPDATE_CHARACTER_POS":
		// TODO replace here
		err := setUserPos(requestObject)
		if err != nil {
			return SystemResponse
		}
		r := []byte(`{"action":"SYSTEM_MESSAGE","status":"OK","error": false}`)
		return r
	}

	return SampleResponse
}
