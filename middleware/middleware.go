package middleware

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id`
	Message string `json:"message"`
}

var db *sql.DB

func Connection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error is happening in database ")
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

func GetStock(w http.ResponseWriter, r *http.Request) {

}

func GetAllStock(w http.ResponseWriter, r *http.Request) {

}

func CreateStock(w http.ResponseWriter, r *http.Request) {

}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {

}
