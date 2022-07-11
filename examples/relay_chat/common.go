package main

import (
	"encoding/json"
	"fmt"
	"io"
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
