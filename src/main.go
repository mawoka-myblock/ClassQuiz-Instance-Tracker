package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mawoka-myblock/ClassQuiz-Instance-Tracker/src/models"
	"net/http"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.LoadHTMLGlob("src/templates/*.html")
	r.GET("/private", func(c *gin.Context) {
		var instances []models.Instance
		models.DB.Order("created_at desc").Find(&instances)
		c.HTML(http.StatusOK, "main.html", gin.H{
			"instances": instances,
		})
	})

	r.POST("/public/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var json struct {
			Users          uint `json:"users"`
			PublicQuizzes  uint `json:"public_quizzes"`
			PrivateQuizzes uint `json:"private_quizzes"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var instance models.Instance

		if err := models.DB.Where("id = ?", id).First(&instance).Error; err != nil {
			if err := models.DB.Create(&models.Instance{
				ID:             id,
				Users:          json.Users,
				PublicQuizzes:  json.PublicQuizzes,
				PrivateQuizzes: json.PrivateQuizzes,
				IP:             c.RemoteIP(),
			}).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "something went wrong"})
				return
			}

			return
		} else {
			models.DB.Model(&instance).Updates(map[string]interface{}{
				"Users":          json.Users,
				"PublicQuizzes":  json.PublicQuizzes,
				"PrivateQuizzes": json.PrivateQuizzes,
				"IP":             c.RemoteIP(),
			})
		}
		fmt.Println(instance, "hi")

	})

	r.Run()
}
