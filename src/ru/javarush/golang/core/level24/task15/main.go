package main

import (
	"fmt"

	"myproject/memstore"
	"myproject/tasker"
)

func main() {
	// Склейка зависимостей делается в main.
	var store tasker.TaskStore = &memstore.Store{}

	svc := tasker.NewService(store)

	_ = svc.AddTask("первая задача")
	_ = svc.AddTask("вторая задача")

	for _, t := range store.List() {
		fmt.Println(t)
	}
}
