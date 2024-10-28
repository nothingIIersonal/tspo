package books

import (
	"encoding/json"
	"io"
	"net/http"
	"proj/common"

	"github.com/gorilla/mux"
)

func init() {
	common.Books = []common.Book{
		{ID: "0", Title: "Book 0", ISBN: "123QWE0", PublishingHouse: "0", Genre: "0"},
		{ID: "1", Title: "Book 1", ISBN: "123QWE1", PublishingHouse: "1", Genre: "1"},
		{ID: "2", Title: "Book 2", ISBN: "123QWE2", PublishingHouse: "2", Genre: "2"},
		{ID: "3", Title: "Book 3", ISBN: "123QWE3", PublishingHouse: "3", Genre: "3"},
		{ID: "4", Title: "Book 4", ISBN: "123QWE4", PublishingHouse: "0", Genre: "4"},
		{ID: "5", Title: "Book 5", ISBN: "123QWE5", PublishingHouse: "1", Genre: "5"},
		{ID: "6", Title: "Book 6", ISBN: "123QWE6", PublishingHouse: "2", Genre: "0"},
		{ID: "7", Title: "Book 7", ISBN: "123QWE7", PublishingHouse: "3", Genre: "1"},
		{ID: "8", Title: "Book 8", ISBN: "123QWE8", PublishingHouse: "0", Genre: "2"},
		{ID: "9", Title: "Book 9", ISBN: "123QWE9", PublishingHouse: "1", Genre: "3"},
	}
}

// @Summary Get all books
// @Description Returns list of books
// @Tags books
// @Produce json
// @Success 200 {object} common.BooksResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /books [get]
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ret, err := common.GetBooksResponse(common.BooksResponse{Message: "books list", Data: common.Books})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Get one book
// @Description Returns book
// @Tags books
// @Produce json
// @Param id path string true "book ID"
// @Success 200 {object} common.BookResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /books/{id} [get]
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	book, _, err := common.GetOneByID(id, common.TBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	ret, err := common.GetBookResponse(common.BookResponse{Message: "book " + id, Data: book.(common.Book)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Create new book
// @Description Add book to list
// @Tags books
// @Accept json
// @Produce json
// @Param book body common.Book true "book data"
// @Success 201 {object} common.BookResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /books [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data common.Book

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, _, err := common.GetOneByID(data.Genre, common.TGenre); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "genre with provided id not exists"}))
		return
	}

	if _, _, err := common.GetOneByID(data.PublishingHouse, common.TPublishingHouse); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "publishing house with provided id not exists"}))
		return
	}

	if _, _, err := common.GetOneByID(data.ID, common.TBook); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "book already exists"}))
		return
	}

	common.Books = append(common.Books, data)

	ret, err := common.GetBookResponse(common.BookResponse{Message: "created book", Data: data})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.Writer.Write(w, ret)
}

// @Summary Update book
// @Description Updates existing book
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "book ID"
// @Param book body common.Book true "book data"
// @Success 200 {object} common.BookResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /books/{id} [put]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data common.Book

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	if _, n, err := common.GetOneByID(id, common.TBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "book not exists"}))
		return
	} else {
		if data.ID != "" {
			if _, _, err := common.GetOneByID(data.ID, common.TBook); err == nil {
				w.WriteHeader(http.StatusBadRequest)
				io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "Book with provided new id already exists"}))
				return
			}
			common.Books[n].ID = data.ID
		}
		if data.Title != "" {
			common.Books[n].Title = data.Title
		}
		if data.ISBN != "" {
			common.Books[n].ISBN = data.ISBN
		}
		if data.PublishingHouse != "" {
			if _, _, err := common.GetOneByID(data.PublishingHouse, common.TPublishingHouse); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "publishing house with provided id not exists"}))
				return
			}
			common.Books[n].PublishingHouse = data.PublishingHouse
		}
		if data.Genre != "" {
			if _, _, err := common.GetOneByID(data.Genre, common.TGenre); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "genre with provided id not exists"}))
				return
			}
			common.Books[n].Genre = data.Genre
		}
	}

	ret, err := common.GetBookResponse(common.BookResponse{Message: "updated book", Data: data})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Delete book
// @Description Deletes existing book
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "book ID"
// @Success 200 {object} common.BookResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /books/{id} [delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var data any
	if dataCopy, n, err := common.GetOneByID(id, common.TBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "book not exists"}))
		return
	} else {
		data = dataCopy
		common.RemoveSliceItemByID(n, common.TBook)
	}

	ret, err := common.GetBookResponse(common.BookResponse{Message: "deleted book", Data: data.(common.Book)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}
