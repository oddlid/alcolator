package main

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const (
	keyLogger          = `logger`
	logTimeStampLayout = `2006-01-02T15:04:05.999-07:00`
	optLogLevel        = `log-level`
)

func logSetup(cCtx *cli.Context) error {
	zerolog.TimeFieldFormat = logTimeStampLayout
	if cCtx.IsSet(optLogLevel) {
		level, err := zerolog.ParseLevel(cCtx.String(optLogLevel))
		if err != nil {
			return err
		}
		zerolog.SetGlobalLevel(level)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	if cCtx.App != nil {
		if cCtx.App.Metadata == nil {
			cCtx.App.Metadata = make(map[string]any)
		}
		cCtx.App.Metadata[keyLogger] = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
	return nil
}

func appLogger(app *cli.App) zerolog.Logger {
	if app != nil {
		if app.Metadata != nil {
			if logger, ok := app.Metadata[keyLogger].(zerolog.Logger); ok {
				return logger
			}
		}
	}
	return log.Logger
}

func cCtxLogger(cCtx *cli.Context) zerolog.Logger {
	return appLogger(cCtx.App)
}

// getLevels returns a list of valid zerolog log levels, formatted as a comma separated string
func getLevels() string {
	levels := []string{
		zerolog.TraceLevel.String(),
		zerolog.DebugLevel.String(),
		zerolog.InfoLevel.String(),
		zerolog.WarnLevel.String(),
		zerolog.ErrorLevel.String(),
		zerolog.FatalLevel.String(),
		zerolog.PanicLevel.String(),
		zerolog.Disabled.String(),
	}
	return strings.Join(levels, ", ")
}
