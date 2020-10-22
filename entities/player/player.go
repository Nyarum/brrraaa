package player

import (
	"github.com/Nyarum/brrraaa/services/packet"
	"github.com/panjf2000/gnet"
)

type Player struct {
	conn gnet.Conn
}

func NewPlayer(conn gnet.Conn) *Player {
	return &Player{
		conn: conn,
	}
}

func (p Player) ManageEvent(obj interface{}) {
	var (
		resData []byte
		err     error
	)

	switch v := obj.(type) {
	case packet.IncomingAuth:
		_ = v.ClientVersion

		resData, err = packet.ConstructPacket(
			packet.OutcomingAuth{
				ErrorCode: 123,
				ErrorText: "Bla bla bla",
			},
		)
		if err != nil {
			// logging error if exist
		}
	}

	if err != nil {
		p.conn.AsyncWrite(resData)
	}

	return
}
