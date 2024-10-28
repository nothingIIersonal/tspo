package publishingHouses

import (
	"encoding/json"
	"io"
	"net/http"
	"proj/common"

	"github.com/gorilla/mux"
)

func init() {
	common.PublishingHouses = []common.PublishingHouse{
		{ID: "0", Title: "Publishing House 0", Phone: "88005553530", Address: "Address 0", City: "0"},
		{ID: "1", Title: "Publishing House 1", Phone: "88005553531", Address: "Address 1", City: "1"},
		{ID: "2", Title: "Publishing House 2", Phone: "88005553532", Address: "Address 2", City: "2"},
		{ID: "3", Title: "Publishing House 3", Phone: "88005553533", Address: "Address 3", City: "3"},
	}
}

// @Summary Get all publishing houses
// @Description Returns list of publishing houses
// @Tags publishingHouses
// @Produce json
// @Success 200 {object} common.PublishingHousesResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /publishingHouses [get]
func GetPublishingHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ret, err := common.GetPublishingHousesResponse(common.PublishingHousesResponse{Message: "publishing houses list", Data: common.PublishingHouses})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Get one publishing house
// @Description Returns publishing house
// @Tags publishingHouses
// @Produce json
// @Param id path string true "publishing house ID"
// @Success 200 {object} common.PublishingHouseResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /publishingHouses/{id} [get]
func GetPublishingHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	publishingHouse, _, err := common.GetOneByID(id, common.TPublishingHouse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	ret, err := common.GetPublishingHouseResponse(common.PublishingHouseResponse{Message: "publishing house " + id, Data: publishingHouse.(common.PublishingHouse)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Create new publishing house
// @Description Add publishing house to list
// @Tags publishingHouses
// @Accept json
// @Produce json
// @Param publishinghouse body common.PublishingHouse true "publishing house data"
// @Success 201 {object} common.PublishingHouseResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /publishingHouses [post]
func CreatePublishingHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data common.PublishingHouse

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, _, err := common.GetOneByID(data.City, common.TCity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "city with provided id not exists"}))
		return
	}

	if _, _, err := common.GetOneByID(data.ID, common.TPublishingHouse); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "publishing house already exists"}))
		return
	}

	common.PublishingHouses = append(common.PublishingHouses, data)

	ret, err := common.GetPublishingHouseResponse(common.PublishingHouseResponse{Message: "created publishing house", Data: data})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.Writer.Write(w, ret)
}

// @Summary Update publishing house
// @Description Updates existing publishing house
// @Tags publishingHouses
// @Accept json
// @Produce json
// @Param id path string true "publishing house ID"
// @Param publishinghouse body common.PublishingHouse true "publishing house data"
// @Success 200 {object} common.PublishingHouseResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /publishingHouses/{id} [put]
func UpdatePublishingHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data common.PublishingHouse

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	if _, n, err := common.GetOneByID(id, common.TPublishingHouse); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "publishing house not exists"}))
		return
	} else {
		if data.ID != "" {
			if _, _, err := common.GetOneByID(data.ID, common.TPublishingHouse); err == nil {
				w.WriteHeader(http.StatusBadRequest)
				io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "Publishing house with provided new id already exists"}))
				return
			}
			common.PublishingHouses[n].ID = data.ID
		}
		if data.Title != "" {
			common.PublishingHouses[n].Title = data.Title
		}
		if data.Phone != "" {
			common.PublishingHouses[n].Phone = data.Phone
		}
		if data.Address != "" {
			common.PublishingHouses[n].Address = data.Address
		}
		if data.City != "" {
			if _, _, err := common.GetOneByID(data.City, common.TCity); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "city with provided id not exists"}))
				return
			}
			common.PublishingHouses[n].City = data.City
		}
	}

	ret, err := common.GetPublishingHouseResponse(common.PublishingHouseResponse{Message: "updated publishing house", Data: data})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}

// @Summary Delete publishing house
// @Description Deletes existing publishing house
// @Tags publishingHouses
// @Accept json
// @Produce json
// @Param id path string true "publishing house ID"
// @Success 200 {object} common.PublishingHouseResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /publishingHouses/{id} [delete]
func DeletePublishingHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var data any
	if dataCopy, n, err := common.GetOneByID(id, common.TPublishingHouse); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: "publishing house not exists"}))
		return
	} else {
		data = dataCopy
		common.RemoveSliceItemByID(n, common.TPublishingHouse)
	}

	ret, err := common.GetPublishingHouseResponse(common.PublishingHouseResponse{Message: "deleted publishing house", Data: data.(common.PublishingHouse)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.Writer.Write(w, common.GetErrorResponse(common.ErrorResponse{Message: err.Error()}))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Writer.Write(w, ret)
}
