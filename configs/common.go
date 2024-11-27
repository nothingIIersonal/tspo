package configs

import (
	"net/url"
	"pr8_1/dtos"
	"strconv"
	"strings"
)

func parseQuery(query *url.Values, limit *int, page *int, sort *string, searchs *[]dtos.Search) {
	for key, value := range *query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			*limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			*page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			*sort = queryValue
			break
		}

		if strings.Contains(key, ".") {
			searchKeys := strings.Split(key, ".")
			search := dtos.Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}
			*searchs = append(*searchs, search)
		}
	}
}
