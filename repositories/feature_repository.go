package repositories

import (
	"context"
	"pr8_1/dtos"
	"pr8_1/models"

	"gorm.io/gorm"
)

type FeatureRepository struct {
	db *gorm.DB
}

func NewFeatureRepository(db *gorm.DB) *FeatureRepository {
	return &FeatureRepository{db: db}
}

func (r *FeatureRepository) Save(Feature *models.Feature) RepositoryResult {
	err := r.db.Save(Feature).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: Feature}
}

func (r *FeatureRepository) FindAll() RepositoryResult {
	var Features models.Features

	err := r.db.Find(&Features).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Features}
}

func (r *FeatureRepository) FindAllWithCtx(ctx *context.Context) RepositoryResult {
	var Features models.Features

	err := r.db.WithContext(*ctx).Find(&Features).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Features}
}

func (r *FeatureRepository) FindAllPaging(limit int, offset int, sort string, searchs []dtos.Search) (RepositoryResult, int64) {
	var Features models.Features
	var total int64

	find := r.db.Limit(limit).Offset(offset).Order(sort)
	err := applySearchs(find, searchs).Find(&Features).Count(&total).Error

	if err != nil {
		return RepositoryResult{Error: err}, 0
	}

	return RepositoryResult{Result: &Features}, total
}

func (r *FeatureRepository) FindOneById(id uint) RepositoryResult {
	var Feature models.Feature

	err := r.db.Where(&models.Feature{FeatureID: id}).Take(&Feature).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Feature}
}

func (r *FeatureRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&models.Feature{FeatureID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *FeatureRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("FeatureID IN (?)", *ids).Delete(&models.Features{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
