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
	Data      string
}

func (l *L2ChanLayer) LayerType() gopacket.LayerType {
	return L2ChanLayerType
}

func (l *L2ChanLayer) LayerContents() []byte {
	panic("LayerContents should never be used")
}

func (l *L2ChanLayer) LayerPayload() []byte {
	return nil
}

var L2ChanLayerType = gopacket.RegisterLayerType(1000, gopacket.LayerTypeMetadata{
	Name:    "L2ChanLayer",
	Decoder: gopacket.DecodeFunc(decodeL2ChanLayer),
})

func init() {
	layers.EthernetTypeMetadata[EthernetTypeL2Chan] = layers.EnumMetadata{
		Name:       "L2ChanLayer",
		LayerType:  L2ChanLayerType,
		DecodeWith: gopacket.DecodeFunc(decodeL2ChanLayer),
	}
}
