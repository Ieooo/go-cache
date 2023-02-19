package lru

import (
	"cache/server/value"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLru(t *testing.T) {
	c := NewCache(12)
	c.Set("key1", value.StrValue("value1"))
	assert.True(t, c.Get("key1").(value.StrValue) == value.String("value1"))
	c.Set("key2", value.StrValue("value2"))
	c.Set("key3", value.StrValue("value3"))
	assert.True(t, c.Get("key1") == nil)
	c.Del("key3")
	assert.True(t, c.Get("key2").(value.StrValue) == value.String("value2"))
	assert.True(t, c.Get("key3") == nil)
}

func TestLRUThread(t *testing.T) {
	c := NewCache(1 << 30)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(g int) {
			for j := 0; j < 10000; j++ {
				c.Set("k"+strconv.Itoa(j+g*10000), value.StrValue("v"+strconv.Itoa(j+g*10000)))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 100000; i++ {
		val := c.Get("k" + strconv.Itoa(i))
		if val == nil {
			t.Fatal(i)
		}
		assert.Equal(t, "v"+strconv.Itoa(i), val.(value.StrValue).String())
	}
}
