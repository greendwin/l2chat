package proto

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const connectionTimeout = 5 * time.Second

type Connection struct {
	agent  *Agent
	hwaddr net.HardwareAddr
	handle *pcap.Handle
	buf    gopacket.SerializeBuffer
}

func (a *Agent) Connect(deviceName string) (*Connection, error) {
	hwaddr, ok, err := FindDeviceHWAddr(deviceName)
	if err != nil {
		return nil, fmt.Errorf("FindDeviceHWAddr: %w", err)
	}

	if !ok {
		return nil, fmt.Errorf("unable to resolve MAC address for device %q", deviceName)
	}

	handle, err := pcap.OpenLive(deviceName, 65536, false, connectionTimeout)
	if err != nil {
		return nil, fmt.Errorf("pcap.OpenLive: %w", err)
	}

	r := &Connection{
		agent:  a,
		hwaddr: hwaddr,
		handle: handle,
		buf:    gopacket.NewSerializeBuffer(),
	}

	return r, nil
}

func (c *Connection) Close() {
	c.handle.Close()
}

func FindDeviceHWAddr(deviceName string) (net.HardwareAddr, bool, error) {
	netIfaces, err := net.Interfaces()
	if err != nil {
		return nil, false, err
	}

	pcapDevices, err := pcap.FindAllDevs()
	if err != nil {
		return nil, false, err
	}

	for _, netIfs := range netIfaces {
		if len(netIfs.HardwareAddr) == 0 {
			// skip broken hardware addresses like in `127.0.0.1`
			continue
		}

		netAddrs, err := netIfs.Addrs()
		if err != nil {
			return nil, false, err
		}

		for _, netAddr := range netAddrs {
			netAddrValue := netAddr.String()

			for _, pcapIface := range pcapDevices {
				if pcapIface.Name != deviceName && pcapIface.Description != deviceName {
					continue
				}

				for _, pcapAddr := range pcapIface.Addresses {
					pcapAddrValue := pcapAddr.IP.String()
					if strings.Contains(netAddrValue, pcapAddrValue) {
						return netIfs.HardwareAddr, true, nil
					}
				}
			}
		}
	}

	return nil, false, nil
}
