package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/libp2p/go-libp2p-core/network"
)

const (
	protocolShakehand = "/shakehand/1.0.0"
	protocolChat      = "/chat/1.0.0"
)

type streamInfo struct {
	Host bool   `json:"host"`
	ID   string `json:"id"`
}

func WriteInfo(writer io.Writer, info streamInfo) error {
	b := make([]byte, 1000)
	data := streamInfo{
		Host: true,
	}
	bb, _ := json.Marshal(data)
	copy(b, bb)
	_, err := writer.Write(b)
	return err
}

func ReadInfo(r io.Reader, info *streamInfo) error {
	b := make([]byte, 1000)
	n, _ := r.Read(b)
	if n != 1000 {
		return fmt.Errorf("not enough")
	}
	return json.Unmarshal(b, info)
}

func handleStream(s network.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw)
	go writeData(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}

func readData(rw *bufio.ReadWriter) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
		}

	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		rw.WriteString(fmt.Sprintf("%s\n", sendData))
		rw.Flush()
	}
}
