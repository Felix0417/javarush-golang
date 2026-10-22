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

	var taskCount int
	fmt.Fscan(in, &taskCount)

	var taskIDs []int

	// TODO: Выведите стартовое состояние слайса taskIDs сразу после объявления
	// TODO: в формате: isNil len cap (через пробел).
	fmt.Fprintln(out, taskIDs == nil, len(taskIDs), cap(taskIDs))

	for i := 0; i < taskCount; i++ {
		var x int
		fmt.Fscan(in, &x)

		// TODO: Добавьте x в слайс taskIDs через append с сохранением результата
		// TODO: (нужно именно присваивание в taskIDs).
		taskIDs = append(taskIDs, x)
	}

	// TODO: Во второй строке выведите элементы итогового taskIDs,
	// TODO: разделяя числа одиночным пробелом.
	// TODO: Если taskCount == 0, вторая строка должна быть пустой (только перевод строки).
	if len(taskIDs) != 0 {
		fmt.Fprintln(out, taskIDs)
		return
	}
	fmt.Fprintln(out)
}
