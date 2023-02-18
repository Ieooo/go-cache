package hash

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	m := NewNodeMap(1, func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	})
	m.SetNode("1", "8", "16")
	assert.Equal(t, m.GetNode("9"), "1")
	assert.Equal(t, m.GetNode("161"), "1")
	assert.Equal(t, m.GetNode("80"), "8")
	assert.Equal(t, m.GetNode("81"), "16")
}
