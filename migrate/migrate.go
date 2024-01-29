package main

import (
	"github.com/asim/go-crud/initializers"
	"github.com/asim/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
