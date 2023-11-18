package database

import (
	"Product-Store/model"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func initSupplierData() []model.Supplier {
	data := []model.Supplier{model.Supplier{
		Id:                 "ea33b52f-abe9-4f06-b01b-ed72cdf0cee9",
		Name:               "a group",
		Email:              "a@gmail.com",
		Phone:              "01111111111",
		StatusId:           1,
		IsVerifiedSupplier: true,
		CreatedAt:          time.Now(),
	}, {
		Id:                 "241acb7f-9b2b-4e97-96a5-830e0b076d53",
		Name:               "b group",
		Email:              "b@gmail.com",
		Phone:              "02222222222",
		StatusId:           1,
		IsVerifiedSupplier: false,
		CreatedAt:          time.Now(),
	}}
	return data
}

func TestCreate(t *testing.T) {
	data := initSupplierData()
	nsr := NewSupplierRepository()
	for i, _ := range data {
		err := nsr.Create(data[i])
		if err != nil {
			return
		}
	}
	dataFrmDb := nsr.Get()
	if !reflect.DeepEqual(data, dataFrmDb) {
		assert.Len(t, data, len(dataFrmDb))
	}
}
