package server

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/greendwin/l2chat/proto"
	"golang.org/x/sync/errgroup"
)

const helloPeriod = 5 * time.Second

type agentInfo struct {
	Name          string
	Id            proto.AgentID
	IsOnline      bool
	LastTimestamp time.Time
}

type server struct {
	agent  proto.Agent
	others map[proto.AgentID]*agentInfo
}

func NewServer(name string) *server {
	return &server{
		agent: proto.NewAgent(name),
	}
}

func (s *server) Run(devices []string) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	gr, ctx := errgroup.WithContext(ctx)

	for _, deviceName := range devices {
		gr.Go(func() error {
			return s.serveDevice(ctx, &s.agent, deviceName)
		})
	}

	return gr.Wait()
}

func (s *server) serveDevice(ctx context.Context, agent *proto.Agent, deviceName string) error {
	conn, err := agent.Connect(deviceName)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.SendHello()
	if err != nil {
		return err
	}

	// try send goodbye on exit
	defer conn.SendBye()

	ch, err := conn.Listen()
	if err != nil {
		return err
	}

	helloTicker := time.NewTicker(helloPeriod)
	defer helloTicker.Stop()

	for {
		select {
		case msg := <-ch:
			err = s.processMessage(msg)
			if err != nil {
				return err
			}

		case <-helloTicker.C:
			err = conn.SendHello()
			if err != nil {
				return err
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) processMessage(msg *proto.L2ChanLayer) error {
	log.Printf("RECV 0x%x %s %q", msg.AgentID, msg.Operation.String(), msg.Data)

	if msg.AgentID == s.agent.Id {
		// skip our own messages if any
		return nil
	}

	switch msg.Operation {
	case proto.OpHello:
		p := s.others[msg.AgentID]
		if p == nil {
			p = &agentInfo{
				Name: msg.Data,
				Id:   msg.AgentID,
			}
			s.others[msg.AgentID] = p
		}

		if !p.IsOnline {
			log.Printf("user %q connected", p.Name)
			p.IsOnline = true
		}

		p.LastTimestamp = time.Now()

	case proto.OpBye:
		if p := s.others[msg.AgentID]; p != nil {
			p.IsOnline = false
			p.LastTimestamp = time.Now()
			log.Printf("user %q disconnected", p.Name)
		}

	case proto.OpEcho:
	case proto.OpEchoReply:
	case proto.OpMessage:
	}

	panic("not implemented")
}
