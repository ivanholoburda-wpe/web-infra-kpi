package repository

import (
	"api-service/internal/models"
	"api-service/pkg"
	"context"
	"gorm.io/gorm"
)

type siteRepositoryGORM struct {
	db *gorm.DB
}

var _ pkg.SiteRepository = (*siteRepositoryGORM)(nil)

func NewSiteRepositoryGORM(db *gorm.DB) pkg.SiteRepository {
	return &siteRepositoryGORM{db: db}
}

func (r *siteRepositoryGORM) Create(ctx context.Context, site *models.Site) error {
	return r.db.WithContext(ctx).Create(site).Error
}

func (r *siteRepositoryGORM) FindByID(ctx context.Context, id uint) (*models.Site, error) {
	var site models.Site
	if err := r.db.WithContext(ctx).First(&site, id).Error; err != nil {
		return nil, err
	}
	return &site, nil
}

func (r *siteRepositoryGORM) FindAll(ctx context.Context) ([]models.Site, error) {
	var sites []models.Site
	if err := r.db.WithContext(ctx).Find(&sites).Error; err != nil {
		return nil, err
	}
	return sites, nil
}

func (r *siteRepositoryGORM) Update(ctx context.Context, site *models.Site) error {
	return r.db.WithContext(ctx).Save(site).Error
}

func (r *siteRepositoryGORM) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Site{}, id).Error
}
