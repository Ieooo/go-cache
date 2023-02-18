package core

import (
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
