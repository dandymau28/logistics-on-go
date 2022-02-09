package dto

type ShipmentCreateDTO struct {
	ShipmentCode     string `json:"shipment_code" form:"shipment_code" binding:"required"`
	SendersName      string `json:"senders_name" form:"senders_name" binding:"required"`
	SendersPhone     string `json:"senders_phone" form:"senders_phone" binding:"required"`
	SendersAddress   string `json:"senders_address" form:"senders_address" binding:"required"`
	SendersZipcode   string `json:"senders_zipcode" form:"senders_zipcode"`
	ReceiversName    string `json:"receivers_name" form:"receivers_name" binding:"required"`
	ReceiversPhone   string `json:"receivers_phone" form:"receivers_phone" binding:"required"`
	ReceiversAddress string `json:"receivers_address" form:"receivers_address" binding:"required"`
	ReceiversZipcode string `json:"receivers_zipcode" form:"receivers_zipcode"`
	TotalPrice       int64  `json:"total_price" form:"total_price" binding:"required"`
}

type ShipmentDTO struct {
	ShipmentCode     string             `json:"shipment_code" form:"shipment_code" binding:"required"`
	SendersName      string             `json:"senders_name" form:"senders_name" binding:"required"`
	SendersPhone     string             `json:"senders_phone" form:"senders_phone" binding:"required"`
	SendersAddress   string             `json:"senders_address" form:"senders_address" binding:"required"`
	SendersZipcode   string             `json:"senders_zipcode" form:"senders_zipcode"`
	ReceiversName    string             `json:"receivers_name" form:"receivers_name" binding:"required"`
	ReceiversPhone   string             `json:"receivers_phone" form:"receivers_phone" binding:"required"`
	ReceiversAddress string             `json:"receivers_address" form:"receivers_address" binding:"required"`
	ReceiversZipcode string             `json:"receivers_zipcode" form:"receivers_zipcode"`
	TotalPrice       int64              `json:"total_price" form:"total_price" binding:"required"`
	Contain          []ContainCreateDTO `json:"detail"`
}

type ShipmentUpdateDTO struct {
	ID               int64              `json:"id" form:"id"`
	ShipmentCode     string             `json:"shipment_code" form:"shipment_code" binding:"required"`
	SendersName      string             `json:"senders_name" form:"senders_name" binding:"required"`
	SendersPhone     string             `json:"senders_phone" form:"senders_phone" binding:"required"`
	SendersAddress   string             `json:"senders_address" form:"senders_address" binding:"required"`
	SendersZipcode   string             `json:"senders_zipcode" form:"senders_zipcode"`
	ReceiversName    string             `json:"receivers_name" form:"receivers_name" binding:"required"`
	ReceiversPhone   string             `json:"receivers_phone" form:"receivers_phone" binding:"required"`
	ReceiversAddress string             `json:"receivers_address" form:"receivers_address" binding:"required"`
	ReceiversZipcode string             `json:"receivers_zipcode" form:"receivers_zipcode"`
	TotalPrice       int64              `json:"total_price" form:"total_price" binding:"required"`
	Contain          []ContainCreateDTO `json:"detail"`
}
