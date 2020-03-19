
package main

import (
"bufio"
"encoding/json"
"fmt"
"net/http"
"os"
"strconv"
"time"
)

func main() {
	result := getUsernames(10)
	stdout, _ := os.Create(os.Getenv("OUTPUT_PATH"))
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	for i, resultItem := range result {
		fmt.Println(resultItem)
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}
	writer.Flush()
}
func getUsernames(threshold int32) []string {
	var userNames []string

	// Store all the page info
	pageInfoList := make([]*PageInfo, 0)

	// Request the first page and get the total page number
	// Based on the total page number decide the next action
	pageInfo, err := client(1)
	if err != nil {
		return nil
	}
	pageInfoList = append(pageInfoList, pageInfo)
	for i := 1; i < pageInfo.TotalPages; i++ {

		pageInfo, _ := client(i + 1)
		pageInfoList = append(pageInfoList, pageInfo)
	}

	for _, v := range pageInfoList {
		x := *v
		y := x.Data
		for _, value := range y {
			if value.SubmissionCount > threshold {
				userNames = append(userNames, value.UserName)
			}
		}
	}
	println(userNames)
	return userNames
}

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





