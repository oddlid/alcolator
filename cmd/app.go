package main

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

const (
	envLogLevel        = `LOG_LEVEL`
	cmdAlcAmount       = `alcoholamount`
	cmdCompare         = `compare`
	optVolumeUnit      = `volume-unit`
	optPriceUnit       = `price-unit`
	optPercentage      = `percentage`
	optOtherPercentage = `other-percentage`
	optVolume          = `volume`
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
			&cli.StringFlag{
				Name:    optVolumeUnit,
				Usage:   "Volume unit",
				Value:   "ml",
				Aliases: []string{"V"},
			},
			&cli.StringFlag{
				Name:    optPriceUnit,
				Usage:   "Price unit",
				Value:   "SEK",
				Aliases: []string{"P"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    cmdAlcAmount,
				Aliases: []string{"a", "amount"},
				Usage:   "Get the amount of alcohol in a drink",
				Action:  getAlcoholAmount,
				Flags: []cli.Flag{
					&cli.Float64Flag{
						Name:    optPercentage,
						Usage:   "Alcohol percentage of drink",
						Aliases: []string{"p"},
					},
					&cli.Float64Flag{
						Name:    optVolume,
						Usage:   "Volume of drink",
						Aliases: []string{"v"},
					},
				},
			},
			{
				Name:    cmdCompare,
				Aliases: []string{"c"},
				Usage:   "Calculate how much it takes of a drink of given strength to match another of different strength",
				Action:  compare,
				Flags: []cli.Flag{
					&cli.Float64Flag{
						Name:    optPercentage,
						Usage:   "Alcohol percentage of drink",
						Aliases: []string{"p"},
					},
					&cli.Float64Flag{
						Name:    optVolume,
						Usage:   "Volume of drink",
						Aliases: []string{"v"},
					},
					&cli.Float64Flag{
						Name:    optOtherPercentage,
						Usage:   "Alcohol percentage of other drink",
						Aliases: []string{"o"},
					},
				},
			},
		},
	}
}
