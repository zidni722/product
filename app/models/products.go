package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	UUID                 string
	Name                 string
	Sku                  string
	Price                int64
	Barcode              string
	Desc                 string
	Img                  string
	Tax                  string
	StockUnit            string `gorm:"column:stock_unit"`
	Type                 string `gorm:"type: enum('single','composite','package'); default: 'single'; not null"`
	EquateStatusAndPrice int8   `gorm:"column:equate_status_and_price"`
	IsAllLocationStock   int8   `gorm:"column:is_all_location_stock"`
	IsAllLocationPrice   int8   `gorm:"column:is_all_location_price"`
	IsSellable           int8   `gorm:"column:is_selleble"`
	IsStockTracked       int8   `gorm:"column:is_stock_tracked"`
	HasAlertstock        int8   `gorm:"column:has_alert_stock"`
	AlertStockLimit      uint32 `gorm:"column:alert_stock_limit"`
	HasModifier          int8   `gorm:"column:has_modifier"`
	HasVariant           int8   `gorm:"column:has_variant"`
	UseOutletTax         int8   `gorm:"column:use_outlet_tax"`
	Brand                string
	Weight               int32
	Height               int32
	Long                 int32
	Wide                 int32
	CompanyID            uint `gorm:"column:company_id"`
	ProductCategoryID    uint `gorm:"column:product_category_id"`
	ParentID             uint
	V1ProductID          uint `gorm:"column:v1_product_id"`
	Show                 int8
	PpobID               uint `gorm:"column:ppob_id"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time `sql:"index"`
}
