package test

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkCache(b *testing.B) {
	b.ResetTimer()
	client := http.DefaultClient
	for i := 0; i < 100000; i++ {
		query := url.Values{}
		query.Add("k", "key"+strconv.Itoa(i))
		query.Add("v", "val"+strconv.Itoa(i))
		req, err := http.NewRequest("GET", "http://127.0.0.1:9000/cache/set", nil)
		if err != nil {
			log.Fatal(err)
		}
		req.URL.RawQuery = query.Encode()
		client.Do(req)
	}
}
func BenchmarkCacheTread(b *testing.B) {
	b.ResetTimer()
	wg := sync.WaitGroup{}
	for j := 0; j < 2; j++ {
		wg.Add(1)
		go func(base int) {
			client := http.DefaultClient
			for i := 0; i < 50000; i++ {
				query := url.Values{}
				query.Add("k", "k"+strconv.Itoa(i+base*10000))
				query.Add("v", "v"+strconv.Itoa(i+base*10000))
				req, err := http.NewRequest("GET", "http://127.0.0.1:9000/cache/set", nil)
				if err != nil {
					log.Fatal(err)
				}
				req.URL.RawQuery = query.Encode()
				client.Do(req)
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
}
