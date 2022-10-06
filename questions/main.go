package main

import (
	"github.com/gin-gonic/gin"
)

// Struct
type question struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Author   string `json:"author"`
}

// Slice or Map
var questions = []question{
	{ID: "1", Question: "What''s something you really resent paying for?", Author: "Kalese Carpenter"},
	{ID: "2", Question: "If you couldn''t be convicted of any one type of crime, what criminal charge would you like to be immune to?", Author: "Mandee Rizzi"},
	{ID: "3", Question: "In the past people were buried with the items they would need in the afterlife, what would you want buried with you so you could use it in the afterlife?", Author: "Zack Raney"},
	{ID: "4", Question: "Who do you go out of your way to be nice to?", Author: "John Butcher"},
	{ID: "5", Question: "What invention doesn''t get a lot of love, but has greatly improved the world?", Author: "Jessie Gorrell"},
	{ID: "6", Question: "What movie quotes do you use on a regular basis?", Author: "Eric Boggs"},
	{ID: "7", Question: "What “old person” things do you do?", Author: "Cha''Diamond Moody"},
}

// Get All Questions
// func getAllQuestions() {

// }

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
