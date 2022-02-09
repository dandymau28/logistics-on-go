package dto

type ContainCreateDTO struct {
	ShipmentID  int64  `json:"shipment_id" form:"shipment_id" binding:"required"`
	ProductName string `json:"product_name" form:"product_name" binding:"required"`
	Weight      int64  `json:"weight" form:"weight" binding:"required"`
	Category    string `json:"category" form:"category"`
	Price       int64  `json:"price" form:"price" binding:"required"`
}

type ContainUpdateDTO struct {
	ID          int64  `json:"id" form:"id"`
	ShipmentID  int64  `json:"shipment_id" form:"shipment_id" binding:"required"`
	ProductName string `json:"product_name" form:"product_name" binding:"required"`
	Weight      int64  `json:"weight" form:"weight" binding:"required"`
	Category    string `json:"category" form:"category"`
	Price       int64  `json:"price" form:"price" binding:"required"`
}
