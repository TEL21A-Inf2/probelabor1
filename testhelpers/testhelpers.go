package testhelpers

import (
	"testing"
)

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
	t.Assert(expected == actual, "\nWerte sind nicht gleich!\n  Erwartet: %v\n  Tatsächlich: %v\n", expected, actual)
}

// Prüft einen booleschen Wert.
// Gibt die angegebene Fehlermeldung aus, wenn die Bedingung nicht zutrifft.
// Die Fehlermeldung darf ein Formatstring sein, es werden die weiteren Werte eingesetzt.
func (t *Test) Assert(value bool, message string, formatValues ...any) {
	if !value {
		t.t.Errorf(message, formatValues...)
	}
}
