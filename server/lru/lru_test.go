package lru

import (
	"cache/server/value"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLru(t *testing.T) {
	c := NewCache(12)
	c.Set("key1", value.StrValue("value1"))
	c.Set("key2", value.StrValue("value2"))
	c.Set("key3", value.StrValue("value3"))
	assert.True(t, c.Get("key1") == nil)
	c.Del("key3")
	assert.True(t, c.Get("key2").(value.StrValue) == value.String("value2"))
	assert.True(t, c.Get("key3") == nil)
}
