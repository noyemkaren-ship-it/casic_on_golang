package main

import (
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	InitDB()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		token, _ := c.Cookie("token")
		if token == "" {
			c.Redirect(302, "/login")
			return
		}
		balance := get_balance(token)
		c.HTML(200, "index.html", gin.H{"Balance": balance})
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	r.POST("/register", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		create_user(name, password)
		token := get_token(name)
		c.SetCookie("token", token, 3600*24*7, "/", "localhost", false, true)
		c.Redirect(302, "/")
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		if login(name, password) {
			token := get_token(name)
			c.SetCookie("token", token, 3600*24*7, "/", "localhost", false, true)
			c.Redirect(302, "/")
		} else {
			c.HTML(200, "login.html", gin.H{"Error": "Неверные данные"})
		}
	})

	r.GET("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.Redirect(302, "/login")
	})

	r.POST("/play", func(c *gin.Context) {
		token, _ := c.Cookie("token")
		if token == "" {
			c.Redirect(302, "/login")
			return
		}

		amount, _ := strconv.Atoi(c.PostForm("amount"))
		currentBalance := get_balance(token)

		if currentBalance < amount {
			c.HTML(200, "index.html", gin.H{"Balance": currentBalance, "Message": "Недостаточно средств"})
			return
		}

		num := rand.Intn(2)
		if num == 0 {
			change_balance(token, amount)
			c.HTML(200, "index.html", gin.H{"Balance": get_balance(token), "Message": "Ты выиграл!"})
		} else {
			change_balance(token, -amount)
			c.HTML(200, "index.html", gin.H{"Balance": get_balance(token), "Message": "Ты проиграл!"})
		}
	})

	r.GET("/deposit/:amount", func(c *gin.Context) {
		amount, _ := strconv.Atoi(c.Param("amount"))
		token, _ := c.Cookie("token")
		change_balance(token, amount)
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.Run(":2222")
}
