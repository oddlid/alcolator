package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

var (
	version  string // to be set by the linker
	compiled string // to be set by the linker
	commitID string // to be set by the linker
	zlog     = zerolog.New(os.Stdout)
)

func getCompileTime() time.Time {
	ts, err := time.Parse(time.RFC3339, compiled)
	if err != nil {
		return time.Time{} // time.Now() might be better, but zero time is easier to see that is wrong
	}
	return ts
}

func newApp() cli.App {
	return cli.App{
		Version:   version,
		Compiled:  getCompileTime(),
		Name:      "Alcolator",
		Copyright: "(C) 2023 Odd Eivind Ebbesen",
		Authors: []*cli.Author{
			{
				Name:  "Odd E. Ebbesen",
				Email: "oddebb@gmail.com",
			},
		},
		Usage: "Calculate drink relations",
		// Before:   nil,
		Flags:    []cli.Flag{},
		Commands: []*cli.Command{},
	}
}
