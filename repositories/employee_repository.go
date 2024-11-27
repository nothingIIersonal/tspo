package repositories

import (
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
	err := r.db.Save(Employee).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

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

func (r *EmployeeRepository) FindOneById(id uint) RepositoryResult {
	var Employee models.Employee

	err := r.db.Preload("User.Role").Where(&models.Employee{UserID: id}).Take(&Employee).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Employee}
}

func (r *EmployeeRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&models.Employee{UserID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *EmployeeRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("UserID IN (?)", *ids).Delete(&models.Employees{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
