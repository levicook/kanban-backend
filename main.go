package main

import (
	"log"
	"os"

	"github.com/levicook/kanban-backend/httpd"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("kanban-backend", "0.0.0")
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
