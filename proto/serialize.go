package proto

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/google/gopacket"
)

var errorTooShort = errors.New("packet data is too short")

const headerSize = 4 + 1 + 1 // AgentID + Operation + DataLen

func (l *L2ChanLayer) CanDecode() gopacket.LayerClass {
	return LayerTypeL2Chan
}

func (l *L2ChanLayer) NextLayerType() gopacket.LayerType {
	return gopacket.LayerTypePayload
}

func (l *L2ChanLayer) DecodeFromBytes(data []byte, df gopacket.DecodeFeedback) error {
	if len(data) < headerSize {
		df.SetTruncated()
		return errorTooShort
	}

	l.AgentID = AgentID(binary.BigEndian.Uint32(data[:4]))
	l.Operation = L2Operation(data[4])
	l.DataLen = data[5]
	l.Data = string(data[headerSize : headerSize+l.DataLen])
	l.Payload = data[headerSize+l.DataLen:]

	return nil
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

var _ gopacket.DecodingLayer = (*L2ChanLayer)(nil)
