package core

import (
	"cache/server/lru"
	"cache/server/value"
)

var lruCache = lru.NewCache(1024)

func Set(key, val string) {
	v := value.String(val)
	lruCache.Set(key, v)
}

func Get(key string) string {
	v := lruCache.Get(key)
	if v != nil {
		return v.(value.StrValue).String()
	}
	return ""
}

func Delete(key string) {
	lruCache.Del(key)
}
