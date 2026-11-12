package memstore

type Store struct {
	tasks []string
}

func (s *Store) Add(text string) error {
	// TODO: Реализуйте сохранение задачи в памяти так, чтобы List() возвращал задачи в порядке добавления.
	if s.tasks == nil {
		s.tasks = make([]string, 0)
	}
	s.tasks = append(s.tasks, text)
	return nil
}

func (s *Store) List() []string {
	// TODO: Реализуйте возврат всех сохранённых задач (в исходном порядке добавления).
	return s.tasks
}
