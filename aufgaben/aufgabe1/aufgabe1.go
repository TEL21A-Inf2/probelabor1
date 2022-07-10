// Thema der Aufgabe: Listen
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe1

/* Aufgabenstellung:
 * In der Datei linkedlist.go  ist ein Datentyp für eine verkettete Liste definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode Erase().
 * Die Methode soll eine Position erwarten und aus der Liste das Element an dieser
 * Stelle entfernen. Kommt das Element nicht vor, soll die Funktion nichts machen.
 */
func (list *LinkedList) Erase(pos int) {
	if list.Empty() {
		return
	}
	if pos == 0 {
		list.Id = list.Next.Id
		pos = 1
	}

	current := list
	for i := 1; !current.Next.Empty(); i++ {
		if i == pos {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
