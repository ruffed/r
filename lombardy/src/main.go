package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Source struct {
	Source string `json:"source"`
}

// String has *receiver* type *Source
func (s *Source) String() string {
	return fmt.Sprintf("%s", s.Source)
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/ace-builds", "./ace-builds")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Landing",
		})
	})
	r.POST("/compile", func(c *gin.Context) {
		// var src Source

		compileText := c.PostForm("compileText")
		//fmt.Println(compileText)

		// TODO: get compiled code here
		cmd := exec.Command("cc", "main.c")
		if err := cmd.Run(); err != nil {
			panic(err)
		}

		cmd = exec.Command("objdump", "-S", "a.out")
		if err := cmd.Run(); err != nil {
			panic(err)
		}

		// objdump -S a.out

		c.HTML(http.StatusOK, "index.html", gin.H{
			"editor2text": compileText,
		})

		// if err := c.BindJSON(&src); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		//c.JSON(http.StatusOK, gin.H{})
	})
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}

// go map -> python dict
// map[string]interface{}
//
// Any
// dict[string, int]

// Go doesn't have subclasses
// - no inheritance
// no polymorphism

// interfaces
// https://gobyexample.com/interfaces
