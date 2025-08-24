package tasks

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(t Task) Task {
	return s.repo.Create(t)
}

func (s *Service) GetTasks() []Task {
	return s.repo.GetAll()
}

func (s *Service) UpdateTask(id int, t Task) (Task, bool) {
	return s.repo.Update(id, t)
}

func (s *Service) DeleteTask(id int) bool {
	return s.repo.Delete(id)
}
