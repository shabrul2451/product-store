package database

import (
	"Product-Store/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func initBrandData() []model.Brand {
	data := []model.Brand{model.Brand{
		Id:        "b8341371-4d70-42a0-a6e8-bfed2e0e5cc5",
		Name:      "apple",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Brand{
		Id:        "c2f86c20-4740-45d3-a2e7-73d12d3d33c7",
		Name:      "samsung",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Brand{
		Id:        "0becbcba-eaa5-41e6-8fcf-e5a0dc6ee471",
		Name:      "oneplus",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Brand{
		Id:        "a3325b39-aaf7-45cf-b59a-82dc8fe28234",
		Name:      "oppo",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Brand{
		Id:        uuid.New().String(),
		Name:      "vivo",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Brand{
		Id:        uuid.New().String(),
		Name:      "xiaomi",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Brand{
		Id:        uuid.New().String(),
		Name:      "titan",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Brand{
		Id:        uuid.New().String(),
		Name:      "casio",
		StatusId:  1,
		CreatedAt: time.Now(),
	}, model.Brand{
		Id:        uuid.New().String(),
		Name:      "rolex",
		StatusId:  1,
		CreatedAt: time.Now(),
	}}
	return data
}
func TestUpdate(t *testing.T) {
	data := initBrandData()
	nbr := NewBrandRepository()
	for _, each := range data {
		err := nbr.Update(each)
		if err != nil {
			return
		}
	}
	dataFrmDb := nbr.Get()
	if !reflect.DeepEqual(data, dataFrmDb) {
		assert.Len(t, data, len(dataFrmDb))
	}
}

/*func TestDelete(t *testing.T) {
	nbr := NewBrandRepository()
	nbr.Delete("b8341371-4d70-42a0-a6e8-bfed2e0e5cc5")
	dataFrmDb := nbr.Get()
	assert.Len(t, dataFrmDb, 8)
}
*/
