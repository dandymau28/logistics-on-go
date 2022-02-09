package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"gitlab.com/logistics-on-go/config"
	"gitlab.com/logistics-on-go/http/controllers"
)

var (
	db                 *bun.DB                        = config.DatabaseConnection()
	shipmentController controllers.ShipmentController = controllers.NewShipmentController(db)
)

func main() {
	r := gin.Default()

	shipmentRoutes := r.Group("api/shipment")
	{
		shipmentRoutes.GET("/", shipmentController.GetAllShipment)
		shipmentRoutes.POST("/", shipmentController.InsertShipment)
		shipmentRoutes.GET("/:id", shipmentController.GetShipmentByID)
		shipmentRoutes.GET("/searchByCode", shipmentController.SearchShipmentByCode)
		shipmentRoutes.PUT("/:id", shipmentController.UpdateShipment)
		shipmentRoutes.DELETE("/:id", shipmentController.DeleteShipmentByID)
	}

	r.Run()
}
