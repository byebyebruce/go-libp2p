package main

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

func RunClient(p peer.AddrInfo) error {
	h, err := libp2p.New(libp2p.ListenAddrs(), libp2p.EnableRelay())
	if err != nil {
		return err
	}
	if err := h.Connect(context.Background(), p); err != nil {
		return err
	}
	s, err := h.NewStream(context.Background(), p.ID, protocolShakehand)
	info := &streamInfo{}
	err := ReadInfo(s, info)
	if err != nil {
		return err
	}
	if info.Host {
		h.SetStreamHandler(protocolChat, func(stream network.Stream) {

		})
	}
	s.Read()
}
