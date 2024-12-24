package main

import (
	"fmt"
	"lab2-client/input"
	"lab2-client/req"
	"time"
)

func main() {
	json, ok := req.GetCurrent()
	if !ok {
		return
	}
	launch := int(json["launch"].(float64))
	seconds := int(json["seconds"].(float64))

	start := time.Now()

	input.Read()
	secs := int(time.Since(start).Seconds())
	launch++
	seconds += secs

	fmt.Println(launch, seconds)

	req.Update(launch, seconds)
}
