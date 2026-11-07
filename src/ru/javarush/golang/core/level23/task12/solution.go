package main

/*
Вы делаете маленькое key-value хранилище для консоли: сохранить значение, загрузить по ключу, удалить, а ещё —
вывести «дамп» в стабильном порядке, чтобы лог сравнения в тестах не прыгал. Контракты вы собираете через embedding:
сначала базовые операции, а потом поверх них административные возможности.

Объявите интерфейсы Saver (Save(key, value string) error), Loader (Load(key string) (string, bool)),
Deleter (Delete(key string) error). Объявите интерфейс KeyValueStore, который включает Saver, Loader и Deleter через
embedding. Дополнительно объявите интерфейс Dumper с методом Dump() []string и интерфейс AdminStore, который включает
KeyValueStore и Dumper через embedding.

Реализуйте тип MapStore на базе map[string]string. Dump должен возвращать слайс строк вида key=value, отсортированный
по key по возрастанию. Программа читает из stdin число N, затем N пар key value (оба без пробелов), затем ключ delKey.
После этого программа удаляет delKey и печатает дамп: по одной строке на элемент.
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Saver — базовая операция записи.
type Saver interface {
	Save(key, value string) error
}

// Loader — базовая операция чтения.
type Loader interface {
	Load(key string) (string, bool)
}

// Deleter — базовая операция удаления.
type Deleter interface {
	Delete(key string) error
}

// KeyValueStore собирает базовые контракты через embedding.
type KeyValueStore interface {
	Saver
	Loader
	Deleter
}

// Dumper — административная возможность: стабильный дамп.
type Dumper interface {
	Dump() []string
}

// AdminStore — расширенный контракт: базовые операции + дамп.
type AdminStore interface {
	KeyValueStore
	Dumper
}

// MapStore — простая реализация на базе map[string]string.
type MapStore struct {
	data map[string]string
}

func (s *MapStore) Save(key, value string) error {
	// map может быть nil у zero-value — создаём при первом сохранении
	if s.data == nil {
		s.data = make(map[string]string)
	}
	s.data[key] = value
	return nil
}

func (s *MapStore) Load(key string) (string, bool) {
	v, ok := s.data[key]
	return v, ok
}

func (s *MapStore) Delete(key string) error {
	// TODO: Реализуйте удаление ключа из хранилища. Удаление отсутствующего ключа не должно приводить к ошибке.
	delete(s.data, key)
	return nil
}

func (s *MapStore) Dump() []string {
	// TODO: Реализуйте дамп хранилища: верните []string в формате key=value, отсортированный по key по возрастанию.
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := make([]string, 0, len(keys))
	for i := range keys {
		result = append(result, fmt.Sprintf("%s=%s", keys[i], s.data[keys[i]]))
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var store MapStore
	for i := 0; i < n; i++ {
		var key, value string
		fmt.Fscan(in, &key, &value)
		store.Save(key, value)
	}

	var delKey string
	fmt.Fscan(in, &delKey)
	store.Delete(delKey)

	for _, line := range store.Dump() {
		fmt.Fprintln(out, line)
	}
}
