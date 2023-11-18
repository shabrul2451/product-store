package database

import (
	"Product-Store/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func initCategoryData() []model.Category {
	mobileId := uuid.New().String()
	watchId := uuid.New().String()
	computerId := uuid.New().String()
	data := []model.Category{model.Category{
		Id:        mobileId,
		Name:      "mobile",
		ParentId:  "",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Category{
		Id:        "6a233df1-b317-4be3-a2ae-8d02ea5205e3",
		Name:      "ios",
		ParentId:  mobileId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}, {
		Id:        "76d33749-54a6-4ae1-a6a7-182292e1efb4",
		Name:      "android",
		ParentId:  mobileId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}, {
		Id:        watchId,
		Name:      "watch",
		ParentId:  "",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, {
		Id:        uuid.New().String(),
		Name:      "smart watch",
		ParentId:  watchId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}, {
		Id:        uuid.New().String(),
		Name:      "analog watch",
		ParentId:  watchId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}, {
		Id:        uuid.New().String(),
		Name:      "digital watch",
		ParentId:  watchId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}, {
		Id:        computerId,
		Name:      "computer",
		ParentId:  "",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, {
		Id:        uuid.New().String(),
		Name:      "desktop",
		ParentId:  computerId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}, {
		Id:        "005886a0-badf-4c03-bcc2-1aa910db9433",
		Name:      "laptop",
		ParentId:  computerId,
		StatusId:  1,
		CreatedAt: time.Now(),
	}}
	return data
}

func TestCategoryRepository_Update(t *testing.T) {
	data := initCategoryData()
	ncr := NewCategoriesRepository()
	for i, _ := range data {
		err := ncr.Create(data[i])
		if err != nil {
			return
		}
	}
	dataFrmDb := ncr.Get()
	if !reflect.DeepEqual(data, dataFrmDb) {
		assert.Len(t, data, len(dataFrmDb))
	}
}
