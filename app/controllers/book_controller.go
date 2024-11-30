package controllers

import (
	"restAPI/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func GetBooks(c *fiber.Ctx) error { 
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true, 
			"msg": err.Error(),
		})
	}

	books, err := db.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
            "msg": 	 "books were not found",
            "count": 0,
            "books": nil,
		})
	}

    // Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false, 
		"msg": nil, 
		"count": len(books), 
		"books": books,
	})
}

// GetBook func gets book by given ID or 404 error.
// @Description Get book by given ID.
// @Summary get book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Router /v1/book/{id} [get]
func GetBook(c *fiber.Ctx) error { 
	id, err := uuid.Parse(c.Params("id"))
	if err != nil { 
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true, 
			"msg": err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true, 
			"msg": err.Error(),
		})
	}

	book, err := db.GetBook(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true, 
			"msg": "book with the given ID is not found",
			"book": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg": nil, 
		"book": book,
	})
}