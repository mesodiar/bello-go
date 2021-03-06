package foobar_test

import "testing"
import "github.com/mesodiar/bello-go/foobar"

func TestFooBarGiven2WantsString2(t *testing.T) {
	given := 2
	wants := "2"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGiven1WantsString1(t *testing.T) {
	given := 1
	wants := "1"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGiven4WantsString4(t *testing.T) {
	given := 4
	wants := "4"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGiven5WantsStringBar(t *testing.T) {
	given := 5
	wants := "Bar"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGiven6WantsStringFoo(t *testing.T) {
	given := 6
	wants := "Foo"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGiven3WantsStringFoo(t *testing.T) {
	given := 3
	wants := "Foo"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGiven12WantsStringFoo(t *testing.T) {
	given := 12
	wants := "Foo"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}

func TestFooBarGiven15WantsStringFooBar(t *testing.T) {
	given := 15
	wants := "FooBar"

	get := foobar.Say(given)

	if wants != get {
		t.Errorf("given %d wants %q but got %q\n", given, wants, get)
	}
}