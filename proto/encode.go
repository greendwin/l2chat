package proto

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const (
	OpHello = iota // payload: username
	OpBye
	OpEcho      // payload: data that should be sent back
	OpEchoReply // payload: data from Echo
	OpMessage   // payload: text
)

var broadcastHWAddr = net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

func (c *Connection) SendHello() error {
	err := c.sendPacket(OpHello, c.agent.Name)
	if err != nil {
		return err
	}

	log.Printf("write 'hello' %v bytes", len(c.buf.Bytes()))

	return nil
}

func (c *Connection) SendBye() error {
	err := c.sendPacket(OpBye, "")
	if err != nil {
		return err
	}

	log.Printf("write 'bye' %v bytes", len(c.buf.Bytes()))

	return nil
}

func (c *Connection) sendPacket(op uint8, data string) error {
	ether := layers.Ethernet{
		SrcMAC:       c.hwaddr,
		DstMAC:       broadcastHWAddr,
		EthernetType: EthernetTypeL2Chan,
	}

	l2chan := L2ChanLayer{
		AgentID:   c.agent.Id,
		Operation: op,
		Data:      data,
	}

	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	err := gopacket.SerializeLayers(c.buf, opts, &ether, &l2chan)
	if err != nil {
		return fmt.Errorf("serialization failed: %w", err)
	}

	err = c.handle.WritePacketData(c.buf.Bytes())
	if err != nil {
		return fmt.Errorf("WritePacketData: %w", err)
	}

	return nil
}
