package main

import (
	"os"

	"github.com/cloud66/cli"
)

func runServiceRestart(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowSubcommandHelp(c)
		os.Exit(2)
	}

	// get stack
	stack := mustStack(c)

	// get serverUID
	var serverUID *string
	flagServer := c.String("server")
	if flagServer != "" {
		server := mustServer(c, *stack, flagServer, false)
		serverUID = &server.Uid
	}

	// get serviceName
	var serviceName *string
	flagService := c.Args()[0]
	if flagService != "" {
		serviceName = &flagService
	}

	asyncId, err := startServiceAction(stack.Uid, serviceName, serverUID, "service_restart")
	if err != nil {
		printFatal(err.Error())
	}
	genericRes, err := endServiceAction(*asyncId, stack.Uid)
	if err != nil {
		printFatal(err.Error())
	}
	printGenericResponse(*genericRes)
	return
}
