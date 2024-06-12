package tmp
// // PrometheusQueryResponse represents the entire response from Prometheus.
// type PrometheusQueryResponse struct {
// 	Status string              `json:"status"`
// 	Data   PrometheusQueryData `json:"data"`
// }

// // PrometheusQueryData represents the data part of the response.
// type PrometheusQueryData struct {
// 	ResultType string                  `json:"resultType"`
// 	Result     []PrometheusQueryResult `json:"result"`
// }

// // PrometheusQueryResult represents each individual result in the data array.
// type PrometheusQueryResult struct {
// 	Metric map[string]string `json:"metric"`
// 	Value  []interface{}     `json:"value"`
// }

// func Prom_RangeQuery(query, range_option string) PrometheusQueryData {
// 	// Define the Prometheus server URL
// 	prometheusURL := PROM_URL + "/api/v1/query_range"
// 	// Create the final URL with parameters
// 	queryURL := fmt.Sprintf("%s?query=%s%s", prometheusURL, query, range_option)
// 	return dorequest(queryURL)
// }

// func Prom_Query(query, range_option string) PrometheusQueryData {
// 	// Define the Prometheus server URL
// 	prometheusURL := PROM_URL + "/api/v1/query"
// 	// Create the final URL with parameters
// 	queryURL := fmt.Sprintf("%s?query=%s%s", prometheusURL, query, range_option)
// 	return dorequest(queryURL)
// }

// func dorequest(queryURL string) PrometheusQueryData {
// 	fmt.Println(queryURL)
// 	// Make the HTTP request
// 	resp, err := http.Get(queryURL)
// 	if err != nil {
// 		fmt.Printf("Error making HTTP request: %v\n", err)
// 		return PrometheusQueryData{}
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("Error reading response body: %v\n", err)
// 		return PrometheusQueryData{}
// 	}

// 	// Parse the JSON response
// 	var queryResponse PrometheusQueryResponse
// 	if err := json.Unmarshal(body, &queryResponse); err != nil {
// 		fmt.Printf("Error unmarshalling JSON response: %v\n", err)
// 		return PrometheusQueryData{}
// 	}

// 	// Check if the query was successful
// 	if queryResponse.Status != "success" {
// 		fmt.Printf("Query failed: %s\n", queryResponse.Status)
// 		return PrometheusQueryData{}
// 	}
// 	return queryResponse.Data
// }
