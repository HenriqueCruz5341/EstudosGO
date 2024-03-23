package main

import (
	"api-and-messagerie/internal/infra/akafka"
	"api-and-messagerie/internal/infra/repository"
	"api-and-messagerie/internal/infra/web"
	"api-and-messagerie/internal/usecase"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	chi "github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost:5432/api-and-messagerie?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryPostgreSql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8080", r)
	println("Server running on port 8080")

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "localhost:9092", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Println(err)
		}
		_, err = createProductUseCase.Execute(dto)
		if err != nil {
			log.Println(err)
		}
	}
}
