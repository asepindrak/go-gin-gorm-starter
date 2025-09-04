package user

import "gorm.io/gorm"

type Repository interface {
	Create(user User) (User, error)
	List() ([]User, error)
	Get(id uint) (User, error)
	Update(id uint, input User) (User, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) List() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) Get(id uint) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) Update(id uint, input User) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return User{}, err
	}
	user.Name = input.Name
	if err := r.db.Save(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
