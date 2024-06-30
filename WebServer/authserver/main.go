package main

import (
	"log"
	"net/http"
	signinup "webser/signInUp"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=qwerty dbname=mydatabase sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	signinup.SetDB(db)

	router := gin.Default()
	router.LoadHTMLGlob("*.html")
	router.StaticFS("/web", http.Dir("web"))
	router.GET("/signin.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/signup.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/welcome.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/sign-up", signinup.SignUp)
	router.POST("/sign-in", signinup.SignIn)

	router.Run(":8888")
}

// type Test struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

// var test = []Test{
// 	{1, "Grisha"},
// 	{2, "Artyom"},
// }

// func getTest(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, test)
// }
