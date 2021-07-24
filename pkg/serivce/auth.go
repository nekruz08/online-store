package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/nekruz08/online-store/models"
	"github.com/nekruz08/online-store/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"

)


const(
	salt =  "dkaskdsa21312das3das"
	signingKey  = "quiche#SKDJASDKA3dsa213#sH"
	tokenTTL = 12 *time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// SERVICE берет у папки Repository

type AuthService struct {
	repo repository.Authorization

}
//AdminChecker  - проверка на Роль Админа
func (s *AuthService) AdminChecker(userId int) (bool, error) {
	return s.repo.AdminChecker(userId)
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

//CreateUser - создает пользователя и хеширует пароль
func (s *AuthService) CreateUser(user models.User) (int, error) {
	//s.repo.CreateUser() // он вызывает интерфейс, а тот в свою очередь логику бд
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)

}

// GenerateToken - Генируерут токен
func (s *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.GetUser(username,  generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	// генерация  токена, если такой юзер существует, то генируем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(), // время жизни токена
		IssuedAt: time.Now().Unix(), // когда токен был создан
	},
		user.Id,
	})

	// вернем подписанный токен
	return token.SignedString([]byte(signingKey))

}

// ParseToken - парсить токен и возвращает его ID
func (s *AuthService) ParseToken(accessToken string) (int, error)  {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok{
		return 0, errors.New("token claims are not of type *tokenClains")
	}

	return claims.UserId, nil
}


// generatePasswordHash - хешируем пароль
func generatePasswordHash(password string)  string {
	hash := sha1.New()
	hash.Write([]byte(password))


	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
