package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var users map[string]string
var roles map[string]string
var blacklist []string

func init() {
	users = map[string]string{
		"user1": "pass1",
		"user2": "pass2",
		"user3": "pass3",
		"user4": "pass4",
	}

	// strict => only get
	// all => all endpoints
	roles = map[string]string{
		"user1": "strict",
		"user2": "all",
		"user3": "strict",
		"user4": "strict",
	}
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	pass, ok := users[creds.Username]
	if !ok || pass != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, refreshToken, err := generateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})
}

func Register(c *gin.Context) {
	var creds RegisterCredentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	_, ok := users[creds.Username]
	if ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user exists"})
		return
	}

	users[creds.Username] = creds.Password
	roles[creds.Username] = creds.Role

	c.JSON(http.StatusOK, gin.H{"message": "successfully registered"})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		isBlocked := false

		for _, tok := range blacklist {
			if tok == tokenString {
				isBlocked = true
			}
		}

		if isBlocked || err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		isBlocked = false

		refreshTokenString := c.GetHeader("Refresh-Token")
		if refreshTokenString != "" {
			for _, tok := range blacklist {
				if tok == refreshTokenString {
					isBlocked = true
				}
			}
			refreshClaims := &Claims{}
			refreshToken, err := jwt.ParseWithClaims(refreshTokenString, refreshClaims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil || !refreshToken.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "refresh token invalid or other error occurs"})
			} else {
				blacklist = append(blacklist, tokenString)
				blacklist = append(blacklist, refreshTokenString)
				newToken, newRefreshToken, err := generateToken(refreshClaims.Username)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
					return
				}

				c.Set("token", newToken)
				c.Set("refreshToken", newRefreshToken)
			}
		}

		c.Next()
	}
}

func extractClaims(tokenStr string) *Claims {
	claims := &Claims{}

	token, _ := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if token.Valid {
		return claims
	}

	return nil
}

func CheckClaims(tokenStr string) bool {
	claims := extractClaims(tokenStr)
	return roles[claims.Username] != "strict"
}

var jwtKey = []byte("so much depends upon a red wheel barrow glazed with rain water beside the white chickens")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(username string) (string, string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	expirationTimeRefresh := time.Now().Add(8 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	claimsRefresh := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeRefresh.Unix(),
		},
	}
	token, err1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err1 != nil {
		return "", "", err1
	}
	refreshToken, err2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh).SignedString(jwtKey)
	if err2 != nil {
		return "", "", err2
	}
	return token, refreshToken, nil
}
