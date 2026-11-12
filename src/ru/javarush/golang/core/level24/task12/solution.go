package main

/*
Вы делаете мини‑хранилище пользователей по ID. В реальном приложении “сохранить пользователя” и “загрузить пользователя” —
разные сценарии, и часто их удобно держать отдельными контрактами. Но иногда нужен общий контракт — тогда вы собираете
его из маленьких интерфейсов через embedding.

Создайте структуру User с полями ID int и Name string. Опишите интерфейс UserSaver с методом Save(User) error и интерфейс
UserLoader с методом Load(id int) (User, bool). Соберите интерфейс UserRepo через embedding: он должен включать UserSaver
и UserLoader.

Реализуйте MemoryUserRepo на базе map[int]User. Save возвращает ошибку "invalid user", если ID <= 0 или Name пустая строка,
и ошибку "user already exists", если такой ID уже сохранён. Load возвращает (User{}, false), если пользователя нет.

Напишите функции Register(s UserSaver, u User) и Show(l UserLoader, id int) с выводом:
Register: "registered: ID Name" или "register error: <ошибка>"
Show: "user: ID Name" или "user not found: ID"

В main выполните:
Register(repo, User{ID: 1, Name: "Ann"})
Register(repo, User{ID: 1, Name: "Bob"})
Show(repo, 1)
Show(repo, 2)

Ожидаемый вывод:
registered: 1 Ann
register error: user already exists
user: 1 Ann
user not found: 2
*/

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

// Маленькие интерфейсы под разные сценарии.
type UserSaver interface {
	Save(User) error
}

type UserLoader interface {
	Load(id int) (User, bool)
}

// Общий контракт собираем через embedding.
type UserRepo interface {
	UserSaver
	UserLoader
}

type MemoryUserRepo struct {
	users map[int]User
}

func NewMemoryUserRepo() *MemoryUserRepo {
	return &MemoryUserRepo{users: make(map[int]User)}
}

func (r *MemoryUserRepo) Save(u User) error {
	// TODO: Реализуйте все правила Save: валидацию (ID и Name), проверку уникальности ID,
	// TODO: точные тексты ошибок и запрет перезаписи уже сохранённого пользователя.

	if u.ID <= 0 || u.Name == "" {
		return errors.New("invalid user")
	}
	if r.users == nil {
		r.users = make(map[int]User)
	}

	if _, ok := r.Load(u.ID); ok {
		return errors.New("user already exists")
	}

	r.users[u.ID] = u
	return nil
}

func (r *MemoryUserRepo) Load(id int) (User, bool) {
	// TODO: Реализуйте загрузку пользователя по id из map с корректным флагом наличия (ok-идиома).
	if r.users == nil || len(r.users) == 0 {
		return User{}, false
	}

	u, ok := r.users[id]
	return u, ok

}

func Register(s UserSaver, u User) {
	if err := s.Save(u); err != nil {
		fmt.Printf("register error: %v\n", err)
		return
	}
	fmt.Printf("registered: %d %s\n", u.ID, u.Name)
}

func Show(l UserLoader, id int) {
	u, ok := l.Load(id)
	if !ok {
		fmt.Printf("user not found: %d\n", id)
		return
	}
	fmt.Printf("user: %d %s\n", u.ID, u.Name)
}

func main() {
	repo := NewMemoryUserRepo()

	Register(repo, User{ID: 1, Name: "Ann"})
	Register(repo, User{ID: 1, Name: "Bob"})
	Show(repo, 1)
	Show(repo, 2)
}
