package repositories

import (
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
