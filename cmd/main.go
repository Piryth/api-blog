package main

import (
	"context"
	"fmt"
	"github.com/dstotijn/go-notion"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
	"piryth.fr/blog/api"
	"piryth.fr/blog/database"
)

func main() {

	//pageBlocksToMarkdown()
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Initialize the database connection
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",
		dbUser, dbName, dbPassword, dbHost, dbPort)
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Create a new Queries instance
	queries := database.New(pool)

	// Initialize the Gin router
	r := gin.Default()

	// Set up the routes
	api.SetupRoutes(r, queries)

	// Start the server
	err = r.Run(":8080")
	if err != nil {
		return
	}
}

func pageBlocksToMarkdown() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	client := notion.NewClient(os.Getenv("NOTION_API_KEY"))

	pagination := notion.PaginationQuery{
		"",
		50,
	}

	blocks, err := client.FindBlockChildrenByID(context.Background(), "21682ad536ce805bb798e59dd58af036", &pagination)
	if err != nil {
		log.Fatalf("Impossible to query Notion", err)
	}

	for index, block := range blocks.Results {
		log.Println(block.MarshalJSON())
		log.Println(index)
	}

	// INPUT : page id

	// QUERY : get block children

	// TREATMENT : iterate over each block

	// FOR EACH : convert the bloc into markdown append it

	// OUTPUT : markdown formatted text

}
