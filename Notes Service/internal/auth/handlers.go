package auth

import (
	"net/http"
	"notes-service/internal/storage"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("khsiji8ncopad20jiucnu4n_dk9-021nmisk")

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterHandler godoc
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя с указанием имени и пароля
// @Tags auth
// @Accept json
// @Produce json
// @Param input body auth.LoginInput true "Данные пользователя"
// @Success 201 {object} map[string]string "Успешная регистрация"
// @Failure 500 {object} map[string]string "Ошибка хеширования или создания пользователя"
// @Router /auth/register [post]
func RegisterHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования"})
		return
	}

	user := User{Username: input.Username, Password: hashedPassword}
	if err := storage.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Успешная регистрация"})
}

// LoginHandler godoc
// @Summary Авторизация пользователя
// @Description Производит авторизацию пользователя по имени пользователя и паролю, возвращая токен для доступа
// @Tags auth
// @Accept json
// @Produce json
// @Param input body auth.LoginInput true "Данные для входа"
// @Success 200 {object} map[string]string "Токен доступа"
// @Failure 400 {object} map[string]string "Ошибка валидации данных"
// @Failure 401 {object} map[string]string "Неправильный логин или пароль"
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := storage.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный логин или пароль"})
		return
	}

	if !CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный логин или пароль"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
