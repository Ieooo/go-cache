package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const BasePath = "/cache"

// remote cache
type rCache struct {
	baseUrl string
	c       *http.Client
}

func NewRCache(endpoint string) *rCache {
	path := "http://" + endpoint
	return &rCache{
		baseUrl: path,
		c:       http.DefaultClient,
	}
}

func (r rCache) Get(key string) (string, error) {
	res, err := httpGet(r.c, r.baseUrl+BasePath+"/get", key, nil)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
func (r rCache) Set(key string, val string) error {
	_, err := httpGet(r.c, r.baseUrl+BasePath+"/set", key, val)
	return err
}
func (r rCache) Del(key string) error {
	_, err := httpGet(r.c, r.baseUrl+BasePath+"/del", key, nil)
	return err
}
func (r rCache) Scan() (map[string]string, error) {
	res, err := httpPost(r.c, r.baseUrl+BasePath+"/scan", nil)
	var m map[string]string
	if len(res) == 0 {
		return nil, nil
	}
	if err := json.Unmarshal(res, m); err != nil {
		return nil, err
	}
	return m, err
}

func httpGet(c *http.Client, u, key string, val interface{}) ([]byte, error) {
	query := url.Values{}
	query.Add("k", key)
	if val != nil {
		query.Add("v", val.(string))
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()
	r, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r.Body)
}

func httpPost(c *http.Client, u string, m map[string]string) ([]byte, error) {
	var b []byte
	if m != nil {
		var err error
		b, err = json.Marshal(m)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest("POST", u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	r, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r.Body)
}
