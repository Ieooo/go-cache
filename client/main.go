package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "cache-cli",
		Usage:       "usage",
		Description: "A tool to connect to go-cache",
		Commands: []*cli.Command{
			cmdSet(),
			cmdGet(),
			cmdDel(),
			cmdScan(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func cmdSet() *cli.Command {
	return &cli.Command{
		Name:   "set",
		Usage:  "Set value by key",
		Action: set,
	}
}

func cmdGet() *cli.Command {
	return &cli.Command{
		Name:   "get",
		Usage:  "Get value according to key",
		Action: get,
	}
}
func cmdDel() *cli.Command {
	return &cli.Command{
		Name:   "del",
		Usage:  "Del key and value",
		Flags:  []cli.Flag{},
		Action: delete,
	}
}

func cmdScan() *cli.Command {
	return &cli.Command{
		Name:   "scan",
		Usage:  "scan all value",
		Flags:  []cli.Flag{},
		Action: scan,
	}
}

func get(c *cli.Context) error {
	params := map[string]string{
		"k": c.Args().Get(0),
	}
	client := NewClient()
	v := client.Get("http://127.0.0.1:9000/cache/"+"get", params)
	fmt.Println(string(v))
	return nil
}

func set(c *cli.Context) error {
	params := map[string]string{
		"k": c.Args().Get(0),
		"v": c.Args().Get(1),
	}
	client := NewClient()
	client.Get("http://127.0.0.1:9000/cache/"+"set", params)
	return nil
}
func delete(c *cli.Context) error {
	params := map[string]string{
		"k": c.Args().Get(0),
	}
	client := NewClient()
	client.Get("http://127.0.0.1:9000/cache/"+"del", params)
	return nil
}

func scan(c *cli.Context) error {
	client := NewClient()
	res := client.Post("http://127.0.0.1:9000/cache/"+"scan", nil)
	fmt.Println(string(res))
	return nil
}

type client struct {
	httpClient http.Client
}

func NewClient() *client {
	return &client{
		httpClient: *http.DefaultClient,
	}
}

func (h *client) Get(path string, params map[string]string) []byte {
	query := url.Values{}
	for k, v := range params {
		query.Add(k, v)
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.URL.RawQuery = query.Encode()
	res, err := h.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func (h *client) Post(path string, params map[string]string) []byte {
	var bs []byte
	if params != nil {
		var err error
		bs, err = json.Marshal(params)
		if err != nil {
			fmt.Println(err)
		}
	}
	req, err := http.NewRequest("POST", path, bytes.NewReader(bs))
	if err != nil {
		fmt.Println(err)
	}
	res, err := h.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if res.Body == nil {
		return nil
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
