package main

import (
	"github.com/whacosta/API-Facturacion-Ecuador/src/services"
	"log"
	"net/http"
)

func main() {
	log.Println("Start Service Port :8080")
	http.HandleFunc("/", services.HelloWorld)
	http.HandleFunc("/hello", services.HelloWorld)
	http.HandleFunc("/invoice", services.GenerateInvoice)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
