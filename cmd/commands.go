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
		"( %.2f %s / 100.0 ) * %.2f %% = %.2f %s\n",
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

func printComparison(w io.Writer, firstDrink, otherDrink *drink.DrinkContainer) {
	if firstDrink == nil || otherDrink == nil {
		return
	}
	c := firstDrink.CompareTo(otherDrink)
	fmt.Fprintf(
		w,
		"( %.2f %% / %.2f %% ) * %.2f %s = %.2f %s [ = %.2f %s alcohol ]\n",
		firstDrink.Drink.Percentage,
		otherDrink.Drink.Percentage,
		firstDrink.Volume,
		firstDrink.VolumeUnit,
		c.Volume,
		c.VolumeUnit,
		c.AlcoholAmount(),
		c.VolumeUnit,
	)
	if firstDrink.Price == 0 || otherDrink.Price == 0 {
		return
	}
}

func compare(cCtx *cli.Context) error {
	volumeUnit := cCtx.String(optVolumeUnit)
	priceUnit := cCtx.String(optPriceUnit)
	firstDrink := drink.DrinkContainer{
		VolumeUnit: volumeUnit,
		PriceUnit:  priceUnit,
		Volume:     cCtx.Float64(optVolume),
		Price:      cCtx.Float64(optPrice),
		Drink: drink.Drink{
			Percentage: cCtx.Float64(optPercentage),
		},
	}
	otherDrink := drink.DrinkContainer{
		VolumeUnit: volumeUnit,
		PriceUnit:  priceUnit,
		Price:      cCtx.Float64(optOtherPrice),
		Drink: drink.Drink{
			Percentage: cCtx.Float64(optOtherPercentage),
		},
	}
	printComparison(os.Stdout, &firstDrink, &otherDrink)
	return nil
}
