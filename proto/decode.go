package proto

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/google/gopacket"
)

var errorTooShort = errors.New("packet data is too short")

const headerSize = 4 + 1 // AgentID + Operation

func decodeL2ChanLayer(data []byte, pb gopacket.PacketBuilder) error {
	if len(data) < headerSize {
		return errorTooShort
	}

	l := L2ChanLayer{
		AgentID:   AgentID(binary.BigEndian.Uint32(data[:4])),
		Operation: L2Operation(data[4]),
		Data:      string(data[headerSize:]),
	}

	pb.AddLayer(&l)

	return pb.NextDecoder(gopacket.DecodePayload)
}

func (l *L2ChanLayer) SerializeTo(b gopacket.SerializeBuffer, opts gopacket.SerializeOptions) error {
	contents, err := b.PrependBytes(headerSize + len(l.Data))
	if err != nil {
		return fmt.Errorf("L2ChanLayer.SerializeTo: %w", err)
	}

	binary.BigEndian.PutUint32(contents[:4], uint32(l.AgentID))
	contents[4] = uint8(l.Operation)
	copy(contents[headerSize:], []byte(l.Data))

	return nil
}
