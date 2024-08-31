package server

import (
	"gomog/proto/rpc"

	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"testing"

	"google.golang.org/grpc"
)

type server struct {
	rpc.UnimplementedServerServer
}

func (s *server) Ping(c context.Context, r *rpc.PingRequest) (*rpc.Pong, error) {
	pong := &rpc.Pong{
		Counter: r.Counter + 1,
		Ok: r.Counter < 100,
	}

	return pong, nil
}

func TestServer(t *testing.T) {
	port := 9001
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("listening on :%d", port)

	s := grpc.NewServer()
	rpc.RegisterServerServer(s, &server{})
	go func() {
		s.Serve(l)
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-sigc
		t.Logf("terminating: %v", sig)
		cancel()
		s.GracefulStop()
	}()

	<-ctx.Done()
	t.Logf("server stop")
}
