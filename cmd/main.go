package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var (
	VERSION string = "0.0.1"
)

type Drink struct {
	Name     string
	VolumeML int
	AlcPCT   float64
	Price    float64
}

//type Drinks []Drink

func (d *Drink) VolumeAsDL() float64 {
	return float64(d.VolumeML) / 100.0
}

func (d *Drink) VolumeAsL() float64 {
	return float64(d.VolumeML) / 1000.0
}

func (d *Drink) AlcML() int {
	return int((float64(d.VolumeML) / 100) * d.AlcPCT)
}

func (d *Drink) WaterML() int {
	return d.VolumeML - d.AlcML()
}

func (d *Drink) PricePerAlcML() float64 {
	return d.Price / float64(d.AlcML())
}

//func (d *Drink) EqOf(d2 Drink) Drink {
//	return Drink{
//		Name: d2.Name,
//	}
//}

func entryPoint(ctx *cli.Context) error {
	sdn := ctx.String("src-drink-name")
	sdv := ctx.Int("src-drink-volume")
	sdp := ctx.Float64("src-drink-pct")
	sdq := ctx.Int("src-drink-qty")
	price := ctx.Float64("src-drink-price")

	d := &Drink{
		Name:     sdn,
		VolumeML: sdv * sdq,
		AlcPCT:   sdp,
		Price:    price,
	}

	fmt.Printf("%d units at %dml each of %.1f%s %q, amounts to:\n", sdq, sdv, sdp, "%", sdn)
	fmt.Printf("Water        : %d ml\n", d.WaterML())
	fmt.Printf("Alcohol      : %d ml\n", d.AlcML())
	fmt.Printf("Price per ml : %.1f\n", d.PricePerAlcML())

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "AlcoLator"
	app.Usage = "Calculate drink equivalents"
	app.Copyright = "(c) 2018 Odd Eivind Ebbesen"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Odd E. Ebbesen",
			Email: "oddebb@gmail.com",
		},
	}
	app.Version = VERSION

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: "Log `level` (options: debug, info, warn, error, fatal, panic)",
		},
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "Run in debug mode",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:  "src-drink-name",
			Usage: "Name for source drink",
		},
		cli.IntFlag{
			Name:  "src-drink-volume",
			Usage: "Source drink volume in ML",
		},
		cli.Float64Flag{
			Name:  "src-drink-pct",
			Usage: "Source drink alcohol percentage",
		},
		cli.IntFlag{
			Name:  "src-drink-qty",
			Usage: "Source drink quantity",
			Value: 1,
		},
		cli.Float64Flag{
			Name:  "src-drink-price",
			Usage: "Price of source drink",
		},
		cli.StringFlag{
			Name:  "dst-drink-name",
			Usage: "Name for drink to compare against",
		},
		cli.Float64Flag{
			Name:  "dst-drink-pct",
			Usage: "Alcohol percentage of drink to compare against",
		},
	}

	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stderr)
		level, err := log.ParseLevel(c.String("log-level"))
		if err != nil {
			log.Fatal(err.Error())
		}
		log.SetLevel(level)
		if !c.IsSet("log-level") && !c.IsSet("l") && c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: false,
			FullTimestamp:    true,
		})
		return nil
	}

	app.Action = entryPoint
	app.Run(os.Args)
}
