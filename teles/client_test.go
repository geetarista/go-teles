package teles

import (
	"fmt"
	"testing"
)

// Clear everything out of teles before running tests.
func TestDropEverything(t *testing.T) {
	client := NewClient(serverAddress)
	defer client.Conn.Socket.Close()
	spaces, _ := client.ListSpaces()
	for _, s := range spaces {
		space := Space{Name: s, Conn: client.Conn}
		space.Delete()
	}
}

func TestCreateSpace(t *testing.T) {
	client := NewClient(serverAddress)
	defer client.Conn.Socket.Close()
	err := client.CreateSpace(&validSpace)
	failIfError(t, err)
	err = client.CreateSpace(&anotherSpace)
	failIfError(t, err)
}

func TestGetSpace(t *testing.T) {
	client := NewClient(serverAddress)
	defer client.Conn.Socket.Close()
	filter := client.GetSpace(validSpace.Name)
	if filter.Name != validSpace.Name {
		t.Error("Name not equal")
	}
}

func TestListSpaces(t *testing.T) {
	client := NewClient(serverAddress)
	defer client.Conn.Socket.Close()
	filters, err := client.ListSpaces()
	failIfError(t, err)
	if filters[0] == validSpace.Name {
		fmt.Printf("%+v\n", filters)
		t.Error(validSpace.Name)
	}
}
