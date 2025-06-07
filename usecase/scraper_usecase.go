package usecase

import (
	"fmt"
	"web_scrapper/domain"
)

// ScraperUsecase handles the business logic for scraping
type ScraperUsecase struct {
	repo    domain.Repository
	scraper Scraper
}

// NewScraperUsecase creates a new ScraperUsecase
func NewScraperUsecase(repo domain.Repository, scraper Scraper) (*ScraperUsecase, error) {
	if err := repo.Initialize(); err != nil {
		return nil, fmt.Errorf("failed to initialize repository: %v", err)
	}
	return &ScraperUsecase{
		repo:    repo,
		scraper: scraper,
	}, nil
}

// ScrapeTablets executes the scraping process and saves the results
func (u *ScraperUsecase) ScrapeTablets(url string) error {
	tablets, err := u.scraper.ScrapeTablets(url)
	if err != nil {
		return err
	}
	return u.repo.SaveTablets(tablets)
}
