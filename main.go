
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/bhakiyakalimuthu/hackerrank/internal/user"
)

type config struct{
	BaseURL string `env:"base_url" envDefault:"https://jsonmock.hackerrank.com/api/article_users/search?page=" `
}

func main() {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	s:= user.NewServiceDefault(client)
	result := s.GetUserName(10)
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







