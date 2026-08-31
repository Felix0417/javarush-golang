package main

import "fmt"

// Слайс — это view (ptr+len+cap) на backing array.
// Если намеренно ограничить cap, то append не сможет "расти" в общий массив
// и будет вынужден выделить новый буфер.
func viewFirstTwo(all []string) []string {
	// TODO: Реализуйте безопасный "view" на первые два элемента (или на весь all, если len < 2).
	// TODO: Важно намеренно ограничить ёмкость результата так, чтобы cap(result) == len(result),
	// TODO: и append к result не мог изменить исходный backing array.
	//
	// Ниже оставлена упрощённая версия, которая НЕ выполняет требование по ограничению cap.
	if len(all) < 2 {
		return all[:len(all):len(all)]
	}
	return all[:2:2]
}

func main() {
	all := make([]string, 4, 4)
	all[0] = "A"
	all[1] = "B"
	all[2] = "C"
	all[3] = "D"

	v := viewFirstTwo(all)

	fmt.Printf("v(initial)=%v len=%d cap=%d\n", v, len(v), cap(v))
	fmt.Printf("all(initial)=%v\n", all)

	grown := append(v, "X")

	fmt.Printf("all(after append)=%v\n", all)
	fmt.Printf("grown(after append)=%v\n", grown)

	// Изменение по индексу — это всё тот же backing array, поэтому all изменится.
	v[0] = "Z"

	fmt.Printf("all(after edit)=%v\n", all)
	fmt.Printf("v(after edit)=%v len=%d cap=%d\n", v, len(v), cap(v))
}
