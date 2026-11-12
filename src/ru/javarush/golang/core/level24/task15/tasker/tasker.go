package tasker

// TaskStore объявляется "там, где используют": в бизнес-логике.
type TaskStore interface {
	Add(text string) error
	List() []string
}

type Service struct {
	// Service зависит только от интерфейса, а не от конкретной реализации хранилища.
	store TaskStore
}

func NewService(store TaskStore) *Service {
	return &Service{store: store}
}

func (s *Service) AddTask(text string) error {
	// Вся логика добавления — делегируется хранилищу.
	return s.store.Add(text)
}
