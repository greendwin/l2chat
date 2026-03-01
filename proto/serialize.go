package proto

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/google/gopacket"
)

var errorTooShort = errors.New("packet data is too short")

const headerSize = 4 + 1 + 1 // AgentID + Operation + DataLen

func decodeL2ChanLayer(data []byte, pb gopacket.PacketBuilder) error {
	if len(data) < headerSize {
		return errorTooShort
	}

	id := AgentID(binary.BigEndian.Uint32(data[:4]))
	op := L2Operation(data[4])
	dataLen := data[5]

	pb.AddLayer(&L2ChanLayer{
		AgentID:   id,
		Operation: op,
		DataLen:   dataLen,
		Data:      string(data[headerSize : headerSize+dataLen]),
		Payload:   data[headerSize+dataLen:],
	})

	return pb.NextDecoder(gopacket.DecodePayload)
}

func (l *L2ChanLayer) SerializeTo(b gopacket.SerializeBuffer, opts gopacket.SerializeOptions) error {
	header, err := b.PrependBytes(headerSize + len(l.Data))
	if err != nil {
		return fmt.Errorf("L2ChanLayer.SerializeTo: %w", err)
	}

	binary.BigEndian.PutUint32(header[:4], uint32(l.AgentID))
	header[4] = uint8(l.Operation)
	header[5] = uint8(len(l.Data))

	copy(header[headerSize:], []byte(l.Data))

	return nil
}
