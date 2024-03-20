package assemblyslicer

import (
	"bufio"
	"heterflow/pkg/codeaid/graph"
)

type Extract interface {
	callInstArgs(s string) (string, error)
	functionName(s string) (string, error)
	verb(s string) (string, error)
	removeleading(scan *bufio.Scanner)
	segmentFeatures(scan *bufio.Scanner) (graph.Features, bool)
	parseInstruction(inst string, funcfeatures graph.Features)
	recordinstructionInfo(inst string, funcfeatures graph.Features)
}
