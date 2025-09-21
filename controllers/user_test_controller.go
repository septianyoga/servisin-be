package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myapp/config"
	"myapp/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	var total int64

	// Ambil semua data
	config.DB.Find(&users)

	// Hitung total
	config.DB.Model(&models.User{}).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  users,
	})
}

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Email: input.Email}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User Created Successfully", "data": user})
}

func UpdateUser(c *gin.Context) {
	// Ambil ID dari URL param
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Binding JSON
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update data
	updatedUser := models.User{
		Name:  input.Name,
		Email: input.Email,
	}
	config.DB.Model(&user).Updates(updatedUser)

	c.JSON(http.StatusOK, gin.H{"message": "User Updated Successfully", "data": user})
}

// DeleteUser godoc
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully", "data": true})
}
