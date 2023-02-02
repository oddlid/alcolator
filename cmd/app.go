package main

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

const (
	envLogLevel   = `LOG_LEVEL`
	cmdAlcAmount  = `alcoholamount`
	optUnit       = `unit`
	optPercentage = `percentage`
	optVolume     = `volume`
)

var (
	version  string // to be set by the linker
	compiled string // to be set by the linker
	commitID string // to be set by the linker
)

func getCompileTime() time.Time {
	ts, err := time.Parse(time.RFC3339, compiled)
	if err != nil {
		return time.Time{} // time.Now() might be better, but zero time is easier to see that is wrong
	}
	return ts
}

func newApp() *cli.App {
	return &cli.App{
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
		Usage:  "Calculate drink relations",
		Before: logSetup,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    optLogLevel,
				Aliases: []string{"l"},
				Value:   zerolog.InfoLevel.String(),
				Usage:   fmt.Sprintf("Log `level` (options: %s)", getLevels()),
				EnvVars: []string{envLogLevel},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   cmdAlcAmount,
				Usage:  "Get the amount of alcohol in a drink",
				Action: getAlcoholAmount,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  optUnit,
						Usage: "Measurement unit",
						Value: "ml",
					},
					&cli.Float64Flag{
						Name:  optPercentage,
						Usage: "Alcohol percentage of drink",
					},
					&cli.Float64Flag{
						Name:  optVolume,
						Usage: "Volume of drink",
					},
				},
			},
		},
	}
}
