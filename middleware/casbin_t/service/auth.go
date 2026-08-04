package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gotest/common/module/cache"
	"gotest/middleware/casbin_t/models"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type AuthService struct {
	db        *gorm.DB
	jwtSecret []byte
}

func NewAuthService(db *gorm.DB, secret string) *AuthService {
	return &AuthService{db: db, jwtSecret: []byte(secret)}
}

func (s *AuthService) Login(username, password string) (string, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	claims := Claims{
		Username: user.Username,
		Role:     s.cachedRole(user.Username, user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.jwtSecret)
}

func (s *AuthService) ParseToken(raw string) (*Claims, error) {
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(raw, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected JWT signing method: %s", token.Header["alg"])
		}
		return s.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		if err == nil {
			err = errors.New("invalid JWT")
		}
		return nil, err
	}
	return claims, nil
}

// cachedRole uses common/module/cache's Redis pool as a five minute cache.
// If Redis is unavailable the database role remains authoritative.
func (s *AuthService) cachedRole(username, fallbackRole string) string {
	conn := cache.RdsPool.Get()
	defer conn.Close()

	key := "casbin_t:role:" + username
	role, err := redis.String(conn.Do("GET", key))
	if err == nil && role != "" {
		return role
	}
	_, _ = conn.Do("SETEX", key, 300, fallbackRole)
	return fallbackRole
}
