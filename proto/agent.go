package proto

import (
	"crypto/rand"
	"encoding/binary"
)

type Agent struct {
	Name string
	Id   AgentID
}

func NewAgent(name string) Agent {
	r := Agent{Name: name}

	// generate unique agentId for this session
	binary.Read(rand.Reader, binary.NativeEndian, &r.Id)

	return r
}
