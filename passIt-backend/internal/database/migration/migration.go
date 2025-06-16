package main

import "passIt/internal/database"

func main() {
	dbService := database.New()
	dbService.Migration()
	defer dbService.Close()
}
