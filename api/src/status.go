package main

import (
	"fmt"
	"net/http"
	"time"

	"notchman.tech/masakari-backend/cache"
)

func getServerStatus(w http.ResponseWriter, r *http.Request) {
	//キャッシュの作成
	cacheManager := cache.NewMemcached("memcached:11211")

	if r.Method != "GET" {
		status := 405
		w.WriteHeader(status)
		err := fmt.Errorf("This endpoint allows only GET method but recieve %s", r.Method)
		w.Write([]byte(err.Error()))

		LoggingHTTPError(status, err)
		return
	}

	cache_item, err := cacheManager.Get(KEY_CACHE_STATUS)
	if len(cache_item) != 0 && err == nil {
		fmt.Println("cached")
		w.Header().Set("Content-Type", "application/json")
		w.Write(cache_item)
		LoggingHTTPError(200, fmt.Errorf("OK"))
		return
	}
	fmt.Println(cache_item)

	//TODO: 結果から検索エンジンID等の情報を削除
	result_json, err := fetchStatus()
	if err != nil {
		status := 500
		w.WriteHeader(status)
		w.Write([]byte("Internal Server Error! Details in the log."))
		LoggingHTTPError(status, err)
		return
	}
	//キャッシュをセット
	cacheManager.SaveFor(30*time.Second, KEY_CACHE_STATUS, result_json)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result_json)
	LoggingHTTPError(200, fmt.Errorf("OK"))
}
