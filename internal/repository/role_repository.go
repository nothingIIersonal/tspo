package repository

import (
	"tspo_final/internal/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Save(Role *models.Role) RepositoryResult {
	err := r.db.Save(Role).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: Role}
}

func (r *RoleRepository) FindAll() RepositoryResult {
	var Roles models.Roles

	err := r.db.Find(&Roles).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Roles}
}

func (r *RoleRepository) FindOneById(id uint) RepositoryResult {
	var Role models.Role

	err := r.db.Where(&models.Role{RoleID: id}).Take(&Role).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Role}
}

func (r *RoleRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&models.Role{RoleID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *RoleRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("RoleID IN (?)", *ids).Delete(&models.Roles{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
