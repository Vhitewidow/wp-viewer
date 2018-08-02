package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var formatValues bool

func main() {
	parseFlags()

	response, err := http.Get("http://localhost:3490/v1/unpulsed")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var whatpulse Whatpulse
	json.Unmarshal(responseData, &whatpulse)
	fmt.Println("Keys:", whatpulse.Keys)
	fmt.Println("Clicks:", whatpulse.Clicks)
	if formatValues {
		fmt.Println("Download:", formatDataValue(whatpulse.Download))
		fmt.Println("Upload:", formatDataValue(whatpulse.Upload))
		fmt.Println("Uptime:", formatTimeValue(whatpulse.Uptime))
	} else {
		fmt.Println("Download:", whatpulse.Download)
		fmt.Println("Upload:", whatpulse.Upload)
		fmt.Println("Uptime:", whatpulse.Uptime)
	}

}

func parseFlags() {
	flag.BoolVar(&formatValues, "format", false, "Determine if values printed should be formatted")
	flag.BoolVar(&formatValues, "f", false, "Determine if values printed should be formatted")
	flag.Parse()
}

func formatDataValue(dataValue int) string {
	var value float32
	divisionCount := 0
	var unit string
	value = float32(dataValue)
	for value >= 1024 {
		value = value / 1024
		divisionCount++
	}
	switch divisionCount {
	case 0:
		unit = "B"
	case 1:
		unit = "kB"
	case 2:
		unit = "MB"
	case 3:
		unit = "GB"
	}
	return fmt.Sprintf("%.2f", value) + unit
}

func formatTimeValue(timeValue int) string {
	return fmt.Sprint(time.Duration(timeValue) * time.Second)
}
