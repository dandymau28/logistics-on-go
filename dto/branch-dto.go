package dto

type BranchCreateDTO struct {
	BranchCode string `json:"branch_code" form:"branch_code" binding:"required"`
	BranchName string `json:"branch_name" form:"branch_name" binding:"required"`
	Address    string `json:"address" form:"address" binding:"required"`
}

type BranchUpdateDTO struct {
	ID         int64  `json:"id" form:"id"`
	BranchCode string `json:"branch_code" form:"branch_code" binding:"required"`
	BranchName string `json:"branch_name" form:"branch_name" binding:"required"`
	Address    string `json:"address" form:"address" binding:"required"`
}
