package user

import (
	"errors"
	"strings"
)

type Service interface {
	Create(name, email string) (*User, error)
	Get(id uint) (*User, error)
	List(page, pageSize int) ([]User, int64, error)
	Update(id uint, name, email string) (*User, error)
	Delete(id uint) error
}

type service struct{ repo Repository }

func NewService(r Repository) Service { return &service{repo: r} }

func (s *service) Create(name, email string) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(strings.ToLower(email))
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}
	u := &User{Name: name, Email: email}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Get(id uint) (*User, error) { return s.repo.FindByID(id) }

func (s *service) List(page, pageSize int) ([]User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return s.repo.FindAll(offset, pageSize)
}
func (s *service) Delete(id uint) error { return s.repo.Delete(id) }
