package main

import (
	"encoding/json"
	"fmt"
	botgolang "github.com/gldmtr/bot-golang"
	"os"
)

type eventsResponse struct {
	OK     bool               `json:"ok"`
	Events []*botgolang.Event `json:"events"`
}

func main() {
	data, err := os.ReadFile("./test.json")
	if err != nil {
		panic(err)
	}

	response := eventsResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Events[0])
}
