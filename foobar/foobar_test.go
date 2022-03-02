package foobar_test

import "testing"
import "github.com/mesodiar/bello-go/foobar"


func TestFooBarGiven1WantsString1(t *testing.T) {
	given := 1
	wants := "1"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}