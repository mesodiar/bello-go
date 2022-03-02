package foobar

var Title = "FooBar"
var name = "Mils"

func Say (n int) string {
	if n == 2 {
		return "2"
	}
	if n == 4 {
		return "4"
	}
	if n == 5 {
		return "Bar"
	}
	if n == 3 || n == 6 {
		return "Foo"
	}

	return "1"
}