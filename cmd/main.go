package main

import (
	"log"
	"web_scrapper/infrastructure/scraper"
	"web_scrapper/infrastructure/storage"
	"web_scrapper/usecase"
)

func main() {
	// Initialize repository
	repo := storage.NewJSONStorage("tablets.json")

	// Initialize scraper
	scraper := scraper.NewCollyScraper()

	// Initialize usecase
	scraperUsecase := usecase.NewScraperUsecase(repo, scraper)

	// Run the scraper
	err := scraperUsecase.ScrapeTablets("https://webscraper.io/test-sites/e-commerce/scroll/computers/tablets")
	if err != nil {
		log.Fatalf("Failed to scrape tablets: %v", err)
	}

	log.Println("Scraping completed successfully")
}