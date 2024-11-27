package models

import "time"

type Role struct {
	RoleID uint   `gorm:"primaryKey;autoIncrement;not null"`
	Role   string `gorm:"type:varchar(10);not null"`
}
type Roles []Role

func (Role) TableName() string {
	return "roles"
}

type Feature struct {
	FeatureID uint   `gorm:"primaryKey;autoIncrement;not null"`
	Feature   string `gorm:"type:varchar(32);not null"`
	IsDeleted bool   `gorm:"type:bool;not null"`
}
type Features []Feature

func (Feature) TableName() string {
	return "features"
}

type User struct {
	UserID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name         string `gorm:"type:varchar(32);not null"`
	Phone        string `gorm:"type:varchar(32);not null"`
	Address      string `gorm:"type:varchar(128);not null"`
	Email        string `gorm:"type:varchar(32);not null"`
	PasswordHash string `gorm:"type:varchar(128);not null"`
	IsDeleted    bool   `gorm:"type:bool;not null"`
	Role         []Role `gorm:"many2many:users_roles;joinForeignKey:user_id;joinReferences:role_id"`
}
type Users []User

func (User) TableName() string {
	return "users"
}

type Employee struct {
	UserID   uint    `gorm:"primaryKey;not null"`
	Salary   float64 `gorm:"type:float8;not null"`
	Position string  `gorm:"type:varchar(32);not null"`
	KPI      int16   `gorm:"type:int2;not null"`
	User     User
}
type Employees []Employee

func (Employee) TableName() string {
	return "employees"
}

type Vendor struct {
	VendorID  uint   `gorm:"primaryKey;autoIncrement;not null"`
	Phone     string `gorm:"type:varchar(32);not null"`
	OrgName   string `gorm:"type:varchar(32);not null"`
	INN       string `gorm:"type:varchar(10);not null"`
	OGRN      string `gorm:"type:varchar(13);not null"`
	Address   string `gorm:"type:varchar(128);not null"`
	IsDeleted bool   `gorm:"type:bool;not null"`
}
type Vendors []Vendor

func (Vendor) TableName() string {
	return "vendors"
}

type Good struct {
	GoodID      uint      `gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `gorm:"type:varchar(32);not null"`
	Description string    `gorm:"type:varchar(128);not null"`
	Price       float64   `gorm:"type:float8;not null"`
	Count       int       `gorm:"type:int4;not null"`
	IsDeleted   bool      `gorm:"type:bool;not null"`
	Feature     []Feature `gorm:"many2many:goods_features;joinForeignKey:good_id;joinReferences:feature_id"`
	Vendor      []Vendor  `gorm:"many2many:goods_vendors;joinForeignKey:good_id;joinReferences:vendor_id"`
}
type Goods []Good

func (Good) TableName() string {
	return "goods"
}

type Order struct {
	OrderID      uint      `gorm:"primaryKey;autoIncrement;not null"`
	DeliveryType string    `gorm:"type:varchar(32);not null"`
	DeliveryTime time.Time `gorm:"type:timestamp;not null"`
	OrderTime    time.Time `gorm:"type:timestamp;not null"`
	TotalPrice   float64   `gorm:"type:float8;not null"`
	Canceled     bool      `gorm:"type:bool;not null"`
}
type Orders []Order

func (Order) TableName() string {
	return "orders"
}

type Favorite struct {
	UserID uint `gorm:"primaryKey;not null"`
	GoodID uint `gorm:"primaryKey;not null"`
}

func (Favorite) TableName() string {
	return "favorites"
}

type Basket struct {
	UserID uint `gorm:"primaryKey;not null"`
	GoodID uint `gorm:"primaryKey;not null"`
	Count  int  `gorm:"type:int4;not null"`
}

func (Basket) TableName() string {
	return "baskets"
}

type UserOrder struct {
	UserID  uint `gorm:"primaryKey;not null"`
	OrderID uint `gorm:"primaryKey;not null"`
}

func (UserOrder) TableName() string {
	return "users_orders"
}

type UserRole struct {
	UserID uint `gorm:"primaryKey;not null"`
	RoleID uint `gorm:"primaryKey;not null"`
}

func (UserRole) TableName() string {
	return "users_roles"
}

type GoodFeature struct {
	GoodID    uint `gorm:"primaryKey;not null"`
	FeatureID uint `gorm:"primaryKey;not null"`
}

func (GoodFeature) TableName() string {
	return "goods_features"
}

type GoodVendor struct {
	GoodID   uint `gorm:"primaryKey;not null"`
	VendorID uint `gorm:"primaryKey;not null"`
}

func (GoodVendor) TableName() string {
	return "goods_vendors"
}

type OrderGood struct {
	OrderID uint `gorm:"primaryKey;not null"`
	GoodID  uint `gorm:"primaryKey;not null"`
	Count   int  `gorm:"type:int4;not null"`
}

func (OrderGood) TableName() string {
	return "orders_goods"
}
