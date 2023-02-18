package peer

import (
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Nodes struct {
	hash     Hash
	keys     []int
	nodeMap  map[int]string
	replicas int // virtual node counts of one physical node
}

func NewNodes(replicas int, hash Hash) *Nodes {
	return &Nodes{
		hash:     hash,
		nodeMap:  make(map[int]string),
		replicas: replicas,
	}
}

func (n *Nodes) AddNode(keys ...string) {
	for _, key := range keys {
		for i := 0; i < n.replicas; i++ {
			h := n.hash([]byte(key + strconv.Itoa(i)))
			n.keys = append(n.keys, int(h))
			n.nodeMap[int(h)] = key
		}
	}
	sort.Ints(n.keys)
}

func (n *Nodes) GetNode(key string) string {
	if len(n.keys) == 0 {
		return ""
	}

	h := int(n.hash([]byte(key)))
	index := sort.Search(len(n.keys), func(i int) bool {
		return h <= n.keys[i]
	})

	return n.nodeMap[n.keys[index%len(n.keys)]]
}

func (n Nodes) Get(key string) string {
	return ""
}
func (n Nodes) Set(key string, val string) {
}
func (n Nodes) Del(key string) {
}

var _ PeerCache = (*Nodes)(nil)
