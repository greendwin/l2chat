package methods

import (
	"fmt"

	"github.com/google/gopacket/pcap"
	"github.com/greendwin/l2chat/proto"
)

func ListDevices() error {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}

	for k, ifs := range ifaces {
		if k > 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("Name:\t%s\nID:\t%s\n", ifs.Description, ifs.Name)

		hwaddr, ok, err := proto.FindDeviceHWAddr(ifs.Name)
		if err != nil {
			return err
		}
		if ok {
			fmt.Printf("MAC:\t%s\n", hwaddr.String())
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
