package proto

import (
	"github.com/vsm0/gomog/proto/rpc"

	"testing"

	"github.com/golang/protobuf/proto"
)

func TestEncodeDecodePing(t *testing.T) {
	ping := &rpc.PingRequest{
		Counter: 0,
	}

	out, err := proto.Marshal(ping)
	if err != nil {
		t.Fatalf("%v", err)
	}

	for _, b := range out {
		t.Logf("%d", b)
	}

	p := &rpc.PingRequest{}

	if err := proto.Unmarshal(out, p); err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("%d", p.Counter)
}

func TestEncodeDecodePong(t *testing.T) {
	pong := &rpc.Pong{
		Counter: 0,
		Ok: true,
	}

	out, err := proto.Marshal(pong)
	if err != nil {
		t.Fatalf("%v", err)
	}

	for _, b := range out {
		t.Logf("%d", b)
	}

	p := &rpc.Pong{}

	if err := proto.Unmarshal(out, p); err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("%d, %t", p.Counter, p.Ok)
}
