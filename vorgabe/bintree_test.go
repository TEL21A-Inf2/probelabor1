package vorgabe

import (
	"testing"

	"github.com/tel21a-inf2/probelabor1/testhelpers"
)

func TestBinTree_Add(t *testing.T) {
	test := testhelpers.NewTest("Test für das Hinzufügen von Elementen zum Baum", t)

	t1 := NewBinTree()
	test.Assert(t1.Empty(), "Ein neuer Baum sollte leer sein.")

	t1.Add(42)
	test.AssertIntsEqual(42, t1.Value)
	test.Assert(t1.Left.Empty(), "Das linke Kind der Wurzel sollte leer sein.")
	test.Assert(t1.Right.Empty(), "Das rechte Kind der sollte leer sein.")

	t1.Add(23)
	t1.Add(55)
	t1.Add(77)
	test.AssertIntsEqual(42, t1.Value)
	test.AssertIntsEqual(23, t1.Left.Value)
	test.AssertIntsEqual(55, t1.Right.Value)
	test.AssertIntsEqual(77, t1.Right.Right.Value)
}

func TestBintree_String(t *testing.T) {
	test := testhelpers.NewTest("Test für die String-Repräsentation des Baumes", t)

	t1 := NewBinTree()
	test.AssertStringsEqual("()", t1.String())

	t1.Add(42)
	test.AssertStringsEqual("(42 () ())", t1.String())

	t1.Add(23)
	t1.Add(55)
	t1.Add(77)
	test.AssertStringsEqual("(42 (23 () ()) (55 () (77 () ())))", t1.String())
}
