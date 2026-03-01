package main

import (
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Interfaces:")
	for k, ifs := range ifaces {
		log.Printf("[%v] %s (0x%x) - %s", k, ifs.Description, ifs.Flags, ifs.Name)
		for _, addr := range ifs.Addresses {
			size, _ := addr.Netmask.Size()
			log.Printf("\t- %v / %v", addr.IP.String(), size)
		}
	}
}
