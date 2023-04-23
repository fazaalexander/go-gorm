package controller

import (
	"net/http"

	"github.com/fazaalexander/go-gorm/config"
	"github.com/fazaalexander/go-gorm/model"

	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	var blogs *[]model.Blog

	if err := config.DB.Preload("User").Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all blogs",
		"blogs":   blogs,
	})
}

func GetBlogController(c echo.Context) error {
	var blog *[]model.Blog
	id := c.Param("id")

	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get blog by id",
		"blog":    blog,
	})
}

func CreateBlogController(c echo.Context) error {
	var blog *model.Blog
	c.Bind(&blog)

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success create new blog",
	})
}

func DeleteBlogController(c echo.Context) error {
	var blog *model.Blog
	id := c.Param("id")

	if err := config.DB.Delete(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Delete blog successful",
	})
}

func UpdateBlogController(c echo.Context) error {
	var blog *[]model.Blog
	id := c.Param("id")

	judul := c.FormValue("judul")
	konten := c.FormValue("konten")

	if err := config.DB.Model(&blog).Where("id = ?", id).Updates(model.Blog{Judul: judul, Konten: konten}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success update blog",
		"blog":    blog,
	})
}
