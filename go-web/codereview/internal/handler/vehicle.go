package handler

import (
	"app/internal"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{service: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	service internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (handler *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		vehicles, err := handler.service.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		// fazemos para atender o formato que o usuario espera
		data := make(map[int]VehicleJSON)
		for id, vehicle := range vehicles {
			data[id] = VehicleJSON{
				ID:              vehicle.Id,
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Weight:          vehicle.Weight,
				Height:          vehicle.Height,
				Length:          vehicle.Length,
				Width:           vehicle.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (handler *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// se tivesse um body que eu preciso pegar que o usuario enviou
		// var reqBody VehicleJSON

		// request (pegar dados e valida-los)
		// request.JSON(r, &reqBody)

		color := chi.URLParam(r, "color")

		if color == "" {
			response.Error(w, 400, "url param color is required")
			return
		}

		year := chi.URLParam(r, "year")

		if year == "" {
			response.Error(w, 400, "url param year is required")
			return
		}

		yearConv, err := strconv.Atoi(year)

		if err != nil {
			response.Error(w, 400, "url param year is invalid")
			return
		}

		// process

		vehicles, err := handler.service.FindByColorAndYear(color, yearConv)

		// resposta

		if len(vehicles) == 0 {
			response.Error(w, 404, fmt.Sprintf("vehicles with color %s and year %s not found", color, year))
			return
		}

		// fazemos para atender o formato que o usuario espera
		data := make(map[int]VehicleJSON)
		for id, vehicle := range vehicles {
			data[id] = VehicleJSON{
				ID:              vehicle.Id,
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Weight:          vehicle.Weight,
				Height:          vehicle.Height,
				Length:          vehicle.Length,
				Width:           vehicle.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
