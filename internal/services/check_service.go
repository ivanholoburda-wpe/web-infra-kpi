package services

import (
	"api-service/internal/models"
	"api-service/pkg"
	"context"
	"log"
	"net/http"
	"time"
)

type checkerService struct {
	repo       pkg.SiteRepository
	httpClient *http.Client
}

var _ pkg.CheckerService = (*checkerService)(nil)

func NewCheckerService(repo pkg.SiteRepository) pkg.CheckerService {
	return &checkerService{
		repo: repo,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *checkerService) CheckSite(ctx context.Context, site *models.Site) error {
	log.Printf("Checking site URL: %s", site.Url)

	req, err := http.NewRequestWithContext(ctx, "GET", site.Url, nil)
	if err != nil {
		log.Printf("Error creating request for site %s: %v", site.Url, err)
		site.HttpStatus = -1
		return s.repo.Update(ctx, site)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		log.Printf("Error during HTTP check for site %s: %v", site.Url, err)
		site.HttpStatus = -1
	} else {
		site.HttpStatus = resp.StatusCode
		resp.Body.Close()
	}

	if err := s.repo.Update(ctx, site); err != nil {
		log.Printf("Error updating site status in DB for %s: %v", site.Url, err)
		return err
	}

	log.Printf("Successfully processed site %s, new status: %d", site.Url, site.HttpStatus)
	return nil
}
