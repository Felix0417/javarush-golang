package main

import "fmt"

// dump печатает слайс так, чтобы было видно len/cap и содержимое.
func dump(name string, s []string) {
	fmt.Printf("%s: len=%d cap=%d %v\n", name, len(s), cap(s), s)
}

func main() {
	// tasks: len=4, cap=4 (заполняем без литерала []string{...})
	tasks := make([]string, 4, 4)
	tasks[0] = "A"
	tasks[1] = "B"
	tasks[2] = "C"
	tasks[3] = "D"

	// TODO: Создайте под-слайс today ровно как tasks[:2], без копирования и без ограничения cap.
	today := tasks[:2]

	dump("tasks(before)", tasks)
	dump("today(before)", today)

	// TODO: Сохраните результат append обратно в today и добавьте строку "X".
	today = append(today, "X")

	dump("tasks(after)", tasks)
	dump("today(after)", today)
}
