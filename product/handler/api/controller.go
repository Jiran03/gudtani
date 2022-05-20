package productAPI

import (
	"net/http"
	"strconv"

	"github.com/Jiran03/gudtani/product/domain"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewProductHandler(service domain.Service) ProductHandler {
	return ProductHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (ph ProductHandler) InsertData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := ph.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	productRes, err := ph.service.InsertData(toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	productObj := fromDomain(productRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    productObj,
	})
}

func (ph ProductHandler) GetDataByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	productResp, err := ph.service.GetDataByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	productObj := fromDomain(productResp)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    productObj,
	})
}

func (ph ProductHandler) UpdateData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	errVal := ph.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	productRes, err := ph.service.UpdateData(id, toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	productObj := fromDomain(productRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    productObj,
	})
}

func (ph ProductHandler) GetAllData(ctx echo.Context) error {
	productRes, err := ph.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	productObj := []ResponseJSON{}
	for _, value := range productRes {
		productObj = append(productObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    productObj,
	})
}

func (ph ProductHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := ph.service.DeleteData(id)

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
