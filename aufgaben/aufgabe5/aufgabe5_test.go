package aufgabe5

import "fmt"

func ExampleBucketList_Size() {
	b1 := NewBucketList()
	fmt.Println(b1.Size())

	b1.Add("Hallo")
	fmt.Println(b1.Size())
	b1.Add("Haus")
	fmt.Println(b1.Size())
	b1.Add("Auto")
	fmt.Println(b1.Size())
	b1.Add("Fahrrad")
	fmt.Println(b1.Size())

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}
