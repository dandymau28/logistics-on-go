package entity

import "github.com/uptrace/bun"

type Contain struct {
	bun.BaseModel `bun:"table:contain,alias:c"`

	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	ShipmentID  int64     `bun:"shipment_id,," json:"shipment_id"`
	ProductName string    `bun:"product_name,," json:"product_name"`
	Weight      int64     `bun:"weight,," json:"weight"`
	Category    string    `bun:"category,," json:"category"`
	Price       int64     `bun:"price,," json:"price"`
	Shipment    *Shipment `bun:"rel:belongs-to,join:shipment_id=id"`
}
