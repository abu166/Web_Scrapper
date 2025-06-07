package main

import (
	"fmt"
	"log"
	"os"
	"web_scrapper/infrastructure/scraper"
	"web_scrapper/infrastructure/storage"
	"web_scrapper/usecase"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"),
		os.Getenv("PGHOST"), os.Getenv("PGPORT"))

	// Initialize repository
	repo, err := storage.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	// Initialize scraper
	scraper := scraper.NewCollyScraper()

	// Initialize usecase
	scraperUsecase, err := usecase.NewScraperUsecase(repo, scraper)
	if err != nil {
		log.Fatalf("Failed to initialize usecase: %v", err)
	}

	// Run the scraper
	err = scraperUsecase.ScrapeTablets("https://webscraper.io/test-sites/e-commerce/scroll/computers/tablets")
	if err != nil {
		log.Fatalf("Failed to scrape tablets: %v", err)
	}

	log.Println("Scraping completed successfully")
}
