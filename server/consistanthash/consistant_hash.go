package hash

import (
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32

type NodeMap struct {
	sync.Mutex

	hash     Hash
	nodeHash []int
	nodeMap  map[int]string
	replicas int // virtual node counts of physical node
}

func NewNodeMap(replicas int, hash Hash) *NodeMap {
	return &NodeMap{
		hash:     hash,
		nodeMap:  make(map[int]string),
		replicas: replicas,
	}
}

// set nodes
func (n *NodeMap) SetNode(keys ...string) {
	n.Lock()
	defer n.Unlock()

	for _, key := range keys {
		for i := 0; i < n.replicas; i++ {
			h := n.hash([]byte(key + strconv.Itoa(i)))
			n.nodeHash = append(n.nodeHash, int(h))
			n.nodeMap[int(h)] = key
		}
	}
	sort.Ints(n.nodeHash)
}

// remove node
func (n *NodeMap) RemoveNode(key string) {
	n.Lock()
	defer n.Unlock()

	for i := 0; i < n.replicas; i++ {
		h := int(n.hash([]byte(key + strconv.Itoa(i))))
		index := sort.Search(len(n.nodeHash), func(i int) bool {
			return n.nodeHash[i] >= h
		})
		if n.nodeHash[index] == h {
			n.nodeHash = append(n.nodeHash[:index], n.nodeHash[index+1:]...)
		}
		delete(n.nodeMap, int(h))
	}

}

// get the node which stored the value
func (n *NodeMap) GetNode(key string) string {
	n.Lock()
	defer n.Unlock()

	if len(n.nodeHash) == 0 {
		return ""
	}

	h := int(n.hash([]byte(key)))
	index := sort.Search(len(n.nodeHash), func(i int) bool {
		return h <= n.nodeHash[i]
	})

	return n.nodeMap[n.nodeHash[index%len(n.nodeHash)]]
}
