package prometheus

import (
	"fmt"
)

func Prom_Rate(metric, duration string) string {
	return fmt.Sprintf("rate(%s[%s])", metric, duration)
}

func Prom_Topk(metric string, k int) string {
	return fmt.Sprintf("topk(%d,%s)", k, metric)
}

func Prom_Sum(metric string) string {
	return fmt.Sprintf("sum(%s)", metric)
}

func Prom_Min(metric string) string {
	return fmt.Sprintf("min(%s)", metric)
}

func Prom_Max(metric string) string {
	return fmt.Sprintf("max(%s)", metric)
}

func Prom_Avg(metric string) string {
	return fmt.Sprintf("avg(%s)", metric)
}

func Prom_Metric(metric, constraint, value string) string {
	if constraint == "" {
		return metric
	}
	return fmt.Sprintf(`%s{%s="%s"}`, metric, constraint, value)
}
