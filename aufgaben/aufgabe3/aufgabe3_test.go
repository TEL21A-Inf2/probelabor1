package aufgabe3

import "fmt"

func ExampleBinTree_DecreaseSmallest_nonEmptyTree() {
	t1 := NewBinTree()

	t1.Add(42)
	t1.Add(23)
	t1.Add(55)
	t1.Add(77)

	fmt.Println(t1)
	t1.DecreaseSmallest()
	fmt.Println(t1)

	// Output:
	// (42 (23 () ()) (55 () (77 () ())))
	// (42 (22 () ()) (55 () (77 () ())))
}

func ExampleBinTree_DecreaseSmallest_emptyTree() {
	t1 := NewBinTree()

	fmt.Println(t1)
	t1.DecreaseSmallest()
	fmt.Println(t1)

	// Output:
	// ()
	// ()
}
