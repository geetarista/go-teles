package teles

import (
	"testing"
)

var (
	serverHost    = "192.168.36.129"
	serverPort    = "2856"
	serverAddress = serverHost + ":" + serverPort
	dummySpace    = Space{Name: "asdf"}
	validSpace    = Space{
		Name: "thing",
		Conn: &Connection{Server: serverAddress},
	}
	anotherSpace = Space{
		Name: "another",
		Conn: &Connection{Server: serverAddress},
	}
)

func failIfError(t *testing.T, err error) {
	if err != nil {
		t.Error(err.Error())
	}
}
