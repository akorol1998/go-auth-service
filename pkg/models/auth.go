package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:"type:varchar(64)" json:"name"`
	Email    string `gorm:"type:varchar(64);unique" json:"email"`
	Password string `gorm:"type:varchar(256)"`
	// Using pointer means we can parse null values to the field when working with model
	Birthday   *string
	Archived   bool
	CreditCard *uint
	Orders     []Order
	Roles      []Role `gorm:"many2many:user_roles;"`
}

type Permission struct {
	ID          uint    `gorm:"primaryKey;"`
	Name        string  `gorm:"type:varchar(64);unique;not null;index;"`
	Description *string `gorm:"type:varchar(128)"`
	Roles       []*Role `gorm:"many2many:role_permissions;"`
}

type Role struct {
	ID          uint          `gorm:"primaryKey;"`
	Name        string        `gorm:"type:varchar(64);unique;not null;index;"`
	Description *string       `gorm:"type:varchar(128);"`
	Users       []User        `gorm:"many2many:user_roles;"`
	Permissions []*Permission `gorm:"many2many:role_permissions;"`
}

type UserRoles struct {
	UserID    uint `gorm:"primaryKey;"`
	RoleID    uint `gorm:"primaryKey;"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type OrderStatus struct {
	Name string
}

type Order struct {
	Name   string
	Status OrderStatus `gorm:"embedded;embeddedPrefix:order_"`
	UserID uint
	gorm.Model
}
