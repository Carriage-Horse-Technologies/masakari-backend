package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"notchman.tech/masakari-backend/cache"
)

// GoogleSearchAPIに検索クエリを投げる
func fetchStatusAsByte() (result_json []byte, err error) {
	const (
		url = "https://jqyf37xokyuyzxqharfzdelepq0lbrjw.lambda-url.ap-northeast-1.on.aws/"
	)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var status ServerStatus
	json.Unmarshal(resBody, &status)

	return resBody, nil
}

func fetchStatus() (result_json ServerStatus, err error) {
	const (
		url = "https://jqyf37xokyuyzxqharfzdelepq0lbrjw.lambda-url.ap-northeast-1.on.aws/"
	)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var status ServerStatus
	json.Unmarshal(resBody, &status)
	return status, nil
}

func getServerStatus() ([]byte, error) {
	//キャッシュの作成
	memcachedUrl := os.Getenv("MEMCACHED_URL")
	if len(memcachedUrl) == 0 {
		return nil, fmt.Errorf("memcached url is not found in env")
	}
	cacheManager := cache.NewMemcached(memcachedUrl)
	cache_item, err := cacheManager.Get(KEY_CACHE_STATUS)
	if len(cache_item) != 0 && err == nil {
		fmt.Println("cached")
		LoggingHTTPError(200, fmt.Errorf("OK"))
		return cache_item, nil
	}
	fmt.Println(cache_item)

	result_json, err := fetchStatusAsByte()
	if err != nil {
		return nil, err
	}
	//キャッシュをセット
	cacheManager.SaveFor(30*time.Second, KEY_CACHE_STATUS, result_json)
	return result_json, nil
}
