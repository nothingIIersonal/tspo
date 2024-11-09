package mock_book

import (
	"net/http"
	"pr7_1/services/auth"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	token, exists := c.Get("token")
	refreshToken, refreshExists := c.Get("refreshToken")

	if !exists || !refreshExists {
		c.JSON(http.StatusOK, gin.H{"message": "get all books", "token": token, "refreshToken": refreshToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "get all books", "token": token, "refreshToken": refreshToken})
}

func CreateBook(c *gin.Context) {
	token, exists := c.Get("token")
	refreshToken, refreshExists := c.Get("refreshToken")

	if !auth.CheckClaims(c.Request.Header["Authorization"][0]) {
		if !exists || !refreshExists {
			c.JSON(http.StatusOK, gin.H{"message": "no permissions"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "no permissions", "token": token, "refreshToken": refreshToken})
		}
		return
	}

	if !exists || !refreshExists {
		c.JSON(http.StatusOK, gin.H{"message": "create a new book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create a new book", "token": token, "refreshToken": refreshToken})
}

func UpdateBook(c *gin.Context) {
	token, exists := c.Get("token")
	refreshToken, refreshExists := c.Get("refreshToken")

	if !auth.CheckClaims(c.Request.Header["Authorization"][0]) {
		if !exists || !refreshExists {
			c.JSON(http.StatusOK, gin.H{"message": "no permissions"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "no permissions", "token": token, "refreshToken": refreshToken})
		}
		return
	}

	if !exists || !refreshExists {
		c.JSON(http.StatusOK, gin.H{"message": "updates a book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updates a book", "token": token, "refreshToken": refreshToken})
}

func DeleteBook(c *gin.Context) {
	token, exists := c.Get("token")
	refreshToken, refreshExists := c.Get("refreshToken")

	if !auth.CheckClaims(c.Request.Header["Authorization"][0]) {
		if !exists || !refreshExists {
			c.JSON(http.StatusOK, gin.H{"message": "no permissions"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "no permissions", "token": token, "refreshToken": refreshToken})
		}
		return
	}

	if !exists || !refreshExists {
		c.JSON(http.StatusOK, gin.H{"message": "deletes a book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deletes a book", "token": token, "refreshToken": refreshToken})
}
