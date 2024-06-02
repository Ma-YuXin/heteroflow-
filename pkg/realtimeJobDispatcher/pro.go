package realtimejobdispatcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// PrometheusResponse represents the structure of the response from Prometheus.
type PrometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// queryPrometheus executes a query against a Prometheus server.
func queryPrometheus(prometheusURL, query string) (*PrometheusResponse, error) {
	// Construct the query URL.
	u, err := url.Parse(prometheusURL)
	if err != nil {
		return nil, fmt.Errorf("invalid Prometheus URL: %v", err)
	}
	u.Path = "/api/v1/query"
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()

	// Send the HTTP request.
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("error querying Prometheus: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var promResp PrometheusResponse
	if err := json.Unmarshal(body, &promResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	if promResp.Status != "success" {
		return nil, fmt.Errorf("Prometheus query failed: %s", promResp.Status)
	}

	return &promResp, nil
}

func main() {
	// Example usage of the queryPrometheus function.
	prometheusURL := "http://localhost:9090"
	query := "up"

	resp, err := queryPrometheus(prometheusURL, query)
	if err != nil {
		log.Fatalf("Error querying Prometheus: %v", err)
	}

	// Print the query results.
	for _, result := range resp.Data.Result {
		fmt.Printf("Metric: %v\n", result.Metric)
		fmt.Printf("Value: %v\n", result.Value)
	}
}
