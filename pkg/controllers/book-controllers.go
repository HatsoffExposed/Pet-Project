package controllers

import (
	"fmt"
	"net/http"
	"petProject/pkg/models"
	"petProject/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	newBooks := models.GetAllBooks()
	// res, _ := json.Marshal(newBooks)
	// ress := string(res)
	// fmt.Printf("%v", ress)
	c.IndentedJSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	// var Book models.Book
	ID := c.Param("bookId")
	bookId, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(bookId)
	// res, _ := json.Marshal(bookDetails)
	c.IndentedJSON(http.StatusOK, bookDetails)
}

func CreateBook(c *gin.Context) {
	CreateBook := new(models.Book)
	utils.ParseBody(c, CreateBook)
	b := CreateBook.CreateBook()
	// res, _ := json.Marshal(b)
	c.IndentedJSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {
	ID := c.Param("bookId")
	bookId, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(bookId)
	// res, _ := json.Marshal(book)
	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var updateBook = &models.Book{}
	utils.ParseBody(c, updateBook)
	ID := c.Param("bookId")
	bookId, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(bookId)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	// res, _ := json.Marshal(bookDetails)
	c.IndentedJSON(http.StatusOK, bookDetails)
}
