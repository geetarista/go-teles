package teles

import (
	"testing"
)

func TestCreateSocket(t *testing.T) {
	conn := Connection{Server: serverAddress}
	conn.createSocket()
	if conn.Socket == nil {
		t.Error("Failed to connect")
	}
	conn.Socket.Close()
}

func TestSend(t *testing.T) {
	conn := Connection{Server: serverAddress}
	defer conn.Socket.Close()
	err := conn.Send("derp")
	failIfError(t, err)
}

func TestReadDerp(t *testing.T) {
	conn := Connection{Server: serverAddress}
	defer conn.Socket.Close()
	err := conn.Send("list spaces")
	failIfError(t, err)
	resp, err := conn.Read()
	failIfError(t, err)
	if resp != "START" {
		t.Error("Got: " + resp)
	}
}
