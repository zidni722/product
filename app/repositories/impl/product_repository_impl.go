package impl

import (
	"github.com/zidni722/pawoon-product/app/models"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryImpl struct {}

func NewProductRepositoryImpl() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) FindAll(db *gorm.DB, entities interface{}) error {
	return db.Find(entities.(*[]models.Product)).Error
}

func (r *ProductRepositoryImpl) FindById(db *gorm.DB, entity interface{}, id int) error {
	return db.First(entity.(*models.Product), id).Error
}

func (r *ProductRepositoryImpl) Create(db *gorm.DB, entity interface{}) error {
	return db.Create(entity.(*models.Product)).Error
}

func (r *ProductRepositoryImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return db.NewRecord(entity.(models.Product))
}

func (r *ProductRepositoryImpl) Update(db *gorm.DB, entity interface{}) error {
	return db.Save(entity.(*models.Product)).Error
}

func (r *ProductRepositoryImpl) Delete(db *gorm.DB, entity interface{}) error {
	return db.Delete(entity.(*models.Product)).Error
}