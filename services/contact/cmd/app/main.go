package main

import (
	"fmt"
	"net/http"
	repository "xamss/microservices/test/pkg/store/postgres"
	"xamss/microservices/test/services/contact/internal/delivery"
)

func main() {
	db, err := repository.ConnPostgres("localhost", "guest", 5432, "pa55word", "mstest")
	if err != nil {
		fmt.Errorf("database connection failed: %s", err)
	}
	defer db.Pool.Close()

	repo := repository.New(db.Pool)
	delivery := delivery.New()
	usecase := usecase.New(repo)

	_ = usecase

	fmt.Println("application started")

	http.ListenAndServe("localhost:4000", delivery.Mux)
}
