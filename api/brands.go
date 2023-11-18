package api

import (
	"Product-Store/common"
	"Product-Store/interfaces/api"
	"Product-Store/interfaces/service"
	"Product-Store/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type brandApi struct {
	bs service.BrandsServiceInterface
}

func (b brandApi) Create(ctx *gin.Context) {
	formData := model.Brand{}
	err := ctx.Bind(&formData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	formData.Id = uuid.New().String()
	formData.StatusId = 1
	err = b.bs.Update(formData)
	if err != nil {
		ctx.JSON(http.StatusConflict, common.ErrorResponse{
			Status:  http.StatusConflict,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, common.SuccessResponse{
			Status:  http.StatusOK,
			Message: "brand created",
			Data:    nil,
		})
	}

}

func (b brandApi) Get(ctx *gin.Context) {
	res := b.bs.Get()
	if res == nil {
		ctx.JSON(http.StatusAccepted, common.SuccessResponse{
			Status:  http.StatusNoContent,
			Message: "no brand found",
			Data:    res,
		})
	} else {
		ctx.JSON(http.StatusAccepted, common.SuccessResponse{
			Status:  http.StatusOK,
			Message: "operation successful",
			Data:    res,
		})
	}

}

func (b brandApi) Update(ctx *gin.Context) {
	formData := model.Brand{}
	err := ctx.Bind(&formData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	id := ctx.Param("id")
	formData.Id = id
	err = b.bs.Update(formData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, common.SuccessResponse{
			Status:  http.StatusOK,
			Message: "updated successful",
			Data:    nil,
		})
	}
}

func (b brandApi) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := b.bs.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, common.ErrorResponse{
			Status:  http.StatusConflict,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, common.SuccessResponse{
			Status:  http.StatusOK,
			Message: "deleted successful",
			Data:    nil,
		})
	}

}

func (b brandApi) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	res := b.bs.GetById(id)
	if res.Id == "" {
		ctx.JSON(http.StatusNoContent, common.SuccessResponse{
			Status:  http.StatusNoContent,
			Message: "no brand found",
			Data:    res,
		})
	} else {
		ctx.JSON(http.StatusOK, common.SuccessResponse{
			Status:  http.StatusOK,
			Message: "operation successful",
			Data:    res,
		})
	}
}

func NewBrandApi(serviceInterface service.BrandsServiceInterface) api.BrandApi {
	return &brandApi{bs: serviceInterface}
}
