package hash

import (
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type NodeMap struct {
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
	for _, key := range keys {
		for i := 0; i < n.replicas; i++ {
			h := n.hash([]byte(key + strconv.Itoa(i)))
			n.nodeHash = append(n.nodeHash, int(h))
			n.nodeMap[int(h)] = key
		}
	}
	sort.Ints(n.nodeHash)
}

// get the node which stored the value
func (n *NodeMap) GetNode(key string) string {
	if len(n.nodeHash) == 0 {
		return ""
	}

	h := int(n.hash([]byte(key)))
	index := sort.Search(len(n.nodeHash), func(i int) bool {
		return h <= n.nodeHash[i]
	})

	return n.nodeMap[n.nodeHash[index%len(n.nodeHash)]]
}
