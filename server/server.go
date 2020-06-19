package server

import (
	"net"

	"github.com/Nyarum/brrraaa/core"
	"github.com/rs/zerolog/log"
)

type Server struct {
	DB     *core.DB
	Client *Client
}

func (s Server) Run() (err error) {
	log.Info().
		Msg("Server is starting")

	l, err := net.Listen("tcp4", ":9999")
	if err != nil {
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		_ = c
	}
}
