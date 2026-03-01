package proto

import (
	"crypto/rand"
	"encoding/binary"
	"log"
)

type Agent struct {
	Name string
	Id   AgentID
}

func NewAgent(name string) Agent {
	r := Agent{Name: name}

	// generate unique agentId for this session
	binary.Read(rand.Reader, binary.NativeEndian, &r.Id)

	log.Printf("AgentName = %q", r.Name)
	log.Printf("AgentID = 0x%08x", r.Id)

	return r
}
