package entity

import "github.com/uptrace/bun"

type Branch struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID         int64  `bun:"id,pk,autoincrement"`
	BranchCode string `bun:"branch_code,,"`
	BranchName string `bun:"branch_name,,"`
	Address    string `bun:"address,,"`
}
