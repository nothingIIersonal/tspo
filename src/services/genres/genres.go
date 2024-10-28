package genres

import (
	"encoding/json"
	"io"
	"net/http"
	"proj/common"

	"github.com/gorilla/mux"
)

func init() {
	common.Genres = []common.Genre{
		{ID: "0", Title: "Drama"},
		{ID: "1", Title: "Fantasy"},
		{ID: "2", Title: "Detective"},
		{ID: "3", Title: "Handbook"},
		{ID: "4", Title: "Sci-fi"},
		{ID: "5", Title: "Dictionary"},
	}
}

// @Summary Get all genres
// @Description Returns list of genres
// @Tags genres
// @Produce json
// @Success 200 {object} common.GenresResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /genres [get]
func GetGenres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ret, err := common.GetGenresResponse(common.GenresResponse{Message: "genres list", Data: common.Genres})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Get one genre
// @Description Returns genre
// @Tags genres
// @Produce json
// @Param id path string true "genre ID"
// @Success 200 {object} common.GenreResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /genres/{id} [get]
func GetGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	genre, _, err := common.GetOneByID(id, common.TGenre)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	ret, err := common.GetGenreResponse(common.GenreResponse{Message: "genre " + id, Data: genre.(common.Genre)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Create new genre
// @Description Add genre to list
// @Tags genres
// @Accept json
// @Produce json
// @Param genre body common.Genre true "genre data"
// @Success 201 {object} common.GenresResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /genres [post]
func CreateGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data common.Genre

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, _, err := common.GetOneByID(data.ID, common.TGenre); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "genre already exists"}))
		return
	}

	common.Genres = append(common.Genres, data)

	ret, err := common.GetGenreResponse(common.GenreResponse{Message: "created genre", Data: data})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.Writer.Write(w, ret)
}

// @Summary Update genre
// @Description Updates existing genre
// @Tags genres
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Param genre body common.Genre true "genre data"
// @Success 200 {object} common.GenreResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /genres/{id} [put]
func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data common.Genre

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	if _, n, err := common.GetOneByID(id, common.TGenre); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "genre not exists"}))
		return
	} else {
		if data.ID != "" {
			if _, _, err := common.GetOneByID(data.ID, common.TGenre); err == nil {
				w.WriteHeader(http.StatusBadRequest)
				io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "Genre with provided new id already exists"}))
				return
			}
			common.Genres[n].ID = data.ID
		}
		if data.Title != "" {
			common.Genres[n].Title = data.Title
		}
	}

	ret, err := common.GetGenreResponse(common.GenreResponse{Message: "updated genre", Data: data})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Delete genre
// @Description Deletes existing genre
// @Tags genres
// @Accept json
// @Produce json
// @Param id path string true "genre ID"
// @Success 200 {object} common.GenreResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /genres/{id} [delete]
func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var data any
	if dataCopy, n, err := common.GetOneByID(id, common.TGenre); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, common.GetBeautyJSONByMap(map[string]any{
			"message": "genre not exists",
		}))
		return
	} else {
		data = dataCopy
		common.RemoveSliceItemByID(n, common.TGenre)
	}

	ret, err := common.GetGenreResponse(common.GenreResponse{Message: "deleted genre", Data: data.(common.Genre)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}
