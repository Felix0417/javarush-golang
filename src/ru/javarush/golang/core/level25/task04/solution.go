package main

/*
Вы пишете маленькое хранилище map[any]string для “меток” в отладочном инструменте. Но есть ловушка: если кто-то попытается
использовать в качестве ключа “нехешируемую” штуку (например, слайс), программа может упасть с panic. Ваш план — сделать
безопасную запись, которая не рушит приложение.

Реализуйте функцию safePut(labels map[any]string, key any, value string) (ok bool). Внутри функция должна пытаться
выполнить присваивание labels[key] = value. Если во время этой операции случается panic, функция должна перехватить её
через recover и вернуть ok=false. Если паники нет, вернуть ok=true.

В main создайте map[any]string и выполните три записи через safePut: с ключом 10 и значением "ten", затем с ключом "x" и
значением "ex", затем с ключом []int{1, 2} и значением "slice". После этого выведите, чем закончилась каждая попытка
(как ok или panic), и итоговый размер карты.

Программа должна вывести 4 строки:
put int: <ok|panic>
put string: <ok|panic>
put slice: <ok|panic>
size=<количество элементов в map после всех попыток>
*/

import "fmt"

// safePut выполняет безопасную запись в map и сообщает, получилось ли выполнить операцию без panic.
func safePut(labels map[any]string, key any, value string) (ok bool) {
	defer func() {
		// TODO: Добавьте recover() в deferred-функцию и установите ok=false, если при записи произошла panic.
		if r := recover(); r != nil {
			ok = false
		}
	}()

	// TODO: Попытайтесь выполнить реальную запись в map: labels[key] = value.
	// TODO: Если паники не было — верните ok=true.
	labels[key] = value
	ok = true
	return ok
}

func main() {
	labels := make(map[any]string)

	okInt := safePut(labels, 10, "ten")
	if okInt {
		fmt.Println("put int: ok")
	} else {
		fmt.Println("put int: panic")
	}

	okString := safePut(labels, "x", "ex")
	if okString {
		fmt.Println("put string: ok")
	} else {
		fmt.Println("put string: panic")
	}

	okSlice := safePut(labels, []int{1, 2}, "slice")
	if okSlice {
		fmt.Println("put slice: ok")
	} else {
		fmt.Println("put slice: panic")
	}

	fmt.Printf("size=%d\n", len(labels))
}
