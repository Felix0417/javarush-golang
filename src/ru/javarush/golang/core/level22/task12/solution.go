package main

import "fmt"

type Job struct {
	Done bool
}

// MarkDone должен менять исходный объект (элемент слайса), поэтому receiver — *Job.
func (j *Job) MarkDone() {
	// TODO: Реализуйте метод так, чтобы он помечал задачу выполненной (Done = true).
	j.Done = true
}

func main() {
	var n int
	fmt.Scan(&n)

	// По требованию: создаём строго через make([]Job, n).
	jobs := make([]Job, n)

	// TODO: Исправьте обход: нужно идти по индексам (for i := range jobs)
	// и вызывать метод на реальном элементе jobs[i]. Иначе меняется не тот объект.
	for i := range jobs {
		jobs[i].MarkDone()
	}

	for i := range jobs {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(jobs[i].Done)
	}
	fmt.Println()
}
