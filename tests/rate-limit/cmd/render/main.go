package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

type Record struct {
	Time     time.Time
	Rate     int
	Offset   time.Duration
	Duration time.Duration
	Status   int
}

func main() {
	var name, filename string
	var padding time.Duration
	flag.StringVar(&filename, "f", "", "Filename to open")
	flag.StringVar(&name, "name", "", "Chart name")
	flag.DurationVar(&padding, "p", padding, "Time padding (total duration)")
	flag.Parse()

	log.Println("Rendering", filename, "-", name)

	// Read JSON data from the file
	jsonData, err := readJSONFile(filename, padding)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		os.Exit(1)
	}

	// Extract offset values and calculate seconds
	var xValues []float64
	for _, record := range jsonData {
		xValues = append(xValues, float64(record.Offset)/float64(time.Second))
	}

	// Create a new chart
	graph := chart.Chart{
		TitleStyle: chart.Style{
			Padding:           chart.Box{IsSet: true, Bottom: 50, Top: 0},
			TextLineSpacing:   70,
			TextVerticalAlign: chart.TextVerticalAlignBottom,
		},
		Canvas: chart.Style{
			Padding: chart.Box{IsSet: true, Top: 60},
		},
		XAxis: chart.XAxis{
			Name: "Offset Seconds",
		},
		YAxis: chart.YAxis{
			Name: name,
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: getYValues(jsonData, 200),
			},
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: getYValues(jsonData, 429),
			},
		},
	}

	// Save the chart as a PNG file
	err = graph.Render(chart.PNG, os.Stdout)
	if err != nil {
		fmt.Println("Error rendering chart:", err)
	}
}

// readJSONFile reads JSON data from the specified file
func readJSONFile(filename string, padding time.Duration) ([]Record, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonData []Record
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		return nil, err
	}

	last := jsonData[len(jsonData)-1]
	if last.Offset < padding {
		jsonData = append(jsonData, Record{Time: last.Time.Add(time.Millisecond), Offset: last.Offset + time.Millisecond, Rate: 0, Status: 429})
		left := padding - last.Offset
		jsonData = append(jsonData, Record{Time: last.Time.Add(left), Offset: last.Offset + left, Rate: 0, Status: 429})
	}

	return jsonData, nil
}

// getYValues extracts Rate values from the JSON data
func getYValues(data []Record, status int) []float64 {
	var yValues []float64
	for _, record := range data {
		if record.Status == status {
			yValues = append(yValues, float64(record.Rate))
		} else {
			yValues = append(yValues, float64(0))
		}
	}
	return yValues
}
