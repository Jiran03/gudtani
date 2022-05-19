package warehouseAPI

import (
	"net/http"
	"strconv"

	"github.com/Jiran03/gudtani/warehouse/domain"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type WarehouseHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewWarehouseHandler(service domain.Service) WarehouseHandler {
	return WarehouseHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (wh WarehouseHandler) InsertData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := wh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	warehouseRes, err := wh.service.InsertData(toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	warehouseObj := fromDomain(warehouseRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    warehouseObj,
	})
}

func (wh WarehouseHandler) GetDataByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	warehouseResp, err := wh.service.GetDataByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	warehouseObj := fromDomain(warehouseResp)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    warehouseObj,
	})
}

func (wh WarehouseHandler) UpdateData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	errVal := wh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	warehouseRes, err := wh.service.UpdateData(id, toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	warehouseObj := fromDomain(warehouseRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    warehouseObj,
	})
}

func (wh WarehouseHandler) GetAllData(ctx echo.Context) error {
	warehouseRes, err := wh.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	warehouseObj := []ResponseJSON{}
	for _, value := range warehouseRes {
		warehouseObj = append(warehouseObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    warehouseObj,
	})
}

func (wh WarehouseHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := wh.service.DeleteData(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
	})
}

func (wh WarehouseHandler) GetDataByAddress(ctx echo.Context) error {
	var warehouseObj []ResponseJSON
	address := ctx.QueryParam("address")
	warehouseRes, err := wh.service.GetDataByAddress(address)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	for _, value := range warehouseRes {
		warehouseObj = append(warehouseObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    warehouseObj,
	})
}
