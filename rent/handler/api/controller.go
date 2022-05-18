package rentAPI

import (
	"net/http"
	"strconv"

	"github.com/Jiran03/gudhani/rent/domain"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RentHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewRentHandler(service domain.Service) RentHandler {
	return RentHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (rh RentHandler) InsertData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := rh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	rentRes, err := rh.service.InsertData(toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	rentObj := fromDomain(rentRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    rentObj,
	})
}

func (rh RentHandler) GetDataByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	rentResp, err := rh.service.GetDataByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	rentObj := fromDomain(rentResp)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    rentObj,
	})
}

func (rh RentHandler) UpdateData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))

	errVal := rh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	rentRes, err := rh.service.UpdateData(id, toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	rentObj := fromDomain(rentRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    rentObj,
	})
}

func (rh RentHandler) GetAllData(ctx echo.Context) error {
	rentRes, err := rh.service.GetAllData()

	if err != nil {
		return err
	}

	rentObj := []ResponseJSON{}
	for _, value := range rentRes {
		rentObj = append(rentObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    rentObj,
	})
}

func (rh RentHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := rh.service.DeleteData(id)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
	})
}
