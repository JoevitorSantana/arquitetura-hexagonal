package main

import (
	"database/sql"

	db2 "github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDb := db2.NewProductDb(db)
	productService := application.NewProductService(productDb)
	product, _ := productService.Create("Product 1", 10)

	productService.Enable(product)
}
