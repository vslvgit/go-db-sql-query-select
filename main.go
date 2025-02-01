package main

import (
	"database/sql"
	. "fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale
	var (
		product string
		volume  int
		date    int
	)
	sales = db.QueryRow("SELECT product, volume, date FROM sales WHERE id = :id", sql.Named("id", client))
	err = row.Scan(&product, &volume, &date)
	if err != nil {
		Println(err)
		return

	}
	return sales, nil
}

func main() {
	client := 208

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	sales, err := selectSales(client)
	if err != nil {
		Println(err)
		return
	}

	for _, sale := range sales {
		Println(sale)
	}
}
