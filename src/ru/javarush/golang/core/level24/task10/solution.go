package main

/*
Представьте конвейер в обработке данных: один компонент выдаёт числа по одному (источник), другой принимает числа и
складывает их “в коробку” (приёмник). А посередине — ваша функция‑фильтр, которая пропускает только чётные значения,
не зная ничего о деталях хранения.

Опишите интерфейс IntSource с методом Next() (int, bool): он возвращает очередное число и true, а когда числа
закончились — любое число и false. Опишите интерфейс IntSink с методом Put(int). Напишите функцию
CopyEvens(src IntSource, dst IntSink), которая читает числа из src и передаёт в dst только чётные.

Сделайте реализацию SliceSource (хранит []int и текущий индекс) и реализацию SliceSink (хранит принятые числа в []int).
В main создайте источник со значениями 1 2 3 4 6 7, выполните CopyEvens и выведите итоговый слайс приёмника
через fmt.Println.
Ожидаемый вывод:
[2 4 6]
*/

import "fmt"

// IntSource отдаёт числа по одному; false = данных больше нет.
type IntSource interface {
	Next() (int, bool)
}

// IntSink принимает числа "в коробку".
type IntSink interface {
	Put(int)
}

// CopyEvens работает только через интерфейсы: читает из src и кладёт в dst только чётные.
func CopyEvens(src IntSource, dst IntSink) {
	// TODO: Реализуйте копирование только чётных чисел из src в dst через интерфейсы.
	// TODO: Читайте значения в цикле через src.Next() и останавливайтесь, когда ok == false.
	// TODO: Передавайте в dst.Put(...) только подходящие значения, сохраняя порядок.
	for {
		i, ok := src.Next()
		if !ok {
			break
		}

		if i%2 == 0 {
			dst.Put(i)
		}
	}

}

// SliceSource — источник из слайса с текущим индексом.
type SliceSource struct {
	data []int
	i    int
}

func (s *SliceSource) Next() (int, bool) {
	// TODO: Реализуйте выдачу следующего элемента из s.data по индексу s.i.
	// TODO: Не забудьте корректно обновлять индекс и возвращать (любое число, false) при завершении данных.
	if s.i >= len(s.data) {
		return 0, false
	}
	v := s.data[s.i]
	s.i++
	return v, true
}

// SliceSink — приёмник, накапливающий значения в порядке поступления.
type SliceSink struct {
	Values []int
}

func (s *SliceSink) Put(v int) {
	s.Values = append(s.Values, v)
}

func main() {
	src := &SliceSource{data: []int{1, 2, 3, 4, 6, 7}}
	dst := &SliceSink{}

	CopyEvens(src, dst)

	fmt.Println(dst.Values)
}
