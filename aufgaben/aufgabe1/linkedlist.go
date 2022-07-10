package aufgabe1

type LinkedList struct {
	Id   string
	Next *LinkedList
}

func MakeLinkedList(Ids ...string) *LinkedList {
	head := &LinkedList{"", nil}
	for i := len(Ids) - 1; i >= 0; i-- {
		newElement := &LinkedList{Ids[i], head}
		head = newElement
	}
	return head
}
