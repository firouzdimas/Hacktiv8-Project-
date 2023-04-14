package controllers

import (
	"fmt"
	"net/http"
	"project-2/config"
	"project-2/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) {
	var results = []models.Book{}
	sqlStatement := `SELECT * FROM books`

	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		fmt.Println("something is wrong")
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": results,
	})
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	var book models.Book
	sqlStatement := `SELECT * FROM books WHERE id = $1`

	row := config.DB.QueryRow(sqlStatement, bookID)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Description); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("book with id %v is not exist", bookID),
			"status":  "not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

func AddBook(ctx *gin.Context) {
	var newBook models.Book
	var book = models.Book{}

	if err := ctx.ShouldBindJSON(&newBook); err != nil || strings.TrimSpace(newBook.Title) == "" || strings.TrimSpace(newBook.Author) == "" || strings.TrimSpace(newBook.Description) == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid inputs, please rechek your input",
		})
		return
	}

	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	RETURNING *`

	err := config.DB.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Description).
		Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "new book successfully created",
		"book":    &book,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	var updatedBook models.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil || strings.TrimSpace(updatedBook.Title) == "" || strings.TrimSpace(updatedBook.Author) == "" || strings.TrimSpace(updatedBook.Description) == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid inputs, please recheck your inputs",
		})
		return
	}

	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1
	`
	res, err := config.DB.Exec(sqlStatement, bookID, updatedBook.Title, updatedBook.Author, updatedBook.Description)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("the book with id %v that you're trying to update is not exist", bookID),
			"status":  "not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfully updated", bookID),
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	sqlStatement := `DELETE FROM books WHERE id = $1`

	res, err := config.DB.Exec(sqlStatement, bookID)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("the book you're trying to delete with id %v is not exist", bookID),
			"status":  "not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deleted", bookID),
	})
}
