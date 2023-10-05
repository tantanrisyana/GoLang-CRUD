package main

import (
	"fmt"
	"os"
	"tugas/controllers"
	"tugas/models"
	"tugas/utils"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	configs := utils.LoadConfigs()

	db, err := utils.InitDB(configs.DBUsername, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}
	//defer db.Close()

	db.AutoMigrate(&models.Siswa{})

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(echojwt.JWT([]byte(os.Getenv("secret"))))

	siswaController := controllers.NewSiswaController(db)

	// Routes
	e.GET("/siswa", siswaController.GetSiswaList)
	e.GET("/siswa/:id", siswaController.GetSiswaByID)
	e.POST("/siswa", siswaController.CreateSiswa)
	e.PUT("/siswa/:id", siswaController.UpdateSiswa)
	e.DELETE("/siswa/:id", siswaController.DeleteSiswa)

	// Start server
	e.Start(":" + os.Getenv("PORT"))
}
