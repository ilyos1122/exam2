package main

import (
	"init/config"
	"init/controller"
	"init/storage/postgres"
	"log"
	"net/http"
)

func main() {

	var cfg = config.Load()

	pgStorage, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	handler := controller.NewController(&cfg, pgStorage)

	http.HandleFunc("/category", handler.Category)
	http.HandleFunc("/product", handler.Product)
	http.HandleFunc("/client", handler.Client)
	http.HandleFunc("/branch", handler.Branch)
	http.HandleFunc("/order", handler.Order)
	http.HandleFunc("/orderproduct", handler.OrderProduct)



	log.Println("Listening:", cfg.ServiceHost+cfg.ServiceHTTPPort, "...")
	if err := http.ListenAndServe(cfg.ServiceHost+cfg.ServiceHTTPPort, nil); err != nil {
		panic("Listent and service panic:" + err.Error())
	}
}
