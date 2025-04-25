package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	. "github.com/dev-rever/cryptoo-pricing/utils"
)

var jwtSecret = []byte("your-secret-key")

type MClaims struct {
	UID uint `json:"uid"`
	jwt.RegisteredClaims
}

func ProvideJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseError(AuthorizedErrorCode, "Missing or invalid Authorization header"))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenStr, &MClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseError(AuthorizedErrorCode, "Invalid token"))
			return
		}

		claims := token.Claims.(*MClaims)
		c.Set("uid", claims.UID)
		c.Next()
	}
}

func GenerateJWT(uid uint) (string, error) {
	claims := MClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
