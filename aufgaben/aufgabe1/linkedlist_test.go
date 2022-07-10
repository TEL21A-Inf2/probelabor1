package aufgabe1

import (
	"testing"

	"github.com/tel21a-inf2/probelabor1/testhelpers"
)

func TestLinkedList_MakeLinkedList(t *testing.T) {
	test := testhelpers.NewTest("Test für das Erzeugen neuer Listen mittels MakeLinkedList", t)

	l1 := MakeLinkedList("A", "B", "C")
	test.AssertStringsEqual(l1.Id, "A")
	test.AssertStringsEqual(l1.Next.Id, "B")
	test.AssertStringsEqual(l1.Next.Next.Id, "C")
	test.AssertStringsEqual(l1.Next.Next.Next.Id, "")

	l2 := MakeLinkedList()
	test.AssertStringsEqual(l2.Id, "")
}
