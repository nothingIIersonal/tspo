package repositories

import (
	"context"
	"fmt"
	"pr8_1/dtos"
	"pr8_1/models"

	"gorm.io/gorm"
)

type VendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *VendorRepository {
	return &VendorRepository{db: db}
}

func (r *VendorRepository) Save(Vendor *models.Vendor) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Save(Vendor).Error
	fmt.Println("[VENDOR] Trying to save...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[VENDOR] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[VENDOR] Commited!")

	return RepositoryResult{Result: Vendor}
}

func (r *VendorRepository) FindAll() RepositoryResult {
	var Vendors models.Vendors

	err := r.db.Find(&Vendors).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Vendors}
}

func (r *VendorRepository) FindAllWithCtx(ctx *context.Context) RepositoryResult {
	var Vendors models.Vendors

	err := r.db.WithContext(*ctx).Find(&Vendors).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Vendors}
}

func (r *VendorRepository) FindAllPaging(limit int, offset int, sort string, searchs []dtos.Search) (RepositoryResult, int64) {
	var Vendors models.Vendors
	var total int64

	find := r.db.Limit(limit).Offset(offset).Order(sort)
	err := applySearchs(find, searchs).Find(&Vendors).Count(&total).Error

	if err != nil {
		return RepositoryResult{Error: err}, 0
	}

	return RepositoryResult{Result: &Vendors}, total
}

func (r *VendorRepository) FindOneById(id uint) RepositoryResult {
	var Vendor models.Vendor

	err := r.db.Where(&models.Vendor{VendorID: id}).Take(&Vendor).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Vendor}
}

func (r *VendorRepository) DeleteOneById(id uint) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Delete(&models.Vendor{VendorID: id}).Error
	fmt.Println("[VENDOR] Trying to delete...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[VENDOR] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[VENDOR] Commited!")

	return RepositoryResult{Result: nil}
}

func (r *VendorRepository) DeleteByIds(ids *[]string) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Where("VendorID IN (?)", *ids).Delete(&models.Vendors{}).Error
	fmt.Println("[VENDOR] Trying to delete...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[VENDOR] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[VENDOR] Commited!")

	return RepositoryResult{Result: nil}
}
