package main

import (
	"fmt"
	"log"
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

func get(c *cli.Context) error {
	key := c.Args().Get(0)
	fmt.Println(key)
	return nil
}

func set(c *cli.Context) error {
	key := c.Args().Get(0)
	value := c.Args().Get(1)
	fmt.Println(key, value)
	return nil
}
func delete(c *cli.Context) error {
	key := c.Args().Get(0)
	fmt.Println(key)
	return nil
}
