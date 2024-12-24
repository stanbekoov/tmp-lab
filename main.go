package main

import (
	"fmt"
	"time"
)

func main() {
	json, ok := GetCurrent()
	if !ok {
		return
	}
	launch := int(json["launch"].(float64))
	seconds := int(json["seconds"].(float64))

	start := time.Now()

	Read()
	secs := int(time.Since(start).Seconds())
	launch++
	seconds += secs

	fmt.Println(launch, seconds)

	Update(launch, seconds)
}
