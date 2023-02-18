package core

import (
	"cache/server/config"
	"cache/server/lru"
	"cache/server/peer"
)

var cache *Cache

type Cache struct {
	Dbs   []*lru.Cache
	Peers []peer.PeerCache
}

func InitCache(endpoints []string) {
	cache = NewCache(endpoints)
}

func NewCache(endpoints []string) *Cache {
	var dbs []*lru.Cache
	var peers []peer.PeerCache
	// init dbs
	for i := 0; i < 16; i++ {
		dbs = append(dbs, lru.NewCache(1024))
	}
	// init nodes
	for _, v := range endpoints {
		if v == config.Conf.IP+config.Conf.Port {
			peers = append(peers, peer.NewNodes(16, func(data []byte) uint32 {
				return 0
			}))
		} else {
			peers = append(peers, peer.NewPeerClient(v))
		}
	}
	return &Cache{
		Dbs:   dbs,
		Peers: peers,
	}
}

func Set(key, val string) {

}

func Get(key string) string {
	return ""
}

func Del(key string) {
}
