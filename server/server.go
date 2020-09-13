package server

import (
	"fmt"

	"github.com/Nyarum/brrraaa/providers/config"
	"github.com/Nyarum/brrraaa/services/storage"
	"github.com/panjf2000/gnet"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*gnet.EventServer `wire:"-"`
	UserStorage       *storage.UserStorage
	Config            config.Config
}

func (s *Server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Info().
		Str("addr", srv.Addr.String()).
		Bool("isMulticore", srv.Multicore).
		Int("numEventLoop", srv.NumEventLoop).
		Msg("Server is started listening")
	return
}

func (s *Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Echo synchronously.
	out = frame
	return

	/*
	   // Echo asynchronously.
	   data := append([]byte{}, frame...)
	   go func() {
	       time.Sleep(time.Second)
	       c.AsyncWrite(data)
	   }()
	   return
	*/
}

func (s *Server) Run() (err error) {
	return gnet.Serve(s, fmt.Sprintf("tcp://%s", s.Config.Addr), gnet.WithMulticore(true))
}
