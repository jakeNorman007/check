package main

import (
    "github.com/JakeNorman007/check/initializers"
    "github.com/JakeNorman007/check/models"
)

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main() {
    initializers.DB.AutoMigrate(&models.Todo{})
}
