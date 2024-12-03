package repositories

import (
	"context"
	"fmt"
	"pr8_1/dtos"
	"pr8_1/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(Order *models.Order) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Save(Order).Error
	fmt.Println("[ORDER] Trying to save...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[ORDER] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[ORDER] Commited!")

	return RepositoryResult{Result: Order}
}

func (r *OrderRepository) FindAll() RepositoryResult {
	var Orders models.Orders

	err := r.db.Find(&Orders).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Orders}
}

func (r *OrderRepository) FindAllWithCtx(ctx *context.Context) RepositoryResult {
	var Orders models.Orders

	err := r.db.WithContext(*ctx).Find(&Orders).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Orders}
}

func (r *OrderRepository) FindAllPaging(limit int, offset int, sort string, searchs []dtos.Search) (RepositoryResult, int64) {
	var Orders models.Orders
	var total int64

	find := r.db.Limit(limit).Offset(offset).Order(sort).Order(sort)
	err := applySearchs(find, searchs).Find(&Orders).Count(&total).Error

	if err != nil {
		return RepositoryResult{Error: err}, 0
	}

	return RepositoryResult{Result: &Orders}, total
}

func (r *OrderRepository) FindOneById(id uint) RepositoryResult {
	var Order models.Order

	err := r.db.Where(&models.Order{OrderID: id}).Take(&Order).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Order}
}

func (r *OrderRepository) DeleteOneById(id uint) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Delete(&models.Order{OrderID: id}).Error
	fmt.Println("[ORDER] Trying to delete...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[ORDER] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[ORDER] Commited!")

	return RepositoryResult{Result: nil}
}

func (r *OrderRepository) DeleteByIds(ids *[]string) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Where("OrderID IN (?)", *ids).Delete(&models.Orders{}).Error
	fmt.Println("[ORDER] Trying to delete...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[ORDER] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[ORDER] Commited!")

	return RepositoryResult{Result: nil}
}
