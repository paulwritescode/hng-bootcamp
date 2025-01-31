package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Slug struct {
	Email            string `json:"email"`
	Current_datetime string `json:"current_datetime"`
	Github_url       string `json:"github_url"`
}

func main() {
	fmt.Println("Server running on port 3000...")
	fmt.Println("http://localhost:3000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json_slug := encodeJson()
		fmt.Fprint(w, string(json_slug))

		// handle CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// status codes
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":3000", nil)
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
