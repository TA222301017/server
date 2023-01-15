package template

const (
	KeyExchange byte = iota
	AddAccessRule
	EditAccessRule
	DeleteAccessRule
	LogAccessEvent
	LogRSSIEvent
	LogHealthcheckEvent
)

type BasePacket struct {
	OpCode    byte
	Data      []byte
	Signature []byte
}

func (p BasePacket) Bytes() []byte {
	temp := []byte{}
	temp = append(temp, byte(p.OpCode))
	temp = append(temp, p.Data...)
	temp = append(temp, p.Signature...)

	return temp
}
