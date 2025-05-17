package httpserver

import (
	"net/http"
)

func NewHandler(shipmentService *shipmentService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/shipment-packs", shipmentService.shipmentPacksHandler())

	return mux
}
