package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"api_test/storage"
	"api_test/storage/kv"
	"net/http"
)

func GetUser(c *gin.Context) {
	userID := c.Query("id")
	if user, err := kv.Get(userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func AddUser(c *gin.Context) {
	var newUser storage.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = uuid.New().String()
	if err := kv.Set(newUser.ID, &newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func UpdateUser(c *gin.Context) {
	var newUser storage.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := kv.Update(newUser.ID, &newUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return	
    }
    c.JSON(http.StatusOK, newUser)
}

func GetAllUsers(c *gin.Context) {
	users, err := kv.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	userID := c.Query("id")
	if err := kv.Delete(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.Status(http.StatusOK)
	}
}
