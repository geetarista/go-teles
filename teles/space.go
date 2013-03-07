// Provides an interface to a single Teles space

package teles

import (
	"fmt"
	"strings"
)

var (
	Units = [5]string{"m", "km", "mi", "y", "ft"}
)

type Space struct {
	Name string
	Conn *Connection
}

func (s *Space) prefix() string {
	return "in " + s.Name + " "
}

// Adds a new object
func (s *Space) Add(name string) (bool, error) {
	cmd := "add object " + name
	resp, err := s.Conn.SendAndReceive(s.prefix() + cmd)
	if err != nil {
		return false, err
	}
	if resp == "Done" {
		return true, nil
	}
	return false, errInvalidResponse(resp)
}

func (s *Space) Delete() (bool, error) {
	cmd := "delete object " + s.Name
	resp, err := s.Conn.SendAndReceive(s.prefix() + cmd)
	if err != nil {
		return false, err
	}
	if resp == "Done" {
		return true, nil
	}
	if resp == "Object does not exist" {
		return false, nil
	}
	return false, errInvalidResponse(resp)
}

// Associates an object with a lat/lng
func (s *Space) Associate(name string, lat, lng float64) (bool, error) {
	cmd := fmt.Sprintf("associate point %f %f with %s", lat, lng, name)
	resp, err := s.Conn.SendAndReceive(s.prefix() + cmd)
	if err != nil {
		return false, err
	}
	if resp == "Done" {
		return true, nil
	}
	if resp == "Object does not exist" {
		return false, nil
	}
	return false, errInvalidResponse(resp)
}

// Disassociates an object with a GID
func (s *Space) Disassociate(gid, name string) (bool, error) {
	cmd := fmt.Sprintf("disassociate %s with %s", gid, name)
	resp, err := s.Conn.SendAndReceive(s.prefix() + cmd)
	if err != nil {
		return false, err
	}
	if resp == "Done" {
		return true, nil
	}
	if resp == "Object does not exist" || resp == "GID not associated" {
		return false, nil
	}
	return false, errInvalidResponse(resp)
}

// List all objects in the space
func (s *Space) ListObjects() (responses []string, err error) {
	err = s.Conn.Send(s.prefix() + "list objects")
	if err != nil {
		return
	}

	responses, err = s.Conn.ReadBlock()
	if err != nil {
		return
	}
	return responses, nil
}

// List object associations
func (s *Space) ListAssociations(name string) (responses map[string][2]string, err error) {
	err = s.Conn.Send(s.prefix() + "list associations with " + name)
	if err != nil {
		return
	}

	resp, err := s.Conn.ReadBlock()
	if err != nil {
		return
	}
	responses = make(map[string][2]string)
	for _, line := range resp {
		split := strings.Split(line, " ")
		items := [2]string{split[1][4:], split[2][4:]}
		responses[split[0][4:]] = items
	}
	return responses, nil
}

// Queries within a bounding box
func (s *Space) QueryWithin(minLat, maxLat, minLng, maxLng float64) (responses []string, err error) {
	if minLat > maxLat || minLng > maxLng {
		return responses, &TelesError{ErrorString: "Minimum lat/lng must be less than maximum lat/lng!"}
	}
	cmd := fmt.Sprintf("query within %f %f %f %f", minLat, maxLat, minLng, maxLng)
	err = s.Conn.Send(s.prefix() + cmd)
	if err != nil {
		return
	}

	responses, err = s.Conn.ReadBlock()
	if err != nil {
		return
	}
	return responses, nil
}

// Queries around a point
func (s *Space) QueryAround(lat, lng, distance float64, unit string) (responses []string, err error) {
	if unit == "" {
		unit = "mi"
	}
	unitFail := true
	for _, u := range Units {
		if u == unit {
			unitFail = false
			break
		}
	}

	if unitFail {
		return responses, &TelesError{ErrorString: "Bad unit provided!"}
	}
	if distance <= 0 {
		return responses, &TelesError{ErrorString: "Bad distance provided!"}
	}
	cmd := fmt.Sprintf("query around %f %f for %f%s", lat, lng, distance, unit)
	err = s.Conn.Send(s.prefix() + cmd)
	if err != nil {
		return
	}

	responses, err = s.Conn.ReadBlock()
	if err != nil {
		return
	}
	return responses, nil
}

// Queries for nearest points around a lat/lng
func (s *Space) QueryNearest(lat, lng float64, num int) (responses []string, err error) {
	if num <= 0 {
		return responses, &TelesError{ErrorString: "Bad num provided!"}
	}
	cmd := fmt.Sprintf("query nearest %d to %f %f", num, lat, lng)
	err = s.Conn.Send(s.prefix() + cmd)
	if err != nil {
		return
	}

	responses, err = s.Conn.ReadBlock()
	if err != nil {
		return
	}
	return responses, nil
}
