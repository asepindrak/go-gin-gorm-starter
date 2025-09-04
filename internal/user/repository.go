package user

import "gorm.io/gorm"

type Repository interface {
	Create(u *User) error
	FindByID(id uint) (*User, error)
	FindAll(offset, limit int) ([]User, int64, error)
	Update(u *User) error
	Delete(id uint) error
}

type gormRepo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &gormRepo{db: db} }

func (r *gormRepo) Create(u *User) error { return r.db.Create(u).Error }

func (r *gormRepo) FindByID(id uint) (*User, error) {
	var u User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *gormRepo) FindAll(offset, limit int) ([]User, int64, error) {
	var (
		users []User
		total int64
	)
	r.db.Model(&User{}).Count(&total)
	q := r.db.Order("id DESC")
	if limit > 0 {
		q = q.Offset(offset).Limit(limit)
	}
	if err := q.Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *gormRepo) Update(u *User) error { return r.db.Save(u).Error }

func (r *gormRepo) Delete(id uint) error { return r.db.Delete(&User{}, id).Error }
