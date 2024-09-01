package client

import (
	"github.com/vsm0/gomog/proto/rpc"

	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestPingService(t *testing.T) {
	port := 9001
	opts := []grpc.DialOption {
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(fmt.Sprintf(":%d", port), opts...)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer conn.Close()

	client := rpc.NewServerClient(conn)

	r := &rpc.PingRequest{
		Counter: 68,
	}

	t.Logf("ping: %d", r.Counter)

	pong, err := client.Ping(context.Background(), r)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("pong: %d (ok: %t)", pong.Counter, pong.Ok)
}
