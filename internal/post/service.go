package post

type Service interface {
	Create(input Post) (Post, error)
	List() ([]Post, error)
	Get(id uint) (Post, error)
	Update(id uint, input Post) (Post, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(input Post) (Post, error) {
	return s.repo.Create(input)
}

func (s *service) List() ([]Post, error) {
	return s.repo.List()
}

func (s *service) Get(id uint) (Post, error) {
	return s.repo.Get(id)
}

func (s *service) Update(id uint, input Post) (Post, error) {
	return s.repo.Update(id, input)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
