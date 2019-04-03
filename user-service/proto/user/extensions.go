package go_micro_srv_user

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// BeforeCreate will run before a creation
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	return scope.SetColumn("Id", uuid.String())
}
