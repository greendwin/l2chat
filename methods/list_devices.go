package methods

import (
	"fmt"

	"github.com/google/gopacket/pcap"
	"github.com/greendwin/l2chat/proto"
)

func ListDevices(showAllDevices bool) error {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}

	first := true

	for _, ifs := range ifaces {
		hwaddr, ok, err := proto.FindDeviceHWAddr(ifs.Name)
		if err != nil {
			return err
		}
		if !ok && !showAllDevices {
			// skip devices without MAC
			continue
		}

		if !first {
			fmt.Printf("\n")
		}
		first = false

		fmt.Printf("Name:\t%s\n", ifs.Description)
		fmt.Printf("ID:\t%s\n", ifs.Name)

		if len(hwaddr) > 0 {
			fmt.Printf("MAC:\t%s\n", hwaddr.String())
		} else {
			fmt.Printf("MAC:\t<MISSING>\n")
		}

		if len(ifs.Addresses) > 0 {
			fmt.Printf("Addresses:\n")
		}
		for _, addr := range ifs.Addresses {
			size, _ := addr.Netmask.Size()
			fmt.Printf("\t- %v / %v\n", addr.IP.String(), size)
		}
	}

	return nil
}
