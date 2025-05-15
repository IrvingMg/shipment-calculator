package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"shipment-calculator/shipmentcalc"
)

type errorResponse struct {
	Error string `json:"error"`
}

type shipmentPacksRequest struct {
	PackSizes []int `json:"pack_sizes"`
	Amount    int   `json:"amount"`
}

type shipmentPacksResponse struct {
	TotalPacks map[int]int `json:"total_packs"`
}

func NewHandler(shipmentCalc *shipmentcalc.ShipmentCalculator) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/shipment-packs", shipmentPacksHandler(shipmentCalc))

	return mux
}

// We could add an abstraction layer (e.g., a shipmentService struct) over ShipmentCalculator
// if additional functionality is needed.
func shipmentPacksHandler(shipmentCalc *shipmentcalc.ShipmentCalculator) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			log.Printf("Method not allowed: %s %s", req.Method, req.URL.Path)
			writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		input := &shipmentPacksRequest{}
		err := json.NewDecoder(req.Body).Decode(input)
		if err != nil {
			log.Printf("JSON decode error: %v", err)
			writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		totalPacks := shipmentCalc.CalculateTotalPacks(input.PackSizes, input.Amount)
		output := &shipmentPacksResponse{
			TotalPacks: totalPacks,
		}

		writeJSON(w, http.StatusOK, output)
	}
}
