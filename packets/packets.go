package packets

import (
	"main/gmessages"
	"net"
)

type Packet struct {
	ID      uint8
	Content []byte
}

type PacketWithConnDTO struct {
	ID      uint8
	Content []byte
	Conn    net.Conn
}

type PacketDTO struct {
	ID           uint8
	Message      interface{}
	Connection   net.Conn
	PacketString string
}

type ReturnPacketDTO struct {
	ID      uint8
	Content []byte
	Conn    net.Conn
}

type ReturnPacket struct {
	ID      uint8
	Content []byte
}

type TokenPacketContent struct {
	Token     Token
	Signature []byte
}

var PACKETS_INT_TO_STRING = map[uint8]string{
	101: "MSG_LOGIN",
	102: "MSG_REGISTER",
}

var PACKETS_STRING_TO_INT = map[string]uint8{
	"MSG_LOGIN":    101,
	"MSG_REGISTER": 102,
}

var PACKET_TO_GAME_MESSAGE = map[string]interface{}{
	"MSG_LOGIN":    gmessages.LoginMessage{},
	"MSG_REGISTER": gmessages.RegisterMessage{},
}
