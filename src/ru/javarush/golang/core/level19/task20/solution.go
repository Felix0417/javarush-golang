package main

import "fmt"

// SnapshotPtrs возвращает слайс указателей на строки из tasks.
//
// Важно: указатели в результате должны быть независимы от исходного tasks.
func SnapshotPtrs(tasks []string) []*string {
	snaps := make([]*string, 0, len(tasks))

	for _, v := range tasks {
		// TODO: Сделайте так, чтобы каждый указатель в snaps указывал на отдельную копию строки,
		// TODO: а не на элемент исходного слайса tasks (и чтобы элементы снимка не были связаны друг с другом).
		snaps = append(snaps, &v)
	}

	return snaps
}

func main() {
	tasks := []string{"a", "b", "c"}

	snaps := SnapshotPtrs(tasks)

	// Меняем исходный слайс — снимок не должен измениться.
	tasks[0] = "X"

	// Меняем значение через указатель из снимка — tasks не должен измениться.
	*snaps[1] = "Y"

	fmt.Println(tasks)
	fmt.Println(*snaps[0], *snaps[1], *snaps[2])

}
