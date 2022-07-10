package aufgabe4

import "fmt"

func ExampleBinTree_HasOddElementNumber() {
	t1 := NewBinTree()
	fmt.Println(t1.HasOddElementNumber())

	t1.Add(42)
	fmt.Println(t1.HasOddElementNumber())
	t1.Add(23)
	fmt.Println(t1.HasOddElementNumber())
	t1.Add(55)
	fmt.Println(t1.HasOddElementNumber())
	t1.Add(77)
	fmt.Println(t1.HasOddElementNumber())

	// Output:
	// false
	// true
	// false
	// true
	// false
}
