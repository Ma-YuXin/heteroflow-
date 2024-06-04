package prometheus

import (
	"fmt"

	"context"

	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func QueryPrometheus(query string) model.Vector {
	// 创建Prometheus API客户端等...
	client, _ := api.NewClient(api.Config{
		Address: PROM_URL,
	})
	v1api := v1.NewAPI(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, warnings, err := v1api.Query(ctx, query, time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		return nil
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	vectorVal, ok := result.(model.Vector)
	if !ok {
		fmt.Println("Query result is not a Vector")
		return nil
	}

	return vectorVal
}

func QueryRangePrometheus(query string, start, end time.Time, step time.Duration) model.Matrix {
	// 创建Prometheus API客户端
	client, _ := api.NewClient(api.Config{
		Address: PROM_URL, // 这里指明了要连接的Prometheus的地址
	})
	v1api := v1.NewAPI(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r := v1.Range{
		Start: start,
		End:   end,
		Step:  step,
	}
	result, warnings, err := v1api.QueryRange(ctx, query, r)
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		return nil
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	matrixVal, ok := result.(model.Matrix)
	if !ok {
		fmt.Println("Query result is not a Matrix")
		return nil
	}
	return matrixVal
}

func PrintRes(result model.Value) {
	switch v := result.(type) {
	case model.Vector:
		// 处理向量类型的结果
		fmt.Println("Result is a Vector:", v)
		for _, sample := range v {
			timestamp := sample.Timestamp.Time() // 转换为标准的time.Time类型
			value := float64(sample.Value)       // 转换为浮点数
			fmt.Printf("At time %v, value is: %v\n", timestamp, value)
		}
	case model.Matrix:
		// 处理矩阵类型的结果
		fmt.Println("Result is a Matrix:", v)
		for _, stream := range v {
			fmt.Printf("Metric: %s\n", stream.Metric)
			for _, sample := range stream.Values {
				timestamp := sample.Timestamp.Time()
				value := float64(sample.Value)
				fmt.Printf("  At time %v, value is: %v\n", timestamp, value)
			}
		}
	default:
		// 未知类型
		fmt.Println("Unknown type of result")
	}
}

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
