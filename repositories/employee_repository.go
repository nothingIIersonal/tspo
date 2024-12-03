package repositories

import (
	"context"
	"fmt"
	"pr8_1/dtos"
	"pr8_1/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Save(Employee *models.Employee) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Save(Employee).Error
	fmt.Println("[EMPLOYEE] Trying to save...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[EMPLOYEE] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[EMPLOYEE] Commited!")

	return RepositoryResult{Result: Employee}
}

func (r *EmployeeRepository) FindAll() RepositoryResult {
	var Employees models.Employees

	err := r.db.Preload("User.Role").Find(&Employees).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Employees}
}

func (r *EmployeeRepository) FindAllWithCtx(ctx *context.Context) RepositoryResult {
	var Employees models.Employees

	err := r.db.WithContext(*ctx).Preload("User.Role").Find(&Employees).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Employees}
}

func (r *EmployeeRepository) FindAllPaging(limit int, offset int, sort string, searchs []dtos.Search) (RepositoryResult, int64) {
	var Employees models.Employees
	var total int64

	find := r.db.Preload("User.Role").Limit(limit).Offset(offset)
	err := applySearchs(find, searchs).Order(sort).Find(&Employees).Count(&total).Error

	if err != nil {
		return RepositoryResult{Error: err}, 0
	}

	return RepositoryResult{Result: &Employees}, total
}

func (r *EmployeeRepository) FindOneById(id uint) RepositoryResult {
	var Employee models.Employee

	err := r.db.Preload("User.Role").Where(&models.Employee{UserID: id}).Take(&Employee).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Employee}
}

func (r *EmployeeRepository) DeleteOneById(id uint) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Delete(&models.Employee{UserID: id}).Error
	fmt.Println("[EMPLOYEE] Trying to delete...")

	if err != nil {

		tx.Rollback()
		fmt.Println("[EMPLOYEE] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[EMPLOYEE] Commited!")

	return RepositoryResult{Result: nil}
}

func (r *EmployeeRepository) DeleteByIds(ids *[]string) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Where("UserID IN (?)", *ids).Delete(&models.Employees{}).Error
	fmt.Println("[EMPLOYEE] Trying to delete...")

	if err != nil {

		tx.Rollback()
		fmt.Println("[EMPLOYEE] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[EMPLOYEE] Commited!")

	return RepositoryResult{Result: nil}
}
