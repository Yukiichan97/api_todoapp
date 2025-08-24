package tasks

import "sync"

type Repository struct {
	tasks  map[int]Task
	nextID int
}

var (
	mu sync.Mutex
)

func NewRepository() *Repository {
	return &Repository{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (r *Repository) Create(task Task) Task {
	mu.Lock()
	defer mu.Unlock()
	task.ID = r.nextID
	r.nextID++
	r.tasks[task.ID] = task
	return task
}

func (r *Repository) GetAll() []Task {
	mu.Lock()
	defer mu.Unlock()
	list := make([]Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		list = append(list, t)
	}
	return list
}

func (r *Repository) Update(id int, updated Task) (Task, bool) {
	mu.Lock()
	defer mu.Unlock()
	_, ok := r.tasks[id]
	if !ok {
		return Task{}, false
	}
	updated.ID = id
	updated.Description = r.tasks[id].Description
	updated.Title = r.tasks[id].Title
	r.tasks[id] = updated
	return updated, true
}

func (r *Repository) Delete(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	_, ok := r.tasks[id]
	if !ok {
		return false
	}
	delete(r.tasks, id)
	return true
}
