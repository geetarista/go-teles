/*
Provides a client abstraction around the Teles interface.

Example:
	client := teles.Client{Server: "10.0.0.30:8673"}
	space := teles.Space{Name: "coolspace"}
	if err := teles.CreateSpace(space); err != nil {
		// handle error
	}
	spaces, _ := teles.ListSpaces()
	fmt.Printf("%+v", spaces[0])
*/
package teles

type Client struct {
	Server     string
	Timeout    int
	Conn       *Connection
	ServerInfo string
	InfoTime   int
}

func NewClient(address string) Client {
	return Client{Server: address, Conn: &Connection{Server: address}}
}

// Creates a new space on the Teles server
func (c *Client) CreateSpace(s *Space) error {
	cmd := "create space " + s.Name

	resp, err := s.Conn.SendAndReceive(cmd)
	if err != nil {
		return err
	}
	if resp != "Done" {
		return errInvalidResponse(resp)
	}
	s.Conn = c.Conn
	return nil
}

// Gets a Space object based on the name
func (c *Client) GetSpace(name string) *Space {
	return &Space{Name: name, Conn: c.Conn}
}

// Lists all the available spaces
func (c *Client) ListSpaces() (responses []string, err error) {
	err = c.Conn.Send("list spaces")
	if err != nil {
		return
	}

	responses, err = c.Conn.ReadBlock()
	if err != nil {
		return
	}
	return responses, nil
}
