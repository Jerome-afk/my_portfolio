package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func InitServer() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("*.html")
	router.Static("/static", "./static")

	router.GET("/", func (c *gin.Context) {
		c.Redirect(http.StatusFound, "/home")
	})

	router.GET("/home", func (c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", nil)
	})

	router.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/static/") {
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"error": "Access forbidden.",
				"message": "Access to the page you are looking for is forbidden.",
			})
			c.Abort()
			return
		}
		c.Next()
	})

	// Handle form submission
	router.POST("/send", func (c *gin.Context)  {
		name := c.PostForm("name")
		email := c.PostForm("email")
		message := c.PostForm("message")

		if name == "" || email == "" || message == "" {
			c.HTML(http.StatusBadRequest, "main.html", gin.H{"error": "All fields are required"})
			return
		}

		// Email credentials
		smtpHost := "smtp.gmail.com"
		smtpPort := 587
		username := "otienojerome7@gmail.com"
		password := "uaye cnmq agyy dbhp"

		response := gin.H{}

		// Recipient email
		recipient := "otienojerome6@gmail.com"

		// Email content
		subject := "New Message from " + name
		body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: \n%s", name, email, message)

		// Send email
		m := gomail.NewMessage()
		m.SetHeader("From", username)
		m.SetHeader("To", recipient)
		m.SetHeader("Subject", subject)
		m.SetBody("text/plain", body)

		d := gomail.NewDialer(smtpHost, smtpPort, username, password)
		if err := d.DialAndSend(m); err != nil {
			log.Println("Failed to send message:", err)
			response = gin.H {
				"status": "error",
				"message": "Failed to send message.Please try again later.",
			}

			c.HTML(http.StatusInternalServerError, "main.html", response)
			return
		}

		response = gin.H{
			"status": "success",
			"message": "Your message was sent successfully!",
		}
		c.HTML(http.StatusOK, "main.html", response)
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "404 Not Found",
			"message": "The page you are looking for does not exist.",
		})
	})

	return router
}
