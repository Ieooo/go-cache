package core

import (
	"cache/server/config"
	hash "cache/server/consistanthash"
	"errors"
)

var cache *Cache

type Storage interface {
	Get(key string) (string, error)
	Set(key, val string) error
	Del(key string) error
}

type Cache struct {
	nodeMap  *hash.NodeMap
	storages map[string]Storage
}

func InitCache(endpoints []string, hashFunc hash.Hash) {
	cache = NewCache(endpoints, hashFunc)
}

func NewCache(endpoints []string, hashFunc hash.Hash) *Cache {
	// init cache storage
	storages := make(map[string]Storage, len(endpoints))
	for _, v := range endpoints {
		// local node use local cache, remote cache use remote cache
		if v == config.Conf.IP+config.Conf.Port {
			storages[v] = NewLCache(16, 1024)

		} else {
			storages[v] = NewRCache(v)
		}
	}

	// init nodeMap
	nodes := hash.NewNodeMap(10, hashFunc)
	nodes.SetNode(endpoints...)

	return &Cache{
		nodeMap:  nodes,
		storages: storages,
	}
}

func Set(key, val string) error {
	node := cache.nodeMap.GetNode(key)
	if storage, ok := cache.storages[node]; ok {
		return storage.Set(key, val)
	}
	return errors.New("key is expired")
}

func Get(key string) (string, error) {
	node := cache.nodeMap.GetNode(key)
	if storage, ok := cache.storages[node]; ok {
		return storage.Get(key)
	}
	return "", errors.New("key is expired")
}

func Del(key string) error {
	node := cache.nodeMap.GetNode(key)
	if storage, ok := cache.storages[node]; ok {
		return storage.Del(key)
	}
	return errors.New("key is expired")
}
