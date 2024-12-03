package repositories

import (
	"context"
	"fmt"
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
	tx := r.db.Begin()

	err := tx.Save(User).Error
	fmt.Println("[USER] Trying to save...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[USER] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[USER] Commited!")

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
	tx := r.db.Begin()

	err := tx.Delete(&models.User{UserID: id}).Error
	fmt.Println("[USER] Trying to delete...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[USER] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[USER] Commited!")

	return RepositoryResult{Result: nil}
}

func (r *UserRepository) DeleteByIds(ids *[]string) RepositoryResult {
	tx := r.db.Begin()

	err := tx.Where("UserID IN (?)", *ids).Delete(&models.Users{}).Error
	fmt.Println("[USER] Trying to delete...")

	if err != nil {
		tx.Rollback()
		fmt.Println("[USER] Rollback...")
		return RepositoryResult{Error: err}
	}

	tx.Commit()
	fmt.Println("[USER] Commited!")

	return RepositoryResult{Result: nil}
}
