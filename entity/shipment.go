package entity

import "github.com/uptrace/bun"

type Shipment struct {
	bun.BaseModel `bun:"table:shipment,alias:s"`

	ID               int64      `bun:"id,pk,autoincrement" json:"id"`
	ShipmentCode     string     `bun:"shipment_code" json:"shipment_code"`
	SendersName      string     `bun:"senders_name" json:"senders_name"`
	SendersPhone     string     `bun:"senders_phone" json:"senders_phone"`
	SendersAddress   string     `bun:"senders_address" json:"senders_address"`
	SendersZipcode   string     `bun:"senders_zipcode" json:"senders_zipcode"`
	ReceiversName    string     `bun:"receivers_name" json:"receivers_name"`
	ReceiversPhone   string     `bun:"receivers_phone" json:"receivers_phone"`
	ReceiversAddress string     `bun:"receivers_address" json:"receivers_address"`
	ReceiversZipcode string     `bun:"receivers_zipcode" json:"receivers_zipcode"`
	TotalPrice       int64      `bun:"total_price" json:"total_price"`
	Contain          []*Contain `bun:"rel:has-many,join:id=shipment_id"`
}
