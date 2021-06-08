package laracom_service_user

import (
	"time"

	uuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	uuid := uuid.New()
	return scope.SetColumn("Id", uuid.String())
}

func (model *User) BeforeSave(scope *gorm.Scope) error {
	_ = scope.SetColumn("UpdatedAt", time.Now().Format(time.RFC3339))
	return nil
}

func (model *PasswordReset) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	return nil
}
