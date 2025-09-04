package main

import (
	"log"

	"github.com/asepindrak/go-gin-gorm-starter/internal/config"
	"github.com/asepindrak/go-gin-gorm-starter/internal/db"
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
	if err := dbConn.AutoMigrate(&user.User{}); err != nil {
		log.Fatal("auto-migrate: ", err)
	}

	r := router.New()

	// Wire DI
	repo := user.NewRepository(dbConn)
	svc := user.NewService(repo)
	h := user.NewHandler(svc)
	h.Register(r)

	addr := ":" + cfg.AppPort
	log.Println("✅ Server running on", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
