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
	h.SetStreamHandler(protocolChat, func(stream network.Stream) {
		handleStream(stream)
	})
	s, err := h.NewStream(context.Background(), p.ID, protocolShakehand)
	info := &streamInfo{}
	err = ReadInfo(s, info)
	if err != nil {
		return err
	}
	s.Close()
	if !info.Host {
		id, err := peer.IDFromString(info.ID)
		if err != nil {
			return err
		}
		s1, err := h.NewStream(context.Background(), id, protocolChat)
		if err != nil {
			return err
		}
		handleStream(s1)
	}
	return nil
}
