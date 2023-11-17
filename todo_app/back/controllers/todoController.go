package controllers

import (
    "github.com/gin-gonic/gin"
)

// create function, creates a todo list item
func TodosCreate(c *gin.Context){
    c.JSON(200, gin.H{
        "message": "pong",
    })
}
