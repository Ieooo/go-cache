package core

import (
	"cache/server/config"
	"hash/crc32"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkCache(b *testing.B) {
	InitCache([]string{"127.0.0.1:9000"}, crc32.ChecksumIEEE)
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		Set("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	}
}

func TestCacheThread(t *testing.T) {
	config.InitConfig("../etc/conf.yaml")
	InitCache([]string{"127.0.0.1:9000"}, crc32.ChecksumIEEE)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(g int) {
			for j := 0; j < 10000; j++ {
				Set("k"+strconv.Itoa(j+g*10000), "v"+strconv.Itoa(j+g*10000))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 100000; i++ {
		val, err := Get("k" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err, i)
		}
		assert.Equal(t, "v"+strconv.Itoa(i), val)
	}
}
