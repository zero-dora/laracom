package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"type:varchar(100)"`
	Email         string `gorm:"type:varchar(100);unique_index"`
	Password      string
	Status        uint8 `gorm:"default:1"`
	StripeId      string
	CardBrand     string
	CardLastFour  string
	RememberToken string
}

func (model *User) ToORM(req *pb.User) (*User, error) {
	if req.Id != "" {
		id, _ := strconv.ParseUint(req.Id, 10, 64)
		model.ID = uint(id)
	}
}
