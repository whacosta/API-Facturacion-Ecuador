package main

import (
	"github.com/whacosta/API-Facturacion-Ecuador/src/services"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", services.HelloWorld)
	http.HandleFunc("/invoice", services.GenerateInvoice)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
