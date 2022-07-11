package main

import (
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/multiformats/go-multiaddr"
)

var ss = map[string]string{}

func RunServer(port int) error {
	h, err := libp2p.New(libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)))
	if err != nil {
		return err
	}

	for _, la := range h.Network().ListenAddresses() {
		if p, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
			log.Printf("/ip4/127.0.0.1/tcp/%v/p2p/%s on another console.\n", p, h.ID().Pretty())
			//port = p
			//break
		}
	}

	//log.Printf("/ip4/127.0.0.1/tcp/%v/p2p/%s on another console.\n", port, h.ID().Pretty())
	log.Println("You can replace 127.0.0.1 with public IP as well.")
	log.Println("Waiting for incoming connection")
	log.Println()
	//mu := sync.Mutex{}
	var s1 network.Stream
	h.SetStreamHandler(protocolShakehand, func(stream network.Stream) {
		if s1 == nil {
			info := streamInfo{
				Host: true,
			}
			err := WriteInfo(stream, info)
			if err != nil {
				s1 = stream
			}
		} else {
			info := streamInfo{
				ID: s1.ID(),
			}
			WriteInfo(stream, info)
			s1 = nil
		}
	})
	return nil
}
