package main

import (
	"database/sql"
	"fmt"

	"github.com/anchietajunior/go/full-cycle/internal/infra/database"
	"github.com/anchietajunior/go/full-cycle/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1",
		Price: 1000.0,
		Tax:   10.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
