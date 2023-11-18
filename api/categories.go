package api

import (
	"Product-Store/common"
	"Product-Store/interfaces/api"
	"Product-Store/interfaces/service"
	"Product-Store/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

type categoriesApi struct {
	cs service.CategoriesServiceInterface
}

type TreeNode struct {
	model.Category
	Children []*TreeNode `json:"children"`
}

func (c categoriesApi) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c categoriesApi) Get(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c categoriesApi) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c categoriesApi) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c categoriesApi) GetById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c categoriesApi) GetTree(ctx *gin.Context) {
	cat := c.cs.Get()
	tree := BuildCategoryTree(cat)
	ctx.JSON(http.StatusOK, common.SuccessResponse{
		Status:  http.StatusOK,
		Message: "operation successful",
		Data:    tree,
	})
}

func BuildCategoryTree(categories []model.Category) []*TreeNode {
	categoryMap := make(map[string]*TreeNode)
	var rootCategories []*TreeNode

	for _, category := range categories {
		copyCategory := category
		node := &TreeNode{Category: copyCategory, Children: nil}
		categoryMap[copyCategory.Id] = node

		if copyCategory.ParentId == "" {
			rootCategories = append(rootCategories, node)
		} else {
			parent, exists := categoryMap[copyCategory.ParentId]
			if exists {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	sort.SliceStable(rootCategories, func(i, j int) bool {
		return rootCategories[i].Category.Sequence < rootCategories[j].Category.Sequence
	})

	for _, node := range categoryMap {
		sort.SliceStable(node.Children, func(i, j int) bool {
			return node.Children[i].Category.Sequence < node.Children[j].Category.Sequence
		})
	}

	return rootCategories
}

func NewCategoriesApi(serviceInterface service.CategoriesServiceInterface) api.CategoriesApi {
	return &categoriesApi{
		cs: serviceInterface,
	}
}
