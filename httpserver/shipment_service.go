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

type shipmentService struct {
	calculator *shipmentcalc.ShipmentCalculator
}

func NewShipmentService(calculator *shipmentcalc.ShipmentCalculator) *shipmentService {
	return &shipmentService{
		calculator: calculator,
	}
}

func (s *shipmentService) shipmentPacksHandler() http.HandlerFunc {
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

		if len(input.PackSizes) == 0 {
			errorMsg := "pack_sizes must not be empty"
			log.Printf("Request validation error: %s: %+v", errorMsg, input)
			writeJSONError(w, http.StatusBadRequest, errorMsg)
			return
		}
		if input.Amount <= 0 {
			errorMsg := "amount must be greater than 0"
			log.Printf("Request validation error: %s: %+v", errorMsg, input)
			writeJSONError(w, http.StatusBadRequest, "amount must be greater than 0")
			return
		}

		totalPacks := s.calculator.CalculateTotalPacks(input.PackSizes, input.Amount)
		output := &shipmentPacksResponse{
			TotalPacks: totalPacks,
		}

		writeJSON(w, http.StatusOK, output)
	}
}
