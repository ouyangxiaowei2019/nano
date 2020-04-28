package io

import (
	"testing"
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/connector"
	"github.com/lonng/nano/serialize/protobuf"
	"github.com/lonng/nano/tests/io/testdata"
)

const (
	addr = "127.0.0.1:13250" // local address
	conc = 1000              // concurrent client count
)

var (
	components component.Components
)

func client(t *testing.T) {
	c := connector.NewConnector()

	chReady := make(chan struct{}, 1)
	c.OnConnected(func() {
		// t.Log("client on connected")
		chReady <- struct{}{}
	})

	if err := c.Start(addr); err != nil {
		panic(err)
	}

	c.On("pong", func(data interface{}) {
		// t.Log("pong received")
	})

	<-chReady
	for {
		c.Notify("TestHandler.Ping", &testdata.Ping{})
		// t.Log("ping send")
		time.Sleep(10 * time.Millisecond)
	}
}

func server(t *testing.T) {
	components.Register(&TestHandler{metrics: 0})
	nano.Listen(addr,
		nano.WithComponents(&components),
		nano.WithSerializer(protobuf.NewSerializer()),
	)
}
