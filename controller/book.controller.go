package controller

import (
	"net/http"

	"github.com/fazaalexander/go-gorm/config"
	"github.com/fazaalexander/go-gorm/model"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success getting all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	var book []model.Book
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting book by id",
		"book":    book,
	})
}

// create new book
func CreateNewBook(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	var book []model.Book
	id := c.Param("id")

	if err := config.DB.Delete(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	var book []model.Book
	id := c.Param("id")
	judul := c.FormValue("judul")
	penulis := c.FormValue("penulis")
	penerbit := c.FormValue("penerbit")

	if err := config.DB.Model(&book).Where("id = ?", id).Updates(model.Book{Judul: judul, Penulis: penulis, Penerbit: penerbit}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success update book",
	})
}
