package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	playlist := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &playlist[i])
	}

	var insertIndex, trackID int
	fmt.Fscan(in, &insertIndex, &trackID)

	if 0 <= insertIndex && insertIndex <= n {
		// TODO: Реализуйте вставку trackID в playlist по индексу insertIndex без создания второго слайса размера O(N).
		// TODO: Вставка должна быть выполнена через увеличение длины append, затем сдвиг хвоста вправо через copy, затем запись trackID в позицию insertIndex.
		playlist = append(playlist, 0)
		copy(playlist[insertIndex+1:], playlist[insertIndex:])
		playlist[insertIndex] = trackID
	} else {
		// TODO: Обработайте некорректный индекс: добавьте trackID в конец без паники и без выхода за границы.
		playlist = append(playlist, trackID)
	}

	for i, v := range playlist {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
}
