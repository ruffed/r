package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	// "os/exec"

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
		var src Source

		f, err := os.CreateTemp("", "example.*.c")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(f.Name()) // clean up

		if err := c.BindJSON(&src); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, err := f.Write([]byte(src.Source)); err != nil {
			f.Close()
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("cc", f.Name())
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		out, err := exec.Command("objdump", "-S", "a.out").Output()
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"res": string(out)})
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
