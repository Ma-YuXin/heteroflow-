package assemblyslicer

import (
	"bufio"
	"errors"
	"heterflow/pkg/codeaid/flowmapmaker"
	"heterflow/pkg/codeaid/util"
	"heterflow/pkg/logger"
	"log"
	"strings"
)

type IntelExtract struct {
}

func (ie *IntelExtract) callInstArgs(s string) (string, error) {
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
func (ie *IntelExtract) functionName(s string) (string, error) {
	if len(s) < 18 {
		return "", errors.New("pasering assembly file occur error , the line is too short , can't get function name. content: " + s)
	}
	return s[18 : len(s)-2], nil
}
func (ie *IntelExtract) verb(s string) (string, error) {
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
func (ie *IntelExtract) removeleading(scan *bufio.Scanner) {
	for i := 0; i < 6; i++ {
		scan.Scan()
	}
}

func (ie *IntelExtract) segmentFeatures(scan *bufio.Scanner) (*flowmapmaker.FuncFeatures, bool) {
	funcfeatures := flowmapmaker.NewFuncFeatures()
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

func (ie *IntelExtract) parseInstruction(inst string, funcfeatures *flowmapmaker.FuncFeatures) {
	if strings.HasPrefix(inst, " ") {
		ie.recordinstructionInfo(inst, funcfeatures)
	} else {
		var err error
		funcfeatures.FuncName, err = ie.functionName(inst)
		if err != nil {
			logger.Info(err.Error())
		}
	}
}

func (ie *IntelExtract) recordinstructionInfo(inst string, funcfeatures *flowmapmaker.FuncFeatures) {
	action, err := ie.verb(inst)
	if err != nil {
		logger.Info(err.Error())
		return
	}
	class := util.ActionClassify(action)
	switch class {
	case util.TransmissionInstruction:
		funcfeatures.TransmissionInstruction++
	case util.IOInstruction:
		funcfeatures.IOInstruction++
	case util.ArithmeticInstruction:
		funcfeatures.ArithmeticInstruction++
	case util.LogicalInstruction:
		funcfeatures.LogicalInstruction++
	case util.StringInstruction:
		funcfeatures.StringInstruction++
	case util.ProgramTransferInstruction:
		funcfeatures.ProgramTransferInstructionsInstruction++
		callee, err := ie.callInstArgs(inst)
		if err != nil {
			logger.Info((err.Error()))
			return
		}
		funcfeatures.AddCallee(callee)
	case util.InterruptInstruction:
		funcfeatures.InterruptInstruction++
	case util.PseudoInstruction:
		funcfeatures.PseudoInstruction++
	case util.ProcessorControlInstruction:
		funcfeatures.ProcessorControlInstruction++
	}
}
