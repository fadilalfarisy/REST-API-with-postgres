package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"challenge-2/config"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID      int    `json:"book_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// CREATE TABLE books (
// 	id SERIAL PRIMARY KEY,
// 	title varchar(50) NOT NULL,
// 	author varchar(50) NOT NULL,
// 	description varchar(50) NOT NULL
// )

func CreateBook(ctx *gin.Context) {
	var newBook = Book{}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	Returning *
	`

	err := config.DB.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Description).
		Scan(&newBook.BookID, &newBook.Title, &newBook.Author, &newBook.Description)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Book Data: %+v \n", newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBooks(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var newBook Book

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

	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1;
	`
	res, err := config.DB.Exec(sqlStatement, number, newBook.Title, newBook.Author, newBook.Description)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("books with id %v failed to updated", number),
		})
		return
	}

	fmt.Println("Updated data amount: ", count)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully updated", bookID),
	})
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	var result = Book{}

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

	sqlStatement := fmt.Sprintf("SELECT * FROM books WHERE id = %d", number)

	rows, err := config.DB.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var books = Book{}

		err = rows.Scan(&books.BookID, &books.Title, &books.Author, &books.Description)

		if err != nil {
			panic(err)
		}

		fmt.Println("Books datas: ", books)

		result = books
	}

	if result.BookID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("book with id %v not found", number),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": result,
	})
}

func GetAllBooks(ctx *gin.Context) {
	var result = []Book{}

	sqlStatement := "SELECT * FROM books"

	rows, err := config.DB.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var books = Book{}

		err = rows.Scan(&books.BookID, &books.Title, &books.Author, &books.Description)

		if err != nil {
			panic(err)
		}

		result = append(result, books)
	}

	fmt.Println("Books datas: ", result)

	ctx.JSON(http.StatusOK, gin.H{
		"books": result,
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

	sqlStatement := `
	DELETE FROM books
	WHERE id = $1;
	`
	res, err := config.DB.Exec(sqlStatement, number)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("books with id %v failed to deleted", number),
		})
		return
	}

	fmt.Println("Deleted data amount: ", count)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("books with id %v has been successfully deleted", number),
	})
}
