package router

import (
	"github.com/gorilla/mux"
	"github.com/sayandas-sd/stocksApiServer/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("Get", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStock).Methods("Get", "OPTIONS")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/newstock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/newstock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
}
