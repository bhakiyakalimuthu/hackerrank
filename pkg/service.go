package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func client(pageNumber int) (*PageInfo, error) {
	hostName := "https://jsonmock.hackerrank.com/api/article_users/search?page=" + strconv.Itoa(pageNumber)
	req, err := http.NewRequest(http.MethodGet, hostName, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var page PageInfo
	err = json.NewDecoder(resp.Body).Decode(&page)
	if err != nil {
		return nil, err
	}
	return &page, nil
}