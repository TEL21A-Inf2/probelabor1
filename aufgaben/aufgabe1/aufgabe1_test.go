package aufgabe1

import "fmt"

func ExampleLinkedList_Erase_firstElement() {
	l1 := MakeLinkedList("A", "B", "C")

	l1.Erase(0)
	fmt.Println(l1)

	// Output:
	// [B C]
}

func ExampleLinkedList_Erase_middleElement() {
	l1 := MakeLinkedList("A", "B", "C")

	l1.Erase(1)
	fmt.Println(l1)

	// Output:
	// [A C]
}

func ExampleLinkedList_Erase_nonExistingElement() {
	l1 := MakeLinkedList("A", "B", "C")

	l1.Erase(5)
	fmt.Println(l1)

	l2 := MakeLinkedList()
	l2.Erase(5)
	fmt.Println(l2)
	l2.Erase(-1)
	fmt.Println(l2)

	// Output:
	// [A B C]
	// []
	// []
}
