package controller

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

func ListDevices() error {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}

	fmt.Printf("Interfaces:\n")
	for k, ifs := range ifaces {
		fmt.Printf("[%v] 0x%x %s - %s\n", k, ifs.Flags, ifs.Description, ifs.Name)
		for _, addr := range ifs.Addresses {
			size, _ := addr.Netmask.Size()
			fmt.Printf("\t- %v / %v\n", addr.IP.String(), size)
		}
	}

	return nil
}
