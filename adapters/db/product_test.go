package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/assert"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	schema := `CREATE TABLE products (
		id string,
		name string,
		price FLOAT,
		status string
	  );`

	stmt, err := db.Prepare(schema)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products values("abs","Teste",0,"disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close() //postergar o fechamento do banco
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abs")
	assert.Nil(t, err)
	assert.Equal(t, "abs", product.GetID())
	assert.Equal(t, "Teste", product.GetName())
	assert.Equal(t, 0.0, product.GetPrice())
	assert.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close() //postergar o fechamento do banco
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25.0
	product.Status = application.DISABLED

	result, err := productDb.Save(product)
	assert.Nil(t, err)
	assert.Equal(t, product.ID, result.GetID())
	assert.Equal(t, "Product Test", result.GetName())
	assert.Equal(t, 25.0, result.GetPrice())
	assert.Equal(t, application.DISABLED, result.GetStatus())

	product.Status = "enabled"
	product.Price = 50.0

	result, err = productDb.Save(product)
	assert.Nil(t, err)
	assert.Equal(t, product.ID, result.GetID())
	assert.Equal(t, "Product Test", result.GetName())
	assert.Equal(t, 50.0, result.GetPrice())
	assert.Equal(t, application.ENABLED, result.GetStatus())
}
