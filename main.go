package main

import (
	"log"
	"os"

	"github.com/levicook/todo-api/httpd"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("todo-api", "0.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"httpd": httpd.CommandFactory,
	}

	status, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(status)
}
