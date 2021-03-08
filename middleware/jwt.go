package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yizheneng/gblog/config"
)

type MyClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`

	jwt.StandardClaims
}

func CreateToken(username string, role string) (token string, err error) {
	aliveTime := time.Now().AddDate(0, 0, 7)
	setClaims := MyClaims{
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: aliveTime.Unix(),
			Issuer:    "gbolg",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)

	token, err = reqClaim.SignedString(config.ServerSettings.JwtKey)
	if err != nil {
		return "", err
	}
	return token, err
}

// 验证token
func CheckToken(token string) (*MyClaims, error) {
	var claims MyClaims

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return config.ServerSettings.JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if setToken != nil {
		if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
			return key, nil
		} else {
			return nil, errors.New("Token error")
		}
	}
	return nil, errors.New("Token error")
}

// jwt 验证Token
func jwtCheckToken(c *gin.Context) (role string) {
	tokenHeader := c.Request.Header.Get("Authorization")
	if tokenHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "Token is emputy",
		})
		c.Abort()
		return
	}
	checkToken := strings.Split(tokenHeader, " ")
	if len(checkToken) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "Token is emputy",
		})
		c.Abort()
		return
	}

	if len(checkToken) != 2 && checkToken[0] != "Bearer" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "Token is emputy",
		})
		c.Abort()
		return
	}
	key, tokenErr := CheckToken(checkToken[1])
	if tokenErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": tokenErr.Error(),
		})
		c.Abort()
		return
	}

	username := c.PostForm("username")

	if key.UserName != username {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "Username error!",
		})
		c.Abort()
		return
	}

	c.Set("token_username", key.UserName)
	role = key.Role
	return
}

// jwt 中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtCheckToken(c)
		c.Next()
	}
}

// 后台接口Token验证
func JwtBackendToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if jwtCheckToken(c) != "admin" {
			c.JSON(http.StatusOK, gin.H{
				"status":  "Error",
				"message": "No permission!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
