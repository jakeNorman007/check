package main

import (
    "github.com/gin-gonic/gin"
    "github.com/JakeNorman007/check/initializers"
    "github.com/JakeNorman007/check/controllers"
    
)

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main() {
    r := gin.Default()

    r.POST("/", controllers.TodosCreate)
    r.GET("/", controllers.TodosIndex)
    r.PUT("/:id", controllers.TodosUpdate)
    r.DELETE("/:id", controllers.TodosDelete)

	r.Run()
}
