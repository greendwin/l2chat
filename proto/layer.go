package proto

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const EthernetTypeL2Chan = 7890

type AgentID uint32

type L2ChanLayer struct {
	AgentID   AgentID
	Operation L2Operation
	DataLen   uint8
	Data      string
	Payload   []byte
}

func (l *L2ChanLayer) LayerType() gopacket.LayerType {
	return LayerTypeL2Chan
}

func (l *L2ChanLayer) LayerContents() []byte {
	panic("LayerContents should never be used")
}

func (l *L2ChanLayer) LayerPayload() []byte {
	return l.Payload
}

var LayerTypeL2Chan = gopacket.RegisterLayerType(1000, gopacket.LayerTypeMetadata{
	Name:    "L2ChanLayer",
	Decoder: gopacket.DecodeFunc(decodeL2ChanLayer),
})

func init() {
	layers.EthernetTypeMetadata[EthernetTypeL2Chan] = layers.EnumMetadata{
		Name:       "L2ChanLayer",
		LayerType:  LayerTypeL2Chan,
		DecodeWith: gopacket.DecodeFunc(decodeL2ChanLayer),
	}
}

func decodeL2ChanLayer(data []byte, pb gopacket.PacketBuilder) error {
	l := L2ChanLayer{}
	err := l.DecodeFromBytes(data, pb)
	if err != nil {
		return err
	}

	pb.AddLayer(&l)

	return pb.NextDecoder(gopacket.DecodePayload)
}
