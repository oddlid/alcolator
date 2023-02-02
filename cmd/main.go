package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// The mother of all contexts. Should be passed on to all levels in the app, and other contexts should be derived from this one.
	// SIGINT is for ctrl-c in a terminal
	// SIGTERM is what Docker and Kubernetes uses to shutdown containers
	// SIGQUIT is just there to be on the safe side
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()
	app := newApp()

	if err := app.RunContext(ctx, os.Args); err != nil {
		if !errors.Is(err, context.Canceled) {
			zlog.Fatal().Err(err).Send()
		}
	}
}
