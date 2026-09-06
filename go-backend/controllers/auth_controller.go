package controllers

import (
	"net/http"
	"strconv"

	"example.com/interview-question-087/go-backend/database"
	"example.com/interview-question-087/go-backend/middleware"
	"example.com/interview-question-087/go-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetCurrentUserID(c *gin.Context) *uint {
	if userIDVal, exists := c.Get("current_user_id"); exists {
		switch v := userIDVal.(type) {
		case uint:
			return &v
		case float64:
			u := uint(v)
			return &u
		case string:
			if parsed, err := strconv.ParseUint(v, 10, 32); err == nil {
				u := uint(parsed)
				return &u
			}
		}
	}
	return nil
}

func Register(c *gin.Context) {
	var input struct {
		Username        string `json:"username" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirmPassword" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "กรุณากรอกข้อมูลให้ครบถ้วน"})
		return
	}

	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Password และ Confirm Password ไม่ตรงกัน"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "ไม่สามารถเข้ารหัสรหัสผ่านได้"})
		return
	}

	user := models.User{
		Username:     input.Username,
		PasswordHash: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "ชื่อผู้ใช้นี้มีอยู่ในระบบแล้ว"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "ลงทะเบียนสำเร็จ"})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "กรุณากรอก Username และ Password"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "ชื่อผู้ใช้งานหรือรหัสผ่านไม่ถูกต้อง"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "ชื่อผู้ใช้งานหรือรหัสผ่านไม่ถูกต้อง"})
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "ไม่สามารถสร้าง Token ได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func GetProfile(c *gin.Context) {
	userID := GetCurrentUserID(c)
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "ไม่พบสิทธิ์การใช้งาน"})
		return
	}

	var user models.User
	if err := database.DB.Where("id = ?", *userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "ไม่พบข้อมูลผู้ใช้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}