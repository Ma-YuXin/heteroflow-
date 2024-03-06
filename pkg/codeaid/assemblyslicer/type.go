package assemblyslicer

import (
	"bufio"
	"heterflow/pkg/codeaid/flowmapmaker"
)

type Extract interface {
	callInstArgs(s string) (string, error)
	functionName(s string) (string, error)
	verb(s string) (string, error)
	removeleading(scan *bufio.Scanner)
	segmentFeatures(scan *bufio.Scanner)(*flowmapmaker.FuncFeatures, bool)
	parseInstruction(inst string, funcfeatures *flowmapmaker.FuncFeatures)
	recordinstructionInfo(inst string, funcfeatures *flowmapmaker.FuncFeatures)
}
