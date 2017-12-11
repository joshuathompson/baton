package api

import (
	"fmt"
	"net/url"
	"strconv"
)

func Search(query, item_types string, limit, offset int) (s searchResults, err error) {
	v := url.Values{}
	v.Add("q", query)
	v.Add("type", item_types)
	v.Add("limit", strconv.Itoa(limit))
	v.Add("offset", strconv.Itoa(offset))

	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"search", v, nil)
	fmt.Println(r.URL.String())
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &s)

	return s, err
}
