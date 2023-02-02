package main

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/oddlid/alcolator/drink"
)

func printAlcoholAmount(w io.Writer, dc *drink.DrinkContainer) {
	if dc == nil {
		return
	}
	fmt.Fprintf(
		w,
		"%.2f %s of %.2f%% has %.2f %s of alcohol\n",
		dc.Volume,
		dc.VolumeUnit,
		dc.Drink.Percentage,
		dc.AlcoholAmount(),
		dc.VolumeUnit,
	)
}

func getAlcoholAmount(cCtx *cli.Context) error {
	unit := cCtx.String(optUnit)
	volume := cCtx.Float64(optVolume)
	percentage := cCtx.Float64(optPercentage)
	dc := drink.DrinkContainer{
		VolumeUnit: unit,
		Volume:     volume,
		Drink: drink.Drink{
			Percentage: percentage,
		},
	}
	printAlcoholAmount(os.Stdout, &dc)
	return nil
}
