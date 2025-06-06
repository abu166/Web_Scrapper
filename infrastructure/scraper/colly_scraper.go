package scraper

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"web_scrapper/domain"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

// CollyScraper implements the Scraper interface
type CollyScraper struct{}

// NewCollyScraper creates a new CollyScraper
func NewCollyScraper() *CollyScraper {
	return &CollyScraper{}
}

// ScrapeTablets scrapes tablet data from the given URL
func (s *CollyScraper) ScrapeTablets(url string) ([]domain.Tablet, error) {
	var tablets []domain.Tablet

	// Create a new chromedp context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set a timeout to avoid hanging
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Variable to store the HTML content
	var htmlContent string

	// Run the browser task
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		// Scroll to load all items
		chromedp.ActionFunc(func(ctx context.Context) error {
			for i := 0; i < 20; i++ {
				chromedp.Run(ctx,
					chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight);`, nil),
					chromedp.Sleep(1000*time.Millisecond),
				)
			}
			return nil
		}),
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to run chromedp: %v", err)
	}

	// Parse HTML using goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	// Extract tablet data
	doc.Find(".product-wrapper").Each(func(i int, e *goquery.Selection) {
		// Extract title
		title := strings.TrimSpace(e.Find(".title").Text())
		if title == "" {
			log.Printf("Title not found for item %d", i)
			return
		}

		// Extract price
		priceStr := strings.TrimPrefix(strings.TrimSpace(e.Find(".price").Text()), "$")
		price, priceErr := strconv.ParseFloat(priceStr, 64)
		if priceErr != nil {
			log.Printf("Price parse error for item %d: %v", i, priceErr)
			return
		}

		// Extract rating by counting star icons
		rating := e.Find(".ratings span.ws-icon.ws-icon-star").Length()
		if rating == 0 {
			log.Printf("No star ratings found for item %d, setting rating to 0", i)
		}

		// Extract image URL
		imageURL, imageOk := e.Find("img").Attr("src")
		if !imageOk {
			log.Printf("Image URL not found for item %d", i)
			return
		}

		// Extract description
		description := strings.TrimSpace(e.Find(".description").Text())
		if description == "" {
			log.Printf("Description not found for item %d", i)
			return
		}

		// Create tablet entry
		tablet := domain.Tablet{
			Title:       title,
			Price:       price,
			Description: description,
			Rating:      rating,
			ImageURL:    imageURL,
		}
		tablets = append(tablets, tablet)
	})

	return tablets, nil
}
