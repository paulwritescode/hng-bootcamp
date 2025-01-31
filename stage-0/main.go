package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Slug struct {
	Email            string `json:"email"`
	Current_datetime string `json:"current_datetime"`
	Github_url       string `json:"github_url"`
}

func main() {
	json_slug := encodeJson()
	fmt.Println(string(json_slug))
}

func encodeJson() []byte {
	this_time := time.Now().Format("2006-01-02T15:04:05Z07:00")

	json_slug := Slug{
		Email:            "kinyattipaul@gmail.com",
		Current_datetime: this_time,
		Github_url:       "https://github.com/paulwritescode/hng-bootcamp",
	}

	finJson, err := json.MarshalIndent(json_slug, "", "  ")
	if err != nil {
		panic(err)
	}

	return finJson
}
