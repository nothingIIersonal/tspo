package cities

import (
	"encoding/json"
	"io"
	"net/http"
	"proj/common"

	"github.com/gorilla/mux"
)

func init() {
	common.Cities = []common.City{
		{ID: "0", Title: "Moscow"},
		{ID: "1", Title: "SPB"},
		{ID: "2", Title: "New York"},
		{ID: "3", Title: "LA"},
	}
}

// @Summary Get all cities
// @Description Returns list of cities
// @Tags cities
// @Produce json
// @Success 200 {object} common.CitiesResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /cities [get]
func GetCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ret, err := common.GetCitiesResponse(common.CitiesResponse{Message: "cities list", Data: common.Cities})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Get one city
// @Description Returns city
// @Tags cities
// @Produce json
// @Param id path string true "city ID"
// @Success 200 {object} common.CityResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /cities/{id} [get]
func GetCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	city, _, err := common.GetOneByID(id, common.TCity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	ret, err := common.GetCityResponse(common.CityResponse{Message: "city " + id, Data: city.(common.City)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Create new city
// @Description Add city to list
// @Tags cities
// @Accept json
// @Produce json
// @Param city body common.City true "city data"
// @Success 201 {object} common.CityResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /cities [post]
func CreateCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data common.City

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, _, err := common.GetOneByID(data.ID, common.TCity); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "city already exists"}))
		return
	}

	common.Cities = append(common.Cities, data)

	ret, err := common.GetCityResponse(common.CityResponse{Message: "created city", Data: data})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.Writer.Write(w, ret)
}

// @Summary Update city
// @Description Updates existing city
// @Tags cities
// @Accept json
// @Produce json
// @Param id path string true "city ID"
// @Param city body common.City true "city data"
// @Success 200 {object} common.CityResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /cities/{id} [put]
func UpdateCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data common.City

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	if _, n, err := common.GetOneByID(id, common.TCity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "city not exists"}))
		return
	} else {
		if data.ID != "" {
			if _, _, err := common.GetOneByID(data.ID, common.TCity); err == nil {
				w.WriteHeader(http.StatusBadRequest)
				io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "City with provided new id already exists"}))
				return
			}
			common.Cities[n].ID = data.ID
		}
		if data.Title != "" {
			common.Cities[n].Title = data.Title
		}
	}

	ret, err := common.GetCityResponse(common.CityResponse{Message: "updated city", Data: data})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Delete city
// @Description Deletes existing city
// @Tags cities
// @Accept json
// @Produce json
// @Param id path string true "city ID"
// @Success 200 {object} common.CityResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /cities/{id} [delete]
func DeleteCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var data any
	if dataCopy, n, err := common.GetOneByID(id, common.TCity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, common.GetBeautyJSONByMap(map[string]any{
			"message": "city not exists",
		}))
		return
	} else {
		data = dataCopy
		common.RemoveSliceItemByID(n, common.TCity)
	}

	ret, err := common.GetCityResponse(common.CityResponse{Message: "deleted city", Data: data.(common.City)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}
