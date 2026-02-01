package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"blink-server/middleware"
	"github.com/gin-gonic/gin"
)

// Config holds the application configuration
type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
}

// Company represents a company listed on NSE/BSE
type Company struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Ticker   string `json:"ticker"`
	Exchange string `json:"exchange"`
	Sector   string `json:"sector"`
	MarketCap string `json:"marketCap"`
}

// CompaniesData holds the list of companies
type CompaniesData struct {
	Companies []Company `json:"companies"`
}

var config Config
var companiesList []Company

// LoadConfig loads the configuration from application.json
func LoadConfig(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	log.Println("Configuration loaded successfully")
	return nil
}

// LoadCompanies loads the companies from companies.json
func LoadCompanies(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read companies file: %w", err)
	}

	var companiesData CompaniesData
	err = json.Unmarshal(data, &companiesData)
	if err != nil {
		return fmt.Errorf("failed to parse companies file: %w", err)
	}

	companiesList = companiesData.Companies
	log.Printf("Loaded %d companies successfully\n", len(companiesList))
	return nil
}

// GetConfig returns the loaded configuration
func GetConfig() Config {
	return config
}

// GetCompanies returns the list of all companies
func GetCompanies() []Company {
	return companiesList
}

// StartServer starts the Gin server with the configured port
func StartServer() {
	// Create a new Gin router with default middleware
	router := gin.Default()

	// Apply custom logging middleware
	router.Use(middleware.LoggerMiddleware())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Welcome endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Blink API",
			"version": "1.0.0",
		})
	})

	// Companies endpoint - returns all listed companies
	router.GET("/companies", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"total":       len(companiesList),
			"companies":   companiesList,
		})
	})

	// Start the server on the configured port
	port := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("Starting server on port %d...\n", config.Server.Port)
	router.Run(port)
}
