package packet

type (
	SubProtocol uint8
	ServiceType uint8
	StreamType  uint8
	NBS         uint8
)

const (
	SubProtocolAudio   SubProtocol = 0x00
	SubProtocolSerial  SubProtocol = 0x20
	SubProtocolText    SubProtocol = 0x40
	SubProtocolService SubProtocol = 0x60
	SubProtocolMask    SubProtocol = 0xE0
)

const (
	ServicePing             ServiceType = 0x00
	ServicePong             ServiceType = 0x00
	ServiceChatUTF8         ServiceType = 0x01
	ServiceRTPacketRegister ServiceType = 0x20
	ServiceRTPacket         ServiceType = 0x21
	ServiceRequestReply     ServiceType = 0x02
	ServiceFnctReply        ServiceType = 0x80
)

const (
	StreamAscii StreamType = 0x00
	StreamUTF8  StreamType = 0x10
	StreamWCHAR StreamType = 0x20
)

const (
	NBSZero NBS = 0
	NBSOne  NBS = 1
)

const (
	HeaderSize          = 28 // VBAN service header size.
	StreamnameMaxLength = 16
)
