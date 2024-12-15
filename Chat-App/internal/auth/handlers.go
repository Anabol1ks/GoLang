package auth

import (
	"chat_app/internal/storage"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterHandler godoc
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя с указанием имени и пароля
// @Tags auth
// @Accept json
// @Produce json
// @Param input body auth.RegisterInput true "Данные пользователя"
// @Failure 201 {object} SuccessResponse "Пользователь успешно зарегистрирован"
// @Failure 500 {object} ErrorResponse "Ошибка хеширования или создания пользователя"
// @Failure 409 {object} ErrorResponse "Имя пользователя или почта уже используется"
// @Router /auth/register [post]
func RegisterHandler(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// проверка уникальности
	var existingUser User
	if err := storage.DB.Where("email = ? OR username = ?", input.Email, input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Имя пользователя или почта уже используется"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования пароля"})
		return
	}

	user := User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := storage.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
}

var jwtSecret = []byte(os.Getenv("JWT_KEY"))

type LoginInput struct {
	EmailOrUsername string `json:"email_or_username" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

// LoginHandler godoc
// @Summary Авторизация пользователя
// @Description Производит авторизацию пользователя по имени пользователя и паролю, возвращая токен для доступа
// @Tags auth
// @Accept json
// @Produce json
// @Param input body auth.LoginInput true "Данные для входа"
// @Success 200 {object} TokenResponse "Токен доступа"
// @Success 500 {object} ErrorResponse "Ошибка генерации токена"
// @Failure 401 {object} ErrorResponse "Неверные данные для входа"
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := storage.DB.Where("email = ? OR username = ?", input.EmailOrUsername, input.EmailOrUsername).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные данные для входа"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные данные для входа"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
