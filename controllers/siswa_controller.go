package controllers

import (
	"net/http"
	"tugas/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SiswaController struct {
	DB *gorm.DB
}

func NewSiswaController(db *gorm.DB) *SiswaController {
	return &SiswaController{DB: db}
}

func (c *SiswaController) GetSiswaList(ctx echo.Context) error {
	var siswaList []models.Siswa
	if err := c.DB.Find(&siswaList).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, siswaList)
}

func (c *SiswaController) GetSiswaByID(ctx echo.Context) error {
	id := ctx.Param("id")
	var siswa models.Siswa
	if err := c.DB.First(&siswa, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": "Siswa not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, siswa)
}

func (c *SiswaController) CreateSiswa(ctx echo.Context) error {
	var siswa models.Siswa
	if err := ctx.Bind(&siswa); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if err := c.DB.Create(&siswa).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, siswa)
}

func (c *SiswaController) UpdateSiswa(ctx echo.Context) error {
	id := ctx.Param("id")
	var siswa models.Siswa
	if err := c.DB.First(&siswa, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": "Siswa not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	if err := ctx.Bind(&siswa); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if err := c.DB.Save(&siswa).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, siswa)
}

func (c *SiswaController) DeleteSiswa(ctx echo.Context) error {
	id := ctx.Param("id")
	var siswa models.Siswa
	if err := c.DB.First(&siswa, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": "Siswa not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	if err := c.DB.Delete(&siswa).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Siswa deleted successfully"})
}
