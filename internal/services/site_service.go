package services

import (
	"api-service/internal/models"
	"api-service/pkg"
	"api-service/pkg/mq"
	"context"
	"encoding/json"
	"log"
)

type siteService struct {
	repo        pkg.SiteRepository
	mqPublisher mq.Publisher
}

func NewSiteService(repo pkg.SiteRepository, mqPublisher mq.Publisher) pkg.SiteService {
	return &siteService{repo: repo, mqPublisher: mqPublisher}
}

func (s *siteService) Create(ctx context.Context, site *models.Site) (*models.Site, error) {
	if err := s.repo.Create(ctx, site); err != nil {
		return nil, err
	}

	msgBody, err := json.Marshal(site)
	if err != nil {
		log.Printf("CRITICAL: Error marshalling site for queue: %v. Site ID: %d", err, site.ID)
		return site, nil
	}

	log.Printf("Publishing site check task for site ID: %d", site.ID)
	if err := s.mqPublisher.Publish(ctx, mq.SiteCheckQueue, msgBody); err != nil {
		log.Printf("ERROR: Failed to publish site check task for site ID: %d. Error: %v", site.ID, err)
	}

	return site, nil
}

func (s *siteService) GetByID(ctx context.Context, id uint) (*models.Site, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *siteService) GetAll(ctx context.Context) ([]models.Site, error) {
	return s.repo.FindAll(ctx)
}

func (s *siteService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
