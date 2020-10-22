package packet

import (
	"bytes"
	"encoding/binary"
)

type ItemAttribute struct {
	ID    uint16
	Value uint16
}

type Item struct {
	ID         uint16
	Num        uint16
	Endures    [2]uint16
	Energies   [2]uint16
	ForgeLv    uint8
	PassValue  uint8
	DbParams   [2]uint32
	InstAttrs  [5]ItemAttribute
	ItemAttrs  [58]uint16
	InitFlag   uint8
	PassValue2 uint8
	Valid      bool
	Change     bool
}

func constructItem(buf *bytes.Buffer, item Item) (err error) {
	err = writeGeneric(buf, item.ID)
	err = writeGeneric(buf, item.Num)

	for _, endure := range item.Endures {
		err = writeGeneric(buf, endure)
	}

	for _, energy := range item.Energies {
		err = writeGeneric(buf, energy)
	}

	err = writeGeneric(buf, item.ForgeLv)
	err = writeGeneric(buf, item.PassValue)

	for _, param := range item.DbParams {
		err = writeGeneric(buf, param)
	}

	for _, attr := range item.InstAttrs {
		err = writeGeneric(buf, attr.ID)
		err = writeGeneric(buf, attr.Value)
	}

	for _, attr := range item.ItemAttrs {
		err = writeGeneric(buf, attr)
	}

	err = writeGeneric(buf, item.InitFlag)
	err = writeGeneric(buf, item.PassValue2)
	err = writeGeneric(buf, item.Valid)
	err = writeGeneric(buf, item.Change)

	return
}

type CharacterLook struct {
	Ver    uint16
	TypeID uint16
	Items  [10]Item
	Hair   uint16
}

func constructCharacterLook(buf *bytes.Buffer, charLook CharacterLook) (err error) {
	err = writeGeneric(buf, uint16(binary.Size(charLook)))
	err = writeGeneric(buf, charLook.Ver)
	err = writeGeneric(buf, charLook.TypeID)

	for _, item := range charLook.Items {
		err = constructItem(buf, item)
	}

	err = writeGeneric(buf, charLook.Hair)

	return
}

type Character struct {
	Flag  uint8
	Name  string
	Job   string
	Level uint16
	Look  CharacterLook
}

func constructCharacter(buf *bytes.Buffer, char Character) (err error) {
	err = writeGeneric(buf, char.Flag)
	err = writeString(buf, char.Name)
	err = writeString(buf, char.Job)
	err = writeGeneric(buf, char.Level)
	err = constructCharacterLook(buf, char.Look)

	return
}
