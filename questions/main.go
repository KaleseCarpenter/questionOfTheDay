package main

import (
	"github.com/gin-gonic/gin"
)

// Struct

// Slice or Map

// Get All Questions
func getAllQuestions() {

}

// Get Question By ID
// func getQuestionByID()  {

// }

// Get a Random Question
// func getRandomQuestions()  {

// }

func main() {
	// Initialize randomization

	// Initialize Gin Router
	router := gin.Default()
	// Initialize my Routes
	router.GET("/questions", getAllQuestions)
	router.GET("/questions", getRandomQuestions)
	router.GET("/questions/:id", getQuestionByID)
	router.Run("0.0.0.0:8080")
}
