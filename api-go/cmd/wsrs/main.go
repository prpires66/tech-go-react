package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rocketseat-education/semana-tech-go-react-server/internal/api"
	"github.com/rocketseat-education/semana-tech-go-react-server/internal/store/pgstore"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	))
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	handler := api.NewHandler(pgstore.New(pool))

	apiServerHost := os.Getenv("API_SERVER_HOST")
	if apiServerHost == "" {
		apiServerHost = "localhost"
	}

	apiServerPort := os.Getenv("API_SERVER_PORT")
	if apiServerPort == "" {
		apiServerPort = "8080"
	}

	// Combina o host e a porta
	serverAddr := apiServerHost + ":" + apiServerPort

	go func() {
		if err := http.ListenAndServe(serverAddr, handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	fmt.Printf("API is running at http://%s\n", serverAddr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
