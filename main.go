package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"shipment-calculator/httpserver"
	"shipment-calculator/shipmentcalc"
)

func main() {
	address := os.Getenv("HTTP_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	shipmentCalculator := shipmentcalc.New()
	shipmentService := httpserver.NewShipmentService(shipmentCalculator)
	handler := httpserver.NewHandler(shipmentService)
	server := httpserver.New(address, handler)
	go server.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	server.Stop()

	log.Println("Server exited cleanly")
}
