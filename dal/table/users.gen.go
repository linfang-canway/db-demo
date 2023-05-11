// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package table

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	Name      *string        `gorm:"column:name;type:longtext" json:"name"`
	Age       *string        `gorm:"column:age;type:varchar(64)" json:"age"`
	Role      *string        `gorm:"column:role;type:varchar(64)" json:"role"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
