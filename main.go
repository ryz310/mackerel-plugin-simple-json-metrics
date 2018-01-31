package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Metrics struct {
	Monitoring []Grouping `json:"monitoring"`
	TimeStump  int        `json:"timestamp"`
}

type Grouping struct {
	Group string `json:"group"`
	Count int    `json:"count"`
}

func main() {
	// 引数から Monitoring API の URL と Access Token を取得する
	title := flag.String("title", "", "Metrics title")
	url := flag.String("url", "", "Monitoring API action")
	api_key := flag.String("api-key", "", "Monitoring API Access Token")
	flag.Parse()

	// Monitoring API に https リクエストする
	req, _ := http.NewRequest("GET", *url, nil)
	req.Header.Set("Authorization", "Bearer "+*api_key)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	// JSONデコード
	var metrics Metrics
	if err := json.Unmarshal(byteArray, &metrics); err != nil {
		log.Fatal(err)
	}

	// Mackerel に投稿するカスタムメトリックを標準出力する
	for _, r := range metrics.Monitoring {
		fmt.Printf("%s.%s\t%d\t%d\n", *title, r.Group, r.Count, metrics.TimeStump)
	}
}
