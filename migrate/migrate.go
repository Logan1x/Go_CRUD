package main

import (
	"github.com/logan1x/go_crud/initializers"
	"github.com/logan1x/go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}
