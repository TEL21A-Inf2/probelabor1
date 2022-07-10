package testhelpers

import "testing"

type Test struct {
	t    *testing.T
	Name string
}

// Liefert ein neues Test-Objekt mit dem angegebenen Namen.
func NewTest(name string, t *testing.T) *Test {
	return &Test{t, name}
}

// Prüft, ob zwei Strings gleich sind.
func (t *Test) AssertStringsEqual(expected, actual string) {
	if expected != actual {
		t.t.Errorf("\nWerte sind nicht gleich!\n  Erwartet: %v\n  Tatsächlich: %v\n", expected, actual)
	}
}
