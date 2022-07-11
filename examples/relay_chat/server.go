package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
)

var ss = map[string]string{}

func RunServer(port int) error {
	h, err := libp2p.New(libp2p.ListenAddrStrings("0.0.0.0:" + fmt.Sprint(port)))
	if err != nil {
		return err
	}
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
