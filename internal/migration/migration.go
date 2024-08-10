package main

import (
	"fmt"
	"smart-rental/internal/database"
	"smart-rental/internal/models"
)

func main() {
	 database.DatabaseConnection()
	 if err := database.DB.Debug().AutoMigrate(&models.User{}); err != nil {
        fmt.Printf("failed to auto migrate: %v", err)
    }
}
