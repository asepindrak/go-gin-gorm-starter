package user

type Service interface {
	Create(input User) (User, error)
	List() ([]User, error)
	Get(id uint) (User, error)
	Update(id uint, input User) (User, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(input User) (User, error) {
	return s.repo.Create(input)
}

func (s *service) List() ([]User, error) {
	return s.repo.List()
}

func (s *service) Get(id uint) (User, error) {
	return s.repo.Get(id)
}

func (s *service) Update(id uint, input User) (User, error) {
	return s.repo.Update(id, input)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
