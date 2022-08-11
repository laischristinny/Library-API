package controller

import (
	"LibraryAPI-GitFlow/src/database"
	"LibraryAPI-GitFlow/src/model"
	"LibraryAPI-GitFlow/src/repository"
	"LibraryAPI-GitFlow/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	books, err := repository.GetAllBooks()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book model.Book
	if err = json.Unmarshal(corpoRequest, &book); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}


	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	book.ID, err = repository.Create(book)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, book)

}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	bookID, erro := strconv.ParseUint(parameters["bookID"], 10, 64)
	if erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, err := database.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	book, err := repository.GetByID(bookID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, book)
}

func CheckoutBook(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	bookID, erro := strconv.ParseUint(parameters["bookID"], 10, 64)
	if erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, err := database.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	book, err := repository.CheckoutBook(bookID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, book)
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	bookID, erro := strconv.ParseUint(parameters["bookID"], 10, 64)
	if erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, err := database.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryBooks(db)
	book, err := repository.ReturnBook(bookID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, book)
}
