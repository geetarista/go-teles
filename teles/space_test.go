package teles

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ok, err := validSpace.Add("derp")
	failIfError(t, err)
	if ok != true {
		t.Error(ok)
	}
}

func TestAssociate(t *testing.T) {
	ok, err := validSpace.Associate("derp", 40.5, -120.5)
	failIfError(t, err)
	if ok != true {
		t.Error(ok)
	}
}

func TestListObjects(t *testing.T) {
	objects, err := validSpace.ListObjects()
	failIfError(t, err)
	if objects[0] != "derp" {
		t.Error("derp not found")
	}
}

func TestListAssociations(t *testing.T) {
	objects, err := validSpace.ListAssociations("derp")
	failIfError(t, err)
	if objects["2057777106"][0] != "40.5000" {
		t.Error("wrong lat")
	}
	if objects["2057777106"][1] != "-120.5000" {
		t.Error("wrong lng")
	}
}

func TestDisassociate(t *testing.T) {
	ok, err := validSpace.Disassociate("2057777106", "derp")
	failIfError(t, err)
	if ok != true {
		t.Error(ok)
	}
	validSpace.Associate("derp", 40.5, -120.5)
}

func TestQueryWithin(t *testing.T) {
	objects, err := validSpace.QueryWithin(0.0, 50.0, -150.0, -100.0)
	failIfError(t, err)
	if objects[0] != "derp" {
		t.Error("not found within")
	}
}

func TestQueryAround(t *testing.T) {
	objects, err := validSpace.QueryAround(40.0, -120.0, 50.0, "mi")
	failIfError(t, err)
	if objects[0] != "derp" {
		t.Error("not found around")
	}
}

func TestQueryNearest(t *testing.T) {
	objects, err := validSpace.QueryNearest(40.0, -120.0, 1)
	failIfError(t, err)
	if objects[0] != "derp" {
		t.Error("not found around")
	}
}
