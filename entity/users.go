package entity

import "github.com/uptrace/bun"

type Users struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID         int64   `bun:"id,pk,autoincrement"`
	BranchCode string  `bun:"branch_code,,"`
	Nik        string  `bun:"nik,,"`
	Name       string  `bun:"name,,"`
	Role       string  `bun:"role,,"`
	Branch     *Branch `bun:"rel:belongs-to,join:branch_code=branch_code"`
}
