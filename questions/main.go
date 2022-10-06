package main

import (
	"math/rand"
	"net/http"
	"time"

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
func getAllQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questions)
}

// Get Question By ID
func getQuestionByID(c *gin.Context) {
	id := c.Param("id")
	//singleQuestion exists,
	for _, q := range questions {
		if q.ID == id {
			c.IndentedJSON(http.StatusOK, q)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Question Not Found"})
}

// Get a Random Question
func getRandomQuestions(c *gin.Context) {
	counter := 0
	randNum := rand.Intn(len(questions))
	for _, v := range questions {
		if counter == randNum {
			c.JSON(http.StatusOK, &v)
		}
		counter++
	}
	// rand.Seed(time.Now().Unix())
	// q := rand.Int() % len(questions)
	// return (questions[q])
}

func main() {
	// Initialize randomization
	rand.Seed(time.Now().Unix())
	// Initialize Gin Router
	router := gin.Default()
	// Initialize my Routes
	router.GET("/questions", getAllQuestions)
	router.GET("/", getRandomQuestions)
	router.GET("/questions/:id", getQuestionByID)
	router.Run("0.0.0.0:9090")
}
