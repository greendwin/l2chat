package proto

import (
	"context"
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func (c *Connection) Listen(ctx context.Context) (<-chan *L2ChanLayer, error) {
	err := c.handle.SetBPFFilter(fmt.Sprintf("ether proto %v", EthernetTypeL2Chan))
	if err != nil {
		return nil, fmt.Errorf("failed to set BPF filter: %w", err)
	}

	output := make(chan *L2ChanLayer, 100)

	go func() {
		defer close(output)

		var eth layers.Ethernet
		var l2chan L2ChanLayer
		parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &l2chan)
		decoded := make([]gopacket.LayerType, 0, 4)

		for {
			// check whether context was cancelled
			select {
			case <-ctx.Done():
				return
			default:
			}

			data, _, err := c.handle.ReadPacketData()
			if err != nil {
				// wait some time and repeat
				time.Sleep(10 * time.Millisecond)
				continue
			}

			err = parser.DecodeLayers(data, &decoded)
			if err != nil {
				// skip broken packet
				continue
			}

			for _, lt := range decoded {
				if lt == LayerTypeL2Chan {
					r := l2chan
					output <- &r
				}
			}
		}
	}()

	return output, nil
}
