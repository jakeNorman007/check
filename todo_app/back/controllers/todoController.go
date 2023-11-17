package controllers

import (
	"github.com/JakeNorman007/check/initializers"
	"github.com/JakeNorman007/check/models"
	"github.com/gin-gonic/gin"
)

// create function, creates a todo list item
func TodosCreate(c *gin.Context){

    // get data off of the body of the request
    var todo struct{
        Body    string
    }

    // Bind method in Gin serializes the body struct into JSON
    c.Bind(&todo)

    // creates a todo item
    item := models.Todo{Body: todo.Body}

    result := initializers.DB.Create(&item)

    if result.Error != nil {
        c.Status(400)
        return
    }

    // returns the todo item
    c.JSON(200, gin.H{
        "item": item,
    })
}

func TodosIndex(c *gin.Context){

    // gets all the todo items
    var todos []models.Todo
    initializers.DB.Find(&todos)

    // responds with all of the todo items that curerntly exist
    c.JSON(200, gin.H{
        "todos": todos,
    })
}

func TodosUpdate(c *gin.Context){
    
    // gets the id off of the url/ id of the row in DB
    id := c.Param("id")

    // get the data off of the body of the todo we are referencing by it's id
    var todo struct{
        Body    string
    }

    // JSON
    c.Bind(&todo)

    // finding the todo item that is needing to be updated
    var item models.Todo
    initializers.DB.First(&item, id)

    // update the todo item
    initializers.DB.Model(&item).Updates(models.Todo{Body: todo.Body,})

    // responds with an updated todo item
    c.JSON(200, gin.H{
        "item": item ,
    })
}

func TodosDelete(c *gin.Context){
    
    // similar to updating, we need to grab the id of the row
    id := c.Param("id")

    // delete the todo item
    initializers.DB.Delete(&models.Todo{}, id)

    // respond with a 200 for confirmation
    c.Status(200)
}











