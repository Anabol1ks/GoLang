package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sqlx.DB
var jwtKey = []byte("dpokfqkl-023fcs")

type ClientLog struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignUp(c *gin.Context) {
	var newUser ClientLog
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Хз что за ошибка"})
		return
	}
	hashPas, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось хешировать"})
		return
	}
	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", newUser.Username, string(hashPas))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user into database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь зарегистрирован"})

}

func Login(c *gin.Context) {
	var user ClientLog
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Хз что за ошибка"})
		return
	}

	var dbUser User
	err := db.Get(&dbUser, "SELECT username, password_hash FROM users WHERE username=$1", user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
		return
	}

	ex := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ex.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Токен не создался"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

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

	router := gin.Default()
	router.LoadHTMLGlob("*.html")
	router.StaticFS("/login", http.Dir("login"))
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

	router.POST("/sign-up", SignUp)
	router.POST("/sign-in", Login)

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
