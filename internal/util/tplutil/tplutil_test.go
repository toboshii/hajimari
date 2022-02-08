package tplutil

import (
	"testing"
	"github.com/toboshii/hajimari/internal/hajimari"
)

// IsStatusCheckEnabled works when Status is set
func TestIsStatusCheckEnabled(t *testing.T) {
	status := "true"
	app := hajimari.App {
		Name : "myApp",
		Icon : "mdi:flower",
		Group : "myGroup",
		URL : "www",
		Status : status,
	}
	want := true
	check := IsStatusCheckEnabled(app)

	if want != check {
		t.Fail()
	}
}

// IsStatusCheckEnabled works when Status is not set
func TestIsStatusCheckEnabledOnEmpty(t *testing.T) {
	app := hajimari.App {
		Name : "myApp",
		Icon : "mdi:flower",
		Group : "myGroup",
		URL : "www",
		Status: "undefined",
	}
	want := false
	check := IsStatusCheckEnabled(app)

	if want != check {
		t.Fail()
	}
}
