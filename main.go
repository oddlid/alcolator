package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/oddlid/alcolator/srv"
	"github.com/oddlid/alcolator/srv/assets"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"github.com/urfave/cli/v2"
)

const (
	exitOK int = iota
	exitTemplateLoad
)

var (
	version   string
	commitID  string
	buildDate string
)

func serve(ctx *cli.Context) error {
	addr := ctx.String("listen")

	// initialize template
	tmpl, err := vfstemplate.ParseFiles(assets.Assets, nil, "/templates/apkform.html")
	if err != nil {
		log.Error().
			Err(err).
			Send()
		return cli.NewExitError(err.Error(), exitTemplateLoad)
	}

	alcSrv := srv.NewServer(tmpl)

	r := mux.NewRouter()
	r.HandleFunc("/", alcSrv.CalcHandler)
	r.PathPrefix("/").Handler(http.FileServer(assets.Assets)) // needed for css files
	log.Info().
		Str("listen_address", addr).
		Msg("Server listening")
	return http.ListenAndServe(addr, r)
}

func main() {
	app := cli.NewApp()
	app.Name = "AlcoLatorSrv"
	app.Usage = "Calculate drink values"
	app.Copyright = "(c) 2018 Odd Eivind Ebbesen"
	app.Version = fmt.Sprintf("%s_%s (Compiled: %s)", version, commitID, buildDate)
	app.Compiled, _ = time.Parse(time.RFC3339, buildDate)

	app.Authors = []*cli.Author{
		{
			Name:  "Odd E. Ebbesen",
			Email: "oddebb@gmail.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"srv"},
			Usage:   "Start server",
			Action:  serve,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "listen",
					Aliases: []string{"l"},
					Usage:   "`ADDR` to listen on",
					Value:   ":9600",
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		// Implement this later
		//&cli.StringFlag{
		//	Name:  "log-level",
		//	Value: "info",
		//	Usage: "Log `level` (options: debug, info, warn, error, fatal, panic)",
		//},
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "Run in debug mode",
			EnvVars: []string{"DEBUG"},
		},
	}

	app.Before = func(c *cli.Context) error {
		//log.SetOutput(os.Stderr)
		//level, err := log.ParseLevel(c.String("log-level"))
		//if err != nil {
		//	log.Fatal(err.Error())
		//}
		//log.SetLevel(level)
		//if !c.IsSet("log-level") && !c.IsSet("l") && c.Bool("debug") {
		//	log.SetLevel(log.DebugLevel)
		//}
		//log.SetFormatter(&log.TextFormatter{
		//	DisableTimestamp: false,
		//	FullTimestamp:    true,
		//})
		if c.Bool("debug") {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
		zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999-07:00"
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error().
			Err(err).
			Send()
	}

	os.Exit(exitOK)
}
