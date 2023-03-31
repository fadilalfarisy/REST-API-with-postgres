package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"challenge-2/database"
	"challenge-2/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()

	var newBook = models.Book{}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Book := models.Book{
		NameBook: newBook.NameBook,
		Author:   newBook.Author,
	}

	err := db.Create(&Book).Error

	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	fmt.Printf("New Book Data: %+v \n", newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": Book,
	})
}

func UpdateBooks(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var newBook = models.Book{}

	var number int
	var err error

	number, err = strconv.Atoi(bookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", number),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	book := models.Book{}

	result := db.First(&book).Where("id = ?", number).
		Updates(models.Book{
			NameBook: newBook.NameBook,
			Author:   newBook.Author,
		})

	if result.Error != nil {
		fmt.Println("Error updating book data: ", err)
		return
	}

	count := result.RowsAffected
	if count == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("books with id %v failed to updated", number),
		})
		return
	}

	fmt.Println("Updated data amount: ", count)

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	var number int
	var err error

	number, err = strconv.Atoi(bookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", number),
		})
		return
	}

	db := database.GetDB()

	book := models.Book{}

	err = db.First(&book, "id = ?", number).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{
				"error_status":  "Data Not Found",
				"error_message": fmt.Sprintf("book with id %v not found", number),
			})
			return
		}
		print("Error finding book: ", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetAllBooks(ctx *gin.Context) {
	db := database.GetDB()

	book := []models.Book{}

	err := db.Find(&book).Error

	if err != nil {
		fmt.Println("Error get user data: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	var number int
	var err error

	number, err = strconv.Atoi(bookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", number),
		})
		return
	}

	db := database.GetDB()

	book := models.Book{}

	result := db.Where("id = ?", number).Delete(&book)

	if result.Error != nil {
		fmt.Println("Error deleting book: ", err.Error())
		return
	}

	count := result.RowsAffected

	if count == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("books with id %v failed to deleted", number),
		})
		return
	}

	fmt.Println("Deleted data amount: ", count)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Book deleted successfully",
	})
}
