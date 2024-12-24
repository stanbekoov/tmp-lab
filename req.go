package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

const API_URL = "http://localhost:8000"

var MAC = ""

func init() {
	ifas, err := net.Interfaces()
	if err != nil {
		log.Fatal(err.Error())
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}

	MAC = as[0]
}

func GetCurrent() (map[string]any, bool) {
	url := API_URL + "/current?mac=" + MAC

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	jsonMap := make(map[string]any)
	json.Unmarshal(body, &jsonMap)

	if response.StatusCode != 200 {
		if jsonMap["message"] == "buy premium" {
			fmt.Println("your trial period has come to an end\n Time to buy premium ;)")
			return nil, false
		}
		log.Fatal()
	}

	return jsonMap, true
}

func Update(launch, secs int) {
	jsonMap := make(map[string]int)
	jsonMap["launch"] = launch
	jsonMap["seconds"] = secs

	body, err := json.Marshal(jsonMap)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(body)

	url := API_URL + "/update?mac=" + MAC

	request, err := http.NewRequest(http.MethodPatch, url, reader)
	if err != nil {
		log.Fatal(err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}

	if response.StatusCode != 200 {
		fmt.Println("someting went wrong")
	}
}
