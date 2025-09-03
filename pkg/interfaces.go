package pkg

import (
	"api-service/internal/models"
	"context"
)

type SiteService interface {
	Create(ctx context.Context, site *models.Site) (*models.Site, error)
	GetByID(ctx context.Context, id uint) (*models.Site, error)
	GetAll(ctx context.Context) ([]models.Site, error)
	Delete(ctx context.Context, id uint) error
}

type CheckerService interface {
	CheckSite(ctx context.Context, site *models.Site) error
}

type SiteRepository interface {
	Create(ctx context.Context, site *models.Site) error
	FindByID(ctx context.Context, id uint) (*models.Site, error)
	FindAll(ctx context.Context) ([]models.Site, error)
	Update(ctx context.Context, site *models.Site) error
	Delete(ctx context.Context, id uint) error
}
