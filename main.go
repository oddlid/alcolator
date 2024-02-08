package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"github.com/urfave/cli/v2"

	"github.com/oddlid/alcolator/srv"
	"github.com/oddlid/alcolator/srv/assets"
)

const (
	eOK int = iota
	eTmplLoad
)

var (
	Version   string
	CommitID  string
	BuildDate string
)

func serve(ctx *cli.Context) error {
	addr := ctx.String("listen")

	// initialize template
	tmpl, err := vfstemplate.ParseFiles(assets.Assets, nil, "/templates/apkform.html")
	if err != nil {
		log.Error().
			Err(err).
			Send()
		return cli.Exit(err.Error(), eTmplLoad)
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

func newApp() *cli.App {
	compiled, _ := time.Parse(time.RFC3339, BuildDate)
	return &cli.App{
		Name:      "AlcoLatorSrv",
		Usage:     "Calculate drink values",
		Copyright: "(c) 2018 Odd Eivind Ebbesen",
		Version:   fmt.Sprintf("%s_%s (Compiled: %s)", Version, CommitID, BuildDate),
		Compiled:  compiled,
		Authors: []*cli.Author{
			{
				Name:  "Odd E. Ebbesen",
				Email: "oddebb@gmail.com",
			},
		},
		Commands: []*cli.Command{
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
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Run in debug mode",
				EnvVars: []string{"DEBUG"},
			},
		},
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}
			zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999-07:00"
			return nil
		},
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	if err := newApp().RunContext(ctx, os.Args); err != nil {
		log.Error().
			Err(err).
			Send()
	}
}
