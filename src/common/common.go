package common

import (
	"encoding/json"
	"errors"
)

type City struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type CityResponse struct {
	Message string `json:"message"`
	Data    City   `json:"data"`
}
type CitiesResponse struct {
	Message string `json:"message"`
	Data    []City `json:"data"`
}

func GetCityResponse(cityResp CityResponse) ([]byte, error) {
	return json.Marshal(cityResp)
}

func GetCitiesResponse(citiesResp CitiesResponse) ([]byte, error) {
	return json.Marshal(citiesResp)
}

type PublishingHouse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	City    string `json:"city"`
}

type PublishingHouseResponse struct {
	Message string          `json:"message"`
	Data    PublishingHouse `json:"data"`
}

type PublishingHousesResponse struct {
	Message string            `json:"message"`
	Data    []PublishingHouse `json:"data"`
}

func GetPublishingHouseResponse(publishingHouseResp PublishingHouseResponse) ([]byte, error) {
	return json.Marshal(publishingHouseResp)
}

func GetPublishingHousesResponse(publishingHousesResp PublishingHousesResponse) ([]byte, error) {
	return json.Marshal(publishingHousesResp)
}

type Genre struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type GenreResponse struct {
	Message string `json:"message"`
	Data    Genre  `json:"data"`
}

type GenresResponse struct {
	Message string  `json:"message"`
	Data    []Genre `json:"data"`
}

func GetGenreResponse(genreResp GenreResponse) ([]byte, error) {
	return json.Marshal(genreResp)
}

func GetGenresResponse(genresResp GenresResponse) ([]byte, error) {
	return json.Marshal(genresResp)
}

type Book struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	ISBN            string `json:"isbn"`
	PublishingHouse string `json:"publishinghouse"`
	Genre           string `json:"genre"`
}

type BookResponse struct {
	Message string `json:"message"`
	Data    Book   `json:"data"`
}

type BooksResponse struct {
	Message string `json:"message"`
	Data    []Book `json:"data"`
}

func GetBookResponse(bookResp BookResponse) ([]byte, error) {
	return json.Marshal(bookResp)
}

func GetBooksResponse(booksResp BooksResponse) ([]byte, error) {
	return json.Marshal(booksResp)
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetErrorResponse(errorResp ErrorResponse) []byte {
	ret, _ := json.Marshal(errorResp)
	return ret
}

type TType string

const (
	TCity            TType = "city"
	TPublishingHouse TType = "publishingHouse"
	TGenre           TType = "genre"
	TBook            TType = "book"
)

var Cities []City
var PublishingHouses []PublishingHouse
var Genres []Genre
var Books []Book

func RemoveSliceItemByID(index int, typ TType) {
	switch typ {
	case TCity:
		Cities = append(Cities[:index], Cities[index+1:]...)
	case TPublishingHouse:
		PublishingHouses = append(PublishingHouses[:index], PublishingHouses[index+1:]...)
	case TGenre:
		Genres = append(Genres[:index], Genres[index+1:]...)
	case TBook:
		Books = append(Books[:index], Books[index+1:]...)
	}
}

func GetOneByID(id string, typ TType) (any, int, error) {
	switch typ {
	case TCity:
		for n, item := range Cities {
			if item.ID == id {
				return item, n, nil
			}
		}
	case TPublishingHouse:
		for n, item := range PublishingHouses {
			if item.ID == id {
				return item, n, nil
			}
		}
	case TGenre:
		for n, item := range Genres {
			if item.ID == id {
				return item, n, nil
			}
		}
	case TBook:
		for n, item := range Books {
			if item.ID == id {
				return item, n, nil
			}
		}
	}

	return nil, 0, errors.New("can't get item")
}

// TODO: delete all items which dependece on deleted item
func RemoveAllDependendecies(index int, typ TType) {
	// switch typ {
	// case TCity:
	// 	Cities = append(Cities[:index], Cities[index+1:]...)
	// case TPublishingHouse:
	// 	PublishingHouses = append(PublishingHouses[:index], PublishingHouses[index+1:]...)
	// case TGenre:
	// 	Genres = append(Genres[:index], Genres[index+1:]...)
	// case TBook:
	// 	Books = append(Books[:index], Books[index+1:]...)
	// }
}

// helper JSON functions
func GetJSONByMap(data map[string]any) []byte {
	res, _ := json.MarshalIndent(data, "", "\t")
	return res
}

func GetBeautyJSONByMap(data map[string]any) string {
	return string(GetJSONByMap(data))
}

func GetBeautyJSONByBytes(data []byte) string {
	return string(data)
}

//// end helper JSON funtions
