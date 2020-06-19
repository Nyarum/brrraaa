package cmd

import "github.com/urfave/cli"

func CmdServer(c *cli.Context) error {
	srv, err := NewServer()
	if err != nil {
		return err
	}

	err = srv.Run()
	if err != nil {
		return err
	}

	return nil
}
