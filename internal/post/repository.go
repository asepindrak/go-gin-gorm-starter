package post

import "gorm.io/gorm"

type Repository interface {
	Create(post Post) (Post, error)
	List() ([]Post, error)
	Get(id uint) (Post, error)
	Update(id uint, input Post) (Post, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(post Post) (Post, error) {
	if err := r.db.Create(&post).Error; err != nil {
		return Post{}, err
	}
	return post, nil
}

func (r *repository) List() ([]Post, error) {
	var posts []Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *repository) Get(id uint) (Post, error) {
	var post Post
	if err := r.db.First(&post, id).Error; err != nil {
		return Post{}, err
	}
	return post, nil
}

func (r *repository) Update(id uint, input Post) (Post, error) {
	var post Post
	if err := r.db.First(&post, id).Error; err != nil {
		return Post{}, err
	}
	post.Title = input.Title
	if err := r.db.Save(&post).Error; err != nil {
		return Post{}, err
	}
	return post, nil
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Post{}, id).Error
}
