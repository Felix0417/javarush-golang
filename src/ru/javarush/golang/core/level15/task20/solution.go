package main

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	fmt.Scan(&m)

	// online — set через map[string]struct{}
	online := map[string]struct{}{}

	// inCount — счётчик входов (IN) по пользователям
	inCount := map[string]int{}

	for i := 0; i < m; i++ {
		var action, userName string
		fmt.Scan(&action, &userName)

		if action == "IN" {
			// TODO: Обработайте вход пользователя: добавьте userName в online-set и увеличьте счётчик входов в inCount.
			online[userName] = struct{}{}
			inCount[userName]++
		} else if action == "OUT" {
			// TODO: Обработайте выход пользователя: удалите userName из online-set через delete (даже если ключа нет).
			delete(online, userName)
		}
	}

	// TODO: Сформируйте список онлайн-пользователей из online (ключи map), отсортируйте и выведите:
	// - onlineCount
	// - onlineCount строк с именами пользователей (лексикографически)
	//
	// Временный код ниже оставлен, чтобы шаблон был минимально запускаемым.
	onlineUsers := make([]string, 0)
	for k := range online {
		onlineUsers = append(onlineUsers, k)
	}

	sort.Strings(onlineUsers)

	fmt.Println(len(onlineUsers))

	// TODO: Выведите имена онлайн-пользователей (по одному в строке) после строки с onlineCount.
	for _, k := range onlineUsers {
		fmt.Println(k)
	}

	// TODO: Найдите maxIn как максимум по значениям inCount и выведите maxIn.
	maxIn := 0
	for k := range inCount {
		if maxIn < inCount[k] {
			maxIn = inCount[k]
		}
	}
	fmt.Println(maxIn)

	// TODO: Сформируйте список лидеров (всех пользователей, у кого входов == maxIn),
	// отсортируйте и выведите по одному имени в строке.
	leaders := make([]string, 0)
	for k, v := range inCount {
		if v == maxIn {
			leaders = append(leaders, k)
		}
	}
	sort.Strings(leaders)

	for _, k := range leaders {
		fmt.Println(k)
	}
}
