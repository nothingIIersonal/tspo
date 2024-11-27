package repositories

import (
	"context"
	"pr8_1/dtos"
	"pr8_1/models"

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

func (r *RoleRepository) FindAllPaging(limit int, offset int, sort string, searchs []dtos.Search) (RepositoryResult, int64) {
	var Roles models.Roles
	var total int64

	find := r.db.Limit(limit).Offset(offset).Order(sort)
	err := applySearchs(find, searchs).Find(&Roles).Count(&total).Error

	if err != nil {
		return RepositoryResult{Error: err}, 0
	}

	return RepositoryResult{Result: &Roles}, total
}

func (r *RoleRepository) FindAllWithCtx(ctx *context.Context) RepositoryResult {
	var Roles models.Roles

	err := r.db.WithContext(*ctx).Find(&Roles).Error

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
