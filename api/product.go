package api

import (
	"Product-Store/common"
	"Product-Store/interfaces/api"
	"Product-Store/interfaces/service"
	"Product-Store/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type productApi struct {
	ps service.ProductServiceInterface
}

func (p productApi) Create(ctx *gin.Context) {
	formData := model.ProductDto{}
	err := ctx.Bind(&formData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	err = p.ps.Create(formData)
	if err != nil {
		ctx.JSON(http.StatusConflict, common.ErrorResponse{
			Status:  http.StatusConflict,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, common.SuccessResponse{
			Status:  http.StatusOK,
			Message: "product created",
			Data:    nil,
		})
	}
}

func (p productApi) Get(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productApi) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productApi) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productApi) GetById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productApi) GetByFilters(ctx *gin.Context) {
	filter := getFiltersValues(ctx)
	res, total := p.ps.GetByFilter(filter)
	hasMore := false
	if total > filter.Limit {
		hasMore = true
	}
	if res == nil {
		ctx.JSON(http.StatusAccepted, common.SuccessResponse{
			Status:  http.StatusNoContent,
			Message: "no brand found",
			Data:    res,
		})
	} else {
		ctx.JSON(http.StatusAccepted, common.PaginationResponse{
			Status:   http.StatusOK,
			Message:  "operation successful",
			Data:     res,
			Total:    total,
			Page:     int(filter.Page),
			PageSize: int(filter.Limit),
			HasMore:  hasMore,
		})
	}

}

func getFiltersValues(ctx *gin.Context) common.ProductFilter {
	filter := common.ProductFilter{
		Name:     ctx.Query("name"),
		Category: ctx.Query("category"),
		Supplier: ctx.Query("supplier"),
	}

	maxPrice, _ := strconv.ParseFloat(ctx.Query("max_price"), 64)
	filter.MaxPrice = maxPrice

	minPrice, _ := strconv.ParseFloat(ctx.Query("min_price"), 64)
	filter.MinPrice = minPrice

	boolValue, _ := strconv.ParseBool(ctx.Query("verified_supplier"))
	filter.IsVerified = boolValue

	brandsStr := ctx.Query("brands")
	if brandsStr != "" {
		brands := strings.Split(brandsStr, ",")
		filter.Brands = brands
	}

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	if page == "" {
		filter.Page = 0
		filter.Limit = 10

	} else {
		filter.Page, _ = strconv.ParseInt(page, 10, 64)
		filter.Limit, _ = strconv.ParseInt(limit, 10, 64)
	}

	return filter
}

func NewProductApi(serviceInterface service.ProductServiceInterface) api.ProductApi {
	return &productApi{ps: serviceInterface}
}
