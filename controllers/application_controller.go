
package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"applicant-api/config"
	"applicant-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

var allowedStages = map[string]bool{
	"applied":   true,
	"screening": true,
	"interview": true,
	"offered":   true,
	"hired":     true,
	"rejected":  true,
}

type CreateApplicantInput struct {
	Name     string `json:"name" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email,max=150"`
	Phone    string `json:"phone" binding:"omitempty,max=20"`
	Position string `json:"position" binding:"required,max=100"`
}

type UpdateStageInput struct {
	Stage string `json:"stage" binding:"required,oneof=applied screening interview offered hired rejected"`
}

//aplikasi
func CreateApplicant(c *gin.Context) {
	var input CreateApplicantInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data applicant tidak valid",
			"details": err.Error(),
		})
		return
	}

	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.ToLower(strings.TrimSpace(input.Email))
	input.Position = strings.TrimSpace(input.Position)
	input.Phone = strings.TrimSpace(input.Phone)

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validasi data gagal",
			"details": err.Error(),
		})
		return
	}

	applicant := models.Applicant{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Position: input.Position,
		Stage:    "applied",
	}

	if err := config.DB.Create(&applicant).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email sudah terdaftar",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal membuat applicant",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Applicant berhasil dibuat",
		"data": applicant,
	})
}

// LIST APPLICANTS
func GetApplicants(c *gin.Context) {
	var applicants []models.Applicant

	result := config.DB.Order("created_at DESC").Find(&applicants)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data applicant",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data applicant berhasil diambil",
		"data": applicants,
	})
}

// UPDATE APPLICANT STAGE
func UpdateApplicantStage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID applicant tidak valid",
		})
		return
	}

	var input UpdateStageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stage tidak valid",
			"details": err.Error(),
		})
		return
	}

	input.Stage = strings.TrimSpace(input.Stage)

	if !allowedStages[input.Stage] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stage tidak diperbolehkan",
		})
		return
	}

	var applicant models.Applicant

	result := config.DB.First(&applicant, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Applicant tidak ditemukan",
		})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mencari applicant",
		})
		return
	}

	if err := config.DB.Model(&applicant).Update("stage", input.Stage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal memperbarui stage",
		})
		return
	}

	applicant.Stage = input.Stage

	c.JSON(http.StatusOK, gin.H{
		"message": "Stage applicant berhasil diperbarui",
		"data": applicant,
	})
}
