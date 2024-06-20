package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func FetchBenchmarkResult(file string) {
	// Open the JSON file
	jsonFile, err := os.Open("/mnt/data/nfs/myx/helloworld/cfp2017-results-20240614-040337.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// Read the file's content
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a slice of Result
	var results []BenchmarkResult
	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		log.Fatal(err)
	}

	// Print the results
	for _, result := range results {
		fmt.Printf("Benchmark: %s\n", result.Benchmark)
		fmt.Printf("Hardware Vendor: %s\n", result.HardwareVendor)
		fmt.Printf("System: %s\n", result.System)
		fmt.Printf("Processor: %s\n", result.Processor)
		fmt.Printf("Result: %s\n", result.Result)
		fmt.Println()
	}
}
