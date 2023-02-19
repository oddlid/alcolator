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
	unit := cCtx.String(optVolumeUnit)
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

func printComparison(w io.Writer, dc *drink.DrinkContainer, other *drink.DrinkContainer) {
	if dc == nil || other == nil {
		return
	}
	c := dc.CompareTo(other)
	// TODO: calculation is wrong, fix
	fmt.Fprintf(
		w,
		"It takes %.2f %s of %.2f%% to match the alcohol amount (%.2f %s) of %.2f %s of %.2f%%\n",
		c.Volume,
		c.VolumeUnit,
		c.Drink.Percentage,
		dc.AlcoholAmount(),
		dc.VolumeUnit,
		dc.Volume,
		dc.VolumeUnit,
		dc.Drink.Percentage,
	)
}

func compare(cCtx *cli.Context) error {
	unit := cCtx.String(optVolumeUnit)
	volume := cCtx.Float64(optVolume)
	percentage := cCtx.Float64(optPercentage)
	percentageOther := cCtx.Float64(optOtherPercentage)
	dc := drink.DrinkContainer{
		VolumeUnit: unit,
		Volume:     volume,
		Drink: drink.Drink{
			Percentage: percentage,
		},
	}
	other := drink.DrinkContainer{
		VolumeUnit: unit,
		Volume:     volume,
		Drink: drink.Drink{
			Percentage: percentageOther,
		},
	}
	printComparison(os.Stdout, &dc, &other)
	return nil
}
