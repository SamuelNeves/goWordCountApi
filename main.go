package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

func HomePage(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	wordCount:=WordCount(string(value))
	fmt.Println(json.Marshal(wordCount))
	output,err:= json.Marshal(wordCount)
	c.JSON(200, gin.H{
		"message": string(output),
	})
}

func QueryString(c *gin.Context) {
	age := c.Query("age")
	name := c.Query("name")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameters(c *gin.Context) {
	age := c.Param("age")
	name := c.Param("name")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	wordCountMap := make(map[string]int)

	for _,word := range words{
		wordCountMap[word]++
	}

	return wordCountMap
}

func main() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryString)              // /query?name=elliot&age=24
	r.GET("/path/:name/:age", PathParameters) // /query?name=elliot&age=24
	r.Run()
}
