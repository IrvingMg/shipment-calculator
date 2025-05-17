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
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./shipmentcalc <address>")
	}
	address := os.Args[1]

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
