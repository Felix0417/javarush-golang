package main

/*
Вы делаете консольный менеджер задач «как в настоящем приложении»: проект хранит список задач и сам назначает им ID, но
наружу не должен торчать ни генератор ID, ни внутренний слайс задач. То есть никакого embedding для этих деталей: внешний
код должен общаться только через методы, а данные получать только копией.

Напишите программу с типами Task (ID int, Title string, Done bool) и IDGen (поле next int, метод Next() int у *IDGen).
Создайте экспортируемый тип Project, который хранит задачи и выдаёт им ID, но не раскрывает наружу генератор и внутреннее
хранилище через embedding.

Project должен содержать Name string, неэкспортируемые именованные поля tasks и idGen (не embedded), а также методы:
	AddTask(title string) (Task, error): пустой title запрещён, ошибка должна иметь текст empty title
	SetDone(id int) bool: помечает задачу Done=true, возвращает true если задача найдена
	Tasks() []Task: возвращает копию списка задач

В main добавьте задачи milk, room, затем добавьте задачу с пустым заголовком, затем пометьте задачу с id=2 как выполненную. Программа должна вывести:
task:1 milk
task:2 room
error: empty title
done=1
*/

import (
	"errors"
	"fmt"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type IDGen struct {
	next int
}

func (g *IDGen) Next() int {
	// TODO: Реализуйте генерацию следующего ID (увеличить внутреннее состояние и вернуть новое значение).
	g.next++
	return g.next
}

type Project struct {
	Name  string
	tasks []Task // не embedded: наружу не должен торчать внутренний слайс
	idGen IDGen  // не embedded: наружу не должен торчать генератор
}

func (p *Project) AddTask(title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("empty title")
	}

	// TODO: Создайте задачу с корректным ID (через генератор), добавьте её во внутреннее хранилище и верните её наружу.
	t := Task{
		ID:    p.idGen.Next(),
		Title: title,
		Done:  false,
	}

	p.tasks = append(p.tasks, t)

	return t, nil
}

func (p *Project) SetDone(id int) bool {
	// TODO: Найдите задачу по ID во внутреннем хранилище, выставьте Done=true и верните true. Если не нашли — верните false.
	for i, t := range p.tasks {
		if t.ID == id {
			p.tasks[i].Done = true
			return true
		}
	}
	return false
}

func (p *Project) Tasks() []Task {
	// TODO: Верните копию слайса задач, чтобы внешний код не мог изменить внутреннее хранилище проекта.
	tasks := make([]Task, len(p.tasks))
	copy(tasks, p.tasks)
	return tasks
}

func main() {
	var p Project
	p.Name = "demo"

	t1, _ := p.AddTask("milk")
	fmt.Printf("task:%d %s\n", t1.ID, t1.Title)

	t2, _ := p.AddTask("room")
	fmt.Printf("task:%d %s\n", t2.ID, t2.Title)

	_, err := p.AddTask("")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	ok := p.SetDone(2)
	if ok {
		fmt.Println("done=1")
	} else {
		fmt.Println("done=0")
	}
}
