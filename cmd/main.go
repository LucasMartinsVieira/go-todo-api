package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/LucasMartinsVieira/go-todo-api/internal/config"
	"github.com/LucasMartinsVieira/go-todo-api/internal/database"
	db "github.com/LucasMartinsVieira/go-todo-api/internal/database/repository"
	"github.com/LucasMartinsVieira/go-todo-api/internal/todo"
)

func main() {
	cfg := config.LoadConfig()

	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "dev"
	}

	config.InitLogger(env)

	pool := database.ConnectDatabase(cfg)
	queries := db.New(pool)

	repo := todo.NewRepository(queries)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service)

	r := gin.Default()
	handler.RegisterRoutes(r)
	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	config.Logger.Info().Str("addr", addr).Msg("🚀 Server running")
	r.Run(addr)
}
