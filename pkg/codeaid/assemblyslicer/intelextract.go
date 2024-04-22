package assemblyslicer

import (
	"bufio"
	"errors"
	"heterflow/pkg/codeaid/graph"
	"heterflow/pkg/codeaid/tools"
	"heterflow/pkg/logger"
	"log"
	"strings"
)

type IntelExtract struct {
}

func (ie IntelExtract) callInstArgs(s string) (string, error) {
	if len(s) < 39 {
		return "", errors.New("pasering assembly file occur error , the line is too short , can't get call instruction args. content: " + s)
	}
	pos := strings.IndexByte(s[39:], '<')
	if pos == -1 {
		return "", errors.New("the call instruction doesn't has callee" + s)
	}
	idx := strings.LastIndex(s, "+0x")
	if idx != -1 && idx > pos+40 {
		return s[pos+40 : idx], nil
	}
	//runtime.panicwrap+0x3a5
	//internal/bytealg.IndexByteString.abi0
	return s[pos+40 : len(s)-1], nil
}
func (ie IntelExtract) functionName(s string) (string, error) {
	if len(s) < 18 {
		return "", errors.New("pasering assembly file occur error , the line is too short , can't get function name. content: " + s)
	}
	return s[18 : len(s)-2], nil
}
func (ie IntelExtract) verb(s string) (string, error) {
	// defer func(s string) {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("[Panic] ", s)
	// 	}
	// }(s)
	if len(s) < 32 {
		return "", errors.New("pasering assembly file occur error , the line is too short , can't get the action of assembly instruction. content: " + s)
	}
	pos := strings.IndexByte(s[32:], ' ')
	if pos == -1 {
		return s[32:], nil
	}
	return s[32 : pos+32], nil
}
func (ie IntelExtract) removeleading(scan *bufio.Scanner) {
	for i := 0; i < 6; i++ {
		scan.Scan()
	}
}

func (ie IntelExtract) segmentFeatures(scan *bufio.Scanner) (graph.Features, bool) {
	funcfeatures := graph.NewFeatures(graph.FuncFeatures)
	for scan.Scan() {
		line := scan.Text()
		if len(line) < 4 {
			return funcfeatures, false
		}
		ie.parseInstruction(line, funcfeatures)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	return funcfeatures, true
}

func (ie IntelExtract) parseInstruction(inst string, funcfeatures graph.Features) {
	if strings.HasPrefix(inst, " ") {
		ie.recordinstructionInfo(inst, funcfeatures)
	} else {
		name, err := ie.functionName(inst)
		if err != nil {
			logger.Info(err.Error())
		}
		funcfeatures.AddInfo(name)
	}
}

func (ie IntelExtract) recordinstructionInfo(inst string, funcfeatures graph.Features) {
	action, err := ie.verb(inst)
	if err != nil {
		logger.Info(err.Error())
		return
	}
	class := tools.ActionClassify(action)
	if ff, ok := funcfeatures.(*graph.Node); ok {
		ff.TotalInstruction++
		switch class {
		case tools.TransmissionInstruction:
			ff.TransmissionInstruction++
		case tools.IOInstruction:
			ff.IOInstruction++
		case tools.ArithmeticInstruction:
			ff.ArithmeticInstruction++
		case tools.LogicalInstruction:
			ff.LogicalInstruction++
		case tools.StringInstruction:
			ff.StringInstruction++
		case tools.ProgramTransferInstruction:
			ff.ProgramTransferInstruction++
			callee, err := ie.callInstArgs(inst)
			if err != nil {
				logger.Info((err.Error()))
				return
			}
			ff.AddCallee(callee)
		case tools.InterruptInstruction:
			ff.InterruptInstruction++
		case tools.PseudoInstruction:
			ff.PseudoInstruction++
		case tools.ProcessorControlInstruction:
			ff.ProcessorControlInstruction++
		case tools.OtherInstruction:
			ff.OtherInstruction++
		}
	}
}
