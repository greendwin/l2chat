package proto

import (
	"fmt"

	"github.com/google/gopacket"
)

func (c *Connection) Listen() (<-chan *L2ChanLayer, error) {
	err := c.handle.SetBPFFilter(fmt.Sprintf("ether proto %v", EthernetTypeL2Chan))
	if err != nil {
		return nil, fmt.Errorf("failed to set BPF filter: %w", err)
	}

	src := gopacket.NewPacketSource(c.handle, c.handle.LinkType())
	src.DecodeOptions.NoCopy = true
	src.DecodeOptions.Lazy = true

	output := make(chan *L2ChanLayer, 100)
	input := src.Packets()

	go func() {
		defer close(output)

		for packet := range input {
			processPacket(packet, output)
		}
	}()

	return output, nil
}

func processPacket(p gopacket.Packet, output chan *L2ChanLayer) {
	l := p.Layer(L2ChanLayerType)
	if l == nil {
		return
	}

	layer, ok := l.(*L2ChanLayer)
	if !ok {
		panic("failed to cast `L2ChanLayer`")
	}

	output <- layer
}
