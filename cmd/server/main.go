package main

import (
	"log"

	"github.com/asepindrak/go-gin-gorm-starter/internal/config"
	"github.com/asepindrak/go-gin-gorm-starter/internal/db"
	"github.com/asepindrak/go-gin-gorm-starter/internal/post"
	"github.com/asepindrak/go-gin-gorm-starter/internal/router"
	"github.com/asepindrak/go-gin-gorm-starter/internal/user"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.Connect(db.Options{DSN: cfg.DSN()})
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate entities here
	if err := db.Migrate(dbConn); err != nil {
		log.Fatal("auto-migrate: ", err)
	}

	r := router.New()

	// USER
	userRepo := user.NewRepository(dbConn)
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc)
	userHandler.Register(r)

	// POST
	postRepo := post.NewRepository(dbConn)
	postSvc := post.NewService(postRepo)
	postHandler := post.NewHandler(postSvc)
	postHandler.Register(r)

	addr := ":" + cfg.AppPort
	log.Println("✅ Server running on", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
