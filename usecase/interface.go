package usecase

import "web_scrapper/domain"

// Scraper interface defines the scraping functionality
type Scraper interface {
	ScrapeTablets(url string) ([]domain.Tablet, error)
}