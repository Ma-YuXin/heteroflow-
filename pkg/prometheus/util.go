package prometheus

import (
	"fmt"
	"strings"
)

type Constraint struct {
	Constraint string
	Value      string
}

func MakeConstraint(values ...string) []Constraint {
	if len(values)%2 != 0 {
		return nil
	}
	constraints := make([]Constraint, 0)
	for i := 0; i < len(values)-1; i += 2 {
		constraint := Constraint{
			Constraint: values[i],
			Value:      values[i+1],
		}
		constraints = append(constraints, constraint)
	}
	return constraints
}

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

func Prom_Metric(metric string, constraint ...Constraint) string {
	if len(constraint) == 0 {
		return metric
	}
	builder := strings.Builder{}
	builder.WriteString(metric)
	builder.WriteByte('{')
	first := constraint[0]
	builder.WriteString(first.Constraint)
	builder.WriteString(`="`)
	builder.WriteString(first.Value)
	builder.WriteByte('"')
	for _, con := range constraint[1:] {
		builder.WriteByte(',')
		builder.WriteString(con.Constraint)
		builder.WriteString(`="`)
		builder.WriteString(con.Value)
		builder.WriteByte('"')
	}
	builder.WriteByte('}')
	return builder.String()
}
