# Web Scraper Project

This project scrapes tablet product data from a test website and stores it in a PostgreSQL database. It uses **Colly** with **ChromeDP** for scraping, **GoQuery** for HTML parsing, and **PostgreSQL** as the storage backend.

## 🧰 Features

* Scrapes tablet product data (title, price, description, rating, image URL)
* Stores scraped data into a PostgreSQL database
* Uses Docker for containerization of both application and database
* Implements clean architecture (domain, infrastructure, use case layers)

## 📦 Prerequisites

Ensure you have the following installed:
* Docker
* Docker Compose

## 🚀 How to Run the Project

### 1. Clone the Repository

```bash
git clone https://github.com/abu166/Web_Scrapper.git
cd web_scrapper
```

### 2. Start the Application Using Docker Compose

```bash
docker-compose up -d --build
```

This will:
* Build and start the Go application container.
* Start a PostgreSQL database container.
* Automatically run the scraper when the app starts.

You should see logs indicating that scraping has completed successfully.

## 🔍 Checking the Scraped Data

To verify that the tablet data was successfully scraped and stored:

**Option 1: Connect via `psql` inside Docker**

```bash
docker exec -it web_scrapper_db_1 psql -U admin -d web_scrapper
```

Then run:

```sql
SELECT * FROM tablets;
```

You should see all the tablet records inserted by the scraper.

## 📁 Project Structure Overview

```
web_scrapper/
├── cmd/
│   └── main.go                 # Entry point of the application
├── domain/
│   ├── model.go               # Tablet struct definition
│   └── repository.go          # Repository interface
├── infrastructure/
│   ├── scraper/
│   │   └── colly_scraper.go   # ChromeDP + Colly implementation
│   └── storage/
│       └── postgres_storage.go # PostgreSQL persistence logic
├── usecase/
│   ├── interface.go           # Scraper interface
│   └── scraper_usecase.go     # Business logic layer
├── migrations/
│   └── ...                    # Database migration files
├── docker-compose.yml
├── Dockerfile
├── .env                       # Environment variables
└── README.md                  # General Information
```

## 🛠️ Configuration

All configuration is handled through the `.env` file:

```env
PGUSER=admin
PGPASSWORD=admin
PGDATABASE=web_scrapper
PGHOST=db
PGPORT=5432
```

Note: `PGHOST` must match the service name defined in `docker-compose.yml`.

## 🧪 Test Site

The scraper targets this URL:
```
https://webscraper.io/test-sites/e-commerce/scroll/computers/tablets
```

It scrolls the page multiple times to simulate infinite scroll behavior and extracts product cards.

## 🧹 Optional: Clean Up

To stop and remove containers:

```bash
docker-compose down
```

To also remove the persistent volume (will delete all stored data):

```bash
docker-compose down -v
```

## 📄 License

MIT License – see [LICENSE](LICENSE) for details.