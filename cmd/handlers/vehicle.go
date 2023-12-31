package handlers

import (
	"app/internal/domain"
	"app/internal/vehicle/service"
	"errors"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NewControllerVehicle returns a new instance of a vehicle controller.
func NewControllerVehicle(st service.ServiceVehicle) *ControllerVehicle {
	return &ControllerVehicle{st: st}
}

// ControllerVehicle is an struct that represents a vehicle controller.
type ControllerVehicle struct {
	// StorageVehicle is the storage of vehicles.
	st service.ServiceVehicle
}

// GetAll returns all vehicles.
type VehicleHandlerGetAll struct {
	Id           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}
type ResponseBodyGetAll struct {
	Message string                  `json:"message"`
	Data    []*VehicleHandlerGetAll `json:"vehicles"`
	Error   bool                    `json:"error"`
}

// Save a list of vehicles
type VehicleHandlerSaveVehicles struct {
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}

type ResponseBodySaveVehicles struct {
	Message string
	Error   bool
}

type VehicleHandlerGetByColorAndYear struct {
	Id           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}

type ResponseBodyGetByColorAndYear struct {
	Message string                             `json:"message"`
	Data    []*VehicleHandlerGetByColorAndYear `json:"vehicles"`
	Error   bool                               `json:"error"`
}

func (c *ControllerVehicle) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...

		// process
		vehicles, err := c.st.GetAll()
		if err != nil {
			var code int
			var body ResponseBodyGetAll // globales
			switch {                    // validacion errors type
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				code = http.StatusNotFound
				body = ResponseBodyGetAll{Message: "Not found", Error: true}
			default:
				code = http.StatusInternalServerError
				body = ResponseBodyGetAll{Message: "Internal server error", Error: true}
			}

			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyGetAll{Message: "Success", Data: make([]*VehicleHandlerGetAll, 0, len(vehicles)), Error: false}
		for _, vehicle := range vehicles {
			body.Data = append(body.Data, &VehicleHandlerGetAll{
				Id:           vehicle.Id,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			})
		}

		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) SaveVehicles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody []VehicleHandlerSaveVehicles
		err := ctx.BindJSON(&reqBody)
		if err != nil {
			code := http.StatusBadRequest
			body := ResponseBodySaveVehicles{Message: "Datos de algún vehículo mal formados o incompletos.", Error: true}
			ctx.JSON(code, body)
			return
		}
		// casting to domain.Vehicle
		vl, err := c.st.GetAll()
		if err != nil {
			var code int
			var body ResponseBodySaveVehicles // globales
			switch {                          // validacion errors type
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				code = http.StatusNotFound
				body = ResponseBodySaveVehicles{Message: "Vehicle not found", Error: true}
			default:
				code = http.StatusInternalServerError
				body = ResponseBodySaveVehicles{Message: "Internal server error: retrieving list", Error: true}
			}
			ctx.JSON(code, body)
			return
		}
		id := len(vl) + 1 // autoincrement

		var vehicleList []domain.Vehicle
		for _, value := range reqBody {
			vehicleList = append(vehicleList, domain.Vehicle{
				Id: id,
				Attributes: domain.VehicleAttributes{
					Brand:        value.Brand,
					Model:        value.Model,
					Registration: value.Registration,
					Year:         value.Year,
					Color:        value.Color,
					MaxSpeed:     value.MaxSpeed,
					FuelType:     value.FuelType,
					Transmission: value.Transmission,
					Passengers:   value.Passengers,
					Height:       value.Height,
					Width:        value.Width,
					Weight:       value.Weight,
				},
			})
			id++
		}
		err = c.st.SaveVehicles(vehicleList)
		if err != nil {
			var code int
			var body ResponseBodySaveVehicles // globales
			switch {                          // validacion errors type
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				code = http.StatusNotFound
				body = ResponseBodySaveVehicles{Message: "Vehicle not found", Error: true}
			case errors.Is(err, service.ErrServiceVehicleAlreadyExist):
				code = http.StatusConflict
				body = ResponseBodySaveVehicles{Message: "Algún vehículo tiene un identificador ya existente.", Error: true}
			default:
				code = http.StatusInternalServerError
				body = ResponseBodySaveVehicles{Message: "Internal server error: saving vehicles", Error: true}
			}
			ctx.JSON(code, body)
			return
		}
		code := http.StatusCreated
		body := ResponseBodySaveVehicles{Message: "Vehículos creados exitosamente.", Error: false}
		ctx.JSON(code, body)
	}

}

func (c *ControllerVehicle) GetByColorAndYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...
		color := ctx.Param("color")
		year, err := strconv.Atoi(ctx.Param("year"))
		if err != nil {
			code := http.StatusNotFound
			message := "error parsing year value to int"
			ctx.JSON(code, message)
			return
		}
		// process
		vehicles, err := c.st.GetByColorAndYear(color, year)
		if err != nil {
			var code int
			var body ResponseBodyGetAll // globales
			switch {                    // validacion errors type
			case errors.Is(err, service.ErrServiceVehiclesNotFound):
				code = http.StatusNotFound
				body = ResponseBodyGetAll{Message: "Vehicle not found", Error: true}
			default:
				code = http.StatusInternalServerError
				body = ResponseBodyGetAll{Message: "Internal server error", Error: true}
			}

			ctx.JSON(code, body)
			return

		}

		// response
		code := http.StatusOK
		body := ResponseBodyGetByColorAndYear{Message: "Success", Data: make([]*VehicleHandlerGetByColorAndYear, 0, len(vehicles)), Error: false}
		for _, vehicle := range vehicles {
			body.Data = append(body.Data, &VehicleHandlerGetByColorAndYear{
				Id:           vehicle.Id,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			})
		}

		ctx.JSON(code, body)
	}
}
