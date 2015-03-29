package httpd

import (
	"log"
	"net/http"

	"github.com/levicook/todo-api/env"
	"github.com/levicook/todo-api/httpd/routes"
	"github.com/mitchellh/cli"
)

func CommandFactory() (cli.Command, error) {
	return command{}, nil
}

type command struct{}

// Help should return long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (command) Help() string {
	return ""
}

// Run should run the actual command with the given CLI instance and
// command-line arguments. It should return the exit status when it is
// finished.
func (command) Run(args []string) int {
	// open database connections
	var (
		err  error
		port string
	)

	if port, err = env.PORT(); err != nil {
		log.Println(err)
		return 1
	}

	bind := ":" + port
	log.Printf("listening on %v", bind)

	if err = http.ListenAndServe(bind, routes.Routes.Handler()); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

// Synopsis should return a one-line, short synopsis of the command.
// This should be less than 50 characters ideally.
func (command) Synopsis() string {
	return "runs foreground server for http requests"
}
