package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sayandas-sd/stocksApiServer/model"
)

type response struct {
	ID      int64  `json:"id`
	Message string `json:"message"`
}

var db *sql.DB

func connection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error is happening in database")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock model.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("unable to decode from body %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID:      insertID,
		Message: "stock craeted successfully",
	}

	json.NewEncoder(w).Encode(res)

}

func GetStock(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to parse the string %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("Unable to get stock")
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()

	if err != nil {
		log.Fatalf("Unable to get all stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)

}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println("error while parsing")
	}

	var stocks model.Stock

	err = json.NewDecoder(r.Body).Decode(&stocks)

	if err != nil {
		log.Fatalf("Unable to decode body data %v", err)
	}

	updateRows := updateStock(int64(id), stocks)

	msg := fmt.Sprint("Stoks update succesfully %v", updateRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("error while parsing")
	}

	deleteRows := deleteStock(int64(id))

	msg := fmt.Sprint("Stocks deleted %v", deleteRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}

func insertStock(stock model.Stock) int64 {

	db := connection()

	defer db.Close()

	sqlQuery := `INSERT INTO stocks(name, price, compnay) VALUES($1, $2, $3) RETURNING stockid`

	var id int64

	err := db.QueryRow(sqlQuery, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func getStock(id int64) (model.Stock, error) {
	db := connection()

	defer db.Close()

	var stock model.Stock

	sqlQuery := `SELECT FROM stocks WHERE stockid=$1`

	row := db.QueryRow(sqlQuery, id)

	err := row.Scan(&stock.StockId, &stock.Name, &stock.Company, &stock.Price)

	switch err {

	case sql.ErrNoRows:
		fmt.Println("rows returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to find the rows %v", err)

	}

	return stock, err
}

func getAllStocks() ([]model.Stock, error) {
	db := connection()

	defer db.Close()

	var stocks []model.Stock

	sqlQuery := `SELECT * FROM stocks`

	row, err := db.Query(sqlQuery)

	if err != nil {
		log.Fatalf("Unbale to find rows %v", err)
	}

	defer row.Close()

	for row.Next() {
		var stock model.Stock

		err = row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the rows %v", err)
		}

		stocks = append(stocks, stock)
	}

	return stocks, err
}

func updateStock(id int64, stock model.Stock) int64 {
	db := connection()

	defer db.Close()

	sqlQuery := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	res, err := db.Exec(sqlQuery, stock.Name, stock.Price, stock.Company, stock.StockId)

	if err != nil {
		log.Fatalf("unable to execute the query %v", err)
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("an error occured %v", err)
	}

	fmt.Printf("Total rows are: %v", rows)

	return rows

}

func deleteStock(id int64) int64 {
	db := connection()

	defer db.Close()

	sqlQuery := `DELETE FROM stocks WHERE stockid=$1`

	res, err := db.Exec(sqlQuery, id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("an error occured %v", err)
	}

	fmt.Printf("Total rows effected %v", rows)

	return rows

}
