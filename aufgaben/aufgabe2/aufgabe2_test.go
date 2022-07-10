package aufgabe2

import "fmt"

func ExampleLinkedList_Find_firstElement() {
	l1 := MakeLinkedList("A", "B", "C")

	fmt.Println(l1.Find("A"))

	// Output:
	// 0
}

func ExampleLinkedList_Find_middleElement() {
	l1 := MakeLinkedList("A", "B", "C")

	fmt.Println(l1.Find("B"))

	// Output:
	// 1
}

func ExampleLinkedList_Find_nonExistingElement() {
	l1 := MakeLinkedList("A", "B", "C")
	fmt.Println(l1.Find("D"))

	l2 := MakeLinkedList()
	fmt.Println(l2.Find("B"))

	// Output:
	// 3
	// 0
}
