package cache_test

import (
	"log"
	"testing"

	"notchman.tech/masakari-backend/cache"
)

func TestMemcache(t *testing.T) {

	m := cache.NewMemcached("memcached:11211")
	if e := m.Save("sample", []byte("This is the SHIOKAZE limited express bound for Okayama")); e != nil {
		log.Println(e.Error())
	}

	if v, e := m.Get("sample"); e != nil {
		log.Println(e.Error())
	} else if v == nil {
		log.Println("value is empty")
	} else {
		log.Println(string(v))
	}
}
