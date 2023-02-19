package core

import (
	"cache/server/lru"
	"cache/server/value"
	"errors"
)

// local cache
type lCache struct {
	dbs []*lru.Cache
}

func NewLCache(dbNums int, maxCacheSize int64) *lCache {
	var dbs []*lru.Cache
	for i := 0; i < dbNums; i++ {
		dbs = append(dbs, lru.NewCache(maxCacheSize))
	}
	return &lCache{
		dbs: dbs,
	}
}

func (l lCache) Get(key string) (string, error) {
	v := l.dbs[0].Get(key)
	if v != nil {
		return v.(value.StrValue).String(), nil
	}
	return "", errors.New("key is not exists")
}
func (l lCache) Set(key string, val string) error {
	l.dbs[0].Set(key, value.String(val))
	return nil
}
func (l lCache) Del(key string) error {
	l.dbs[0].Del(key)
	return nil
}

func (l lCache) Scan() (map[string]string, error) {
	m := l.dbs[0].Scan()
	resMap := make(map[string]string, len(m))
	for k, v := range m {
		if v != nil {
			resMap[k] = v.(value.StrValue).String()
		}
	}
	return resMap, nil
}
