package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/oddlid/alcolator"
	"github.com/oddlid/alcolator/srv/assets"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"github.com/urfave/cli"
)

const (
	E_OK int = iota
	E_TMPL_LOAD
)

var (
	VERSION    string = "undef"
	COMMIT_ID  string = "undef"
	BUILD_DATE string = "undef"
	ctmpl      *template.Template
)

type FormData struct {
	HasData               bool
	DrinkName             string
	DrinkVolML            int
	DrinkPCT              float64
	DrinkPrice            float64
	ResultWaterML         int
	ResultAlcML           int
	ResultPricePerAlcML   float64
	ResultPricePerDrinkML float64
}


func CalcHandler(w http.ResponseWriter, r *http.Request) {
	hasData := r.FormValue("dsub") != ""

	if !hasData {
		ctmpl.Execute(w, nil)
		return
	}

	fd := FormData{}
	fd.HasData = hasData
	if !fd.HasData {
		log.Warnf("Should have form data, but it seems to be missing... From: %q", r.RemoteAddr)
	}

	fd.DrinkName = r.FormValue("dname")
	var err error
	fd.DrinkVolML, err = strconv.Atoi(r.FormValue("dvolml"))
	if err != nil {
		log.Errorf("Failed to convert dvolml. Input: %q from %q", r.FormValue("dvolml"), r.RemoteAddr)
	}
	dpct := strings.Replace(r.FormValue("dpct"), ",", ".", -1) // replace comma with dot in input
	fd.DrinkPCT, err = strconv.ParseFloat(dpct, 64)
	if err != nil {
		log.Errorf("Failed to convert dpct. Input: %q from %q", r.FormValue("dpct"), r.RemoteAddr)
	}
	fd.DrinkPrice, err = strconv.ParseFloat(r.FormValue("dprice"), 64)
	if err != nil {
		log.Errorf("Failed to convert dprice. Input: %q from %q", r.FormValue("dprice"), r.RemoteAddr)
	}

	d := &alcolator.Drink{
		Name:     fd.DrinkName,
		VolumeML: fd.DrinkVolML,
		AlcPCT:   fd.DrinkPCT,
		Price:    fd.DrinkPrice,
	}

	fd.ResultWaterML = d.WaterML()
	fd.ResultAlcML = d.AlcML()
	fd.ResultPricePerAlcML = d.PricePerAlcML()
	fd.ResultPricePerDrinkML = d.PricePerDrinkML()

	ctmpl.Execute(w, &fd)
}

// initialize template via vfsgen
func initTmpl() error {
	log.Debug("Loading template...")
	tmpl, err := vfstemplate.ParseFiles(assets.Assets, nil, "/templates/apkform.html")
	if err != nil {
		return err
	}

	ctmpl = tmpl

	return nil
}

func serve(ctx *cli.Context) error {
	addr := ctx.String("listen")

	// initialize template
	err := initTmpl()
	if err != nil {
		log.Error(err)
		return cli.NewExitError(err.Error(), E_TMPL_LOAD)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", CalcHandler)
	r.PathPrefix("/").Handler(http.FileServer(assets.Assets)) // needed for css files
	log.Infof("Server listening on %s", addr)
	return http.ListenAndServe(addr, r)
}

func main() {
	app := cli.NewApp()
	app.Name = "AlcoLatorSrv"
	app.Usage = "Calculate drink values"
	app.Copyright = "(c) 2018 Odd Eivind Ebbesen"
	app.Version = fmt.Sprintf("%s_%s (Compiled: %s)", VERSION, COMMIT_ID, BUILD_DATE)
	app.Compiled, _ = time.Parse(time.RFC3339, BUILD_DATE)

	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Odd E. Ebbesen",
			Email: "oddebb@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-level",
			Value: "info",
			Usage: "Log `level` (options: debug, info, warn, error, fatal, panic)",
		},
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "Run in debug mode",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:  "listen, l",
			Usage: "`ADDR` to listen on",
			Value: ":9600",
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

	app.Action = serve
	app.Run(os.Args)
}
