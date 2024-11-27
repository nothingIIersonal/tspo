package repositories

import (
	"context"
	"pr8_1/dtos"
	"pr8_1/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(User *models.User) RepositoryResult {
	err := r.db.Save(User).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: User}
}

func (r *UserRepository) FindAll() RepositoryResult {
	var Users models.Users

	err := r.db.Preload("Role").Find(&Users).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Users}
}

func (r *UserRepository) FindAllWithCtx(ctx *context.Context) RepositoryResult {
	var Users models.Users

	err := r.db.WithContext(*ctx).Preload("Role").Find(&Users).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Users}
}

func (r *UserRepository) FindAllPaging(limit int, offset int, sort string, searchs []dtos.Search) (RepositoryResult, int64) {
	var Users models.Users
	var total int64

	find := r.db.Preload("Role").Limit(limit).Offset(offset).Order(sort)
	err := applySearchs(find, searchs).Find(&Users).Count(&total).Error

	if err != nil {
		return RepositoryResult{Error: err}, 0
	}

	return RepositoryResult{Result: &Users}, total
}

func (r *UserRepository) FindOneById(id uint) RepositoryResult {
	var User models.User

	err := r.db.Preload("Role").Where(&models.User{UserID: id}).Take(&User).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &User}
}

func (r *UserRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&models.User{UserID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *UserRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("UserID IN (?)", *ids).Delete(&models.Users{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
