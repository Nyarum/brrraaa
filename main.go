package main

import (
	"os"

	"github.com/Nyarum/brrraaa/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "starting the emulator server",
				Action:  cmd.CmdServer,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to running a new server")
	}
}
