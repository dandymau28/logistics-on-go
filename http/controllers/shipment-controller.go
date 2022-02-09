package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/uptrace/bun"
	"gitlab.com/logistics-on-go/dto"
	"gitlab.com/logistics-on-go/entity"
	"gitlab.com/logistics-on-go/helper"
)

type ShipmentController interface {
	InsertShipment(ctx *gin.Context)
	GetAllShipment(ctx *gin.Context)
	GetShipmentByID(ctx *gin.Context)
	UpdateShipment(ctx *gin.Context)
	DeleteShipmentByID(ctx *gin.Context)
	SearchShipmentByCode(ctx *gin.Context)
}

type shipmentController struct {
	connection *bun.DB
}

func NewShipmentController(dbConn *bun.DB) ShipmentController {
	return &shipmentController{
		connection: dbConn,
	}
}

func (s *shipmentController) InsertShipment(ctx *gin.Context) {
	tx, errTx := s.connection.BeginTx(ctx, &sql.TxOptions{})

	if errTx != nil {
		res := helper.BuildErrorResponse("Failed to process request", errTx.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var shipmentDTO dto.ShipmentDTO
	errShipment := ctx.ShouldBind(&shipmentDTO)

	if errShipment != nil {
		res := helper.BuildErrorResponse("Failed to process request", errShipment.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var containCreateDTO []dto.ContainCreateDTO = shipmentDTO.Contain

	shipment := entity.Shipment{}
	err := smapping.FillStruct(&shipment, smapping.MapFields(&shipmentDTO))

	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var lastId int64

	_, errDB := tx.NewInsert().Model(&shipment).Returning("id").Exec(ctx, &lastId)

	if errDB != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDB.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	contains := []entity.Contain{}
	contain := entity.Contain{}
	for _, row := range containCreateDTO {
		row.ShipmentID = lastId
		errFill := smapping.FillStruct(&contain, smapping.MapFields(&row))

		if errFill != nil {
			res := helper.BuildErrorResponse("Failed to process request", errFill.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		contains = append(contains, contain)
	}

	result, errDBContain := tx.NewInsert().Model(&contains).Exec(ctx)

	if errDBContain != nil {
		errRb := tx.Rollback()

		if errRb != nil {
			res := helper.BuildErrorResponse("Failed to RB", errRb.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := helper.BuildErrorResponse("Failed to insert contain", errDBContain.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errCom := tx.Commit()

	if errCom != nil {
		res := helper.BuildErrorResponse("Failed to commit", errCom.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, res)
}

func (s *shipmentController) UpdateShipment(ctx *gin.Context) {
	tx, errTx := s.connection.BeginTx(ctx, &sql.TxOptions{})

	if errTx != nil {
		res := helper.BuildErrorResponse("Failed to process request", errTx.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var shipmentUpdateDTO dto.ShipmentUpdateDTO
	err := ctx.ShouldBind(&shipmentUpdateDTO)

	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	shipmentID, errID := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if errID != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	shipmentUpdateDTO.ID = shipmentID

	shipment := entity.Shipment{}
	errMap := smapping.FillStruct(&shipment, smapping.MapFields(&shipmentUpdateDTO))

	if errMap != nil {
		res := helper.BuildErrorResponse("Failed to process request", errMap.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := tx.NewUpdate().Model(&shipment).Where("id = ?", shipmentID).Exec(ctx)

	if err != nil {
		errRB := tx.Rollback()

		if errRB != nil {
			log.Fatalf("Rollback failed: %v\n", errRB.Error())
		}

		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	contains := []entity.Contain{}
	delete, errDel := tx.NewDelete().Model(&contains).Where("shipment_id = ?", shipmentID).Exec(ctx)

	fmt.Printf("delete result: %v\n", delete)

	if errDel != nil {
		errRB := tx.Rollback()

		if errRB != nil {
			log.Fatalf("Rollback failed: %v\n", errRB.Error())
		}

		res := helper.BuildErrorResponse("Failed to process request", errDel.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var containCreateDTO []dto.ContainCreateDTO = shipmentUpdateDTO.Contain

	contain := entity.Contain{}

	for _, row := range containCreateDTO {
		row.ShipmentID = shipmentID
		errFil := smapping.FillStruct(&contain, smapping.MapFields(&row))

		if errFil != nil {
			log.Fatalf("Error fill contain: %v\n", errFil.Error())
		}

		contains = append(contains, contain)
	}

	_, errIns := tx.NewInsert().Model(&contains).Exec(ctx)

	if errIns != nil {
		errRB := tx.Rollback()

		if errRB != nil {
			log.Fatalf("Error rollback: %v\n", errRB.Error())
		}

		res := helper.BuildErrorResponse("Failed to process request", errIns.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errCom := tx.Commit()

	if errCom != nil {
		res := helper.BuildErrorResponse("Failed to commit", errCom.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, res)
}

func (s *shipmentController) GetAllShipment(ctx *gin.Context) {
	var shipment []entity.Shipment
	err := s.connection.NewSelect().Model(&shipment).Relation("Contain").Scan(ctx)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", shipment)
	ctx.JSON(http.StatusOK, res)
}

func (s *shipmentController) GetShipmentByID(ctx *gin.Context) {
	var shipment entity.Shipment

	errQuery := s.connection.NewSelect().Model(&shipment).Relation("Contain").Where("id = ?", ctx.Param("id")).Limit(1).Scan(ctx)
	if errQuery != nil {
		res := helper.BuildErrorResponse("Failed to get query", errQuery.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", shipment)
	ctx.JSON(http.StatusOK, res)
}

func (s *shipmentController) DeleteShipmentByID(ctx *gin.Context) {
	shipment := entity.Shipment{}
	contains := []entity.Contain{}

	_, errDelContains := s.connection.NewDelete().Model(&contains).Where("shipment_id = ?", ctx.Param("id")).Exec(ctx)

	if errDelContains != nil {
		res := helper.BuildErrorResponse("Failed to get query", errDelContains.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, errDelShipment := s.connection.NewDelete().Model(&shipment).Where("id = ?", ctx.Param("id")).Exec(ctx)

	if errDelShipment != nil {
		res := helper.BuildErrorResponse("Failed to get query", errDelShipment.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, res)
}

func (s *shipmentController) SearchShipmentByCode(ctx *gin.Context) {
	var shipment entity.Shipment

	code, err := ctx.GetQuery("code")
	fmt.Printf("query: %v \n", code)

	if !err {
		res := helper.BuildErrorResponse("Failed to get query", "Shipment code required", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errQuery := s.connection.NewSelect().Model(&shipment).Relation("Contain").Where("shipment_code = ?", code).Limit(1).Scan(ctx)
	if errQuery != nil {
		res := helper.BuildErrorResponse("Failed to get query", errQuery.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", shipment)
	ctx.JSON(http.StatusOK, res)
}
