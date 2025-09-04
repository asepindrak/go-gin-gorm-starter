package db

import (
	"github.com/asepindrak/go-gin-gorm-starter/internal/post"
	"github.com/asepindrak/go-gin-gorm-starter/internal/user"
	"gorm.io/gorm"
)

// Migrate menambahkan semua model ke AutoMigrate
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&post.Post{},
	)
}
