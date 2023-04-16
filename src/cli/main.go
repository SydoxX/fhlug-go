package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

var input = flag.String("i", "points.json", "Input file")
var output = flag.String("o", "output.yaml", "Output file")

func errorHandling(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	flag.Parse()

	jsonFile, err := os.Open(*input)
	errorHandling(err)
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	errorHandling(err)

	var points []Point
	json.Unmarshal(bytes, &points)

	for _, point := range points {
		fmt.Println(point.PrettyPrint())
	}

	yamlData, err := yaml.Marshal(&points)
	errorHandling(err)

	yamlFile, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0600)
	errorHandling(err)

	yamlFile.Write(yamlData)
}
