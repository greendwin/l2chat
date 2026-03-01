package methods

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/greendwin/l2chat/proto"
	"golang.org/x/sync/errgroup"
)

var (
	helloPeriod = 5 * time.Second
)

func RunServer(name string, devices []string) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	gr, ctx := errgroup.WithContext(ctx)

	agent := proto.NewAgent(name)

	for _, deviceName := range devices {
		gr.Go(func() error {
			return serve(ctx, &agent, deviceName)
		})
	}

	return gr.Wait()
}

func serve(ctx context.Context, agent *proto.Agent, deviceName string) error {
	conn, err := agent.Connect(deviceName)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.SendHello()
	if err != nil {
		return err
	}

	defer func() {
		// try send goodbye on exit
		_ = conn.SendBye()
	}()

	t := time.NewTicker(helloPeriod)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			err = conn.SendHello()
			if err != nil {
				return err
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
