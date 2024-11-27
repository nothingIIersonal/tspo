package repositories

import (
	"pr8_1/models"

	"gorm.io/gorm"
)

type GoodRepository struct {
	db *gorm.DB
}

func NewGoodRepository(db *gorm.DB) *GoodRepository {
	return &GoodRepository{db: db}
}

func (r *GoodRepository) Save(Good *models.Good) RepositoryResult {
	err := r.db.Save(Good).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: Good}
}

func (r *GoodRepository) FindAll() RepositoryResult {
	var Goods models.Goods

	err := r.db.Preload("Feature").Preload("Vendor").Find(&Goods).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Goods}
}

func (r *GoodRepository) FindOneById(id uint) RepositoryResult {
	var Good models.Good

	err := r.db.Preload("Feature").Preload("Vendor").Where(&models.Good{GoodID: id}).Take(&Good).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Good}
}

func (r *GoodRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&models.Good{GoodID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *GoodRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("GoodID IN (?)", *ids).Delete(&models.Goods{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
