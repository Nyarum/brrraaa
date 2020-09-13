package packet

import (
	"bytes"
)

type OutcomingAuth struct {
	ErrorCode uint16
	// ErrorText string (if ErrorCode != 0)
	LenKey uint16
	Key    []byte
	//Characters []common.CharacterSub
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32
}

func ConstructAuth(auth OutcomingAuth) ([]byte, error) {
	return nil, nil
}

type IncomingAuth struct {
	Key           string
	Login         string
	Password      string
	MAC           string
	IsCheat       uint16
	ClientVersion uint16
}

func DeconstructAuth(data []byte) (auth IncomingAuth, err error) {
	buf := bytes.NewBuffer(data)

	auth.Key, err = readString(buf)
	auth.Login, err = readString(buf)
	auth.Password, err = readString(buf)
	auth.MAC, err = readString(buf)
	auth.Password, err = readString(buf)
	auth.IsCheat, err = readUint16(buf)
	auth.ClientVersion, err = readUint16(buf)

	return
}
