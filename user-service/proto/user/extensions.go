package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BeforeCreate will run before a creation
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, _ := uuid.NewV4()

	return scope.SetColumn("Id", uuid.String())
}
