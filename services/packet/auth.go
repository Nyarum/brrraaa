package packet

import (
	"bytes"
)

var (
	customAuthKey    = []byte{0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49}
	customAuthDWFlag = 12820
)

type OutcomingAuth struct {
	ErrorCode  uint16
	ErrorText  string // (if ErrorCode != 0)
	Key        []byte // customAuthKey
	Characters []Character
	Pincode    uint8
	Encryption uint32
	DwFlag     uint32 // customAuthDWFlag
}

func constructAuth(auth OutcomingAuth) (res []byte, err error) {
	buf := bytes.NewBuffer([]byte{})

	err = writeGeneric(buf, auth.ErrorCode)

	if auth.ErrorCode == 0 {
		err = writeString(buf, string(customAuthKey))

		err = writeGeneric(buf, uint8(len(auth.Characters)))
		for _, char := range auth.Characters {
			err = constructCharacter(buf, char)
		}

		err = writeGeneric(buf, auth.Pincode)
		err = writeGeneric(buf, auth.Encryption)
		err = writeGeneric(buf, uint32(customAuthDWFlag))
	} else {
		err = writeString(buf, auth.ErrorText)
	}

	if err != nil {
		res = buf.Bytes()
	}

	return
}

type IncomingAuth struct {
	Key           string
	Login         string
	Password      string
	MAC           string
	IsCheat       uint16
	ClientVersion uint16
}

func deconstructAuth(data []byte) (auth IncomingAuth, err error) {
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
