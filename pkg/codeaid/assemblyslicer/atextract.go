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

type ATExtract struct {
}

func (ate *ATExtract) callInstArgs(s string) (string, error) {
	if len(s) < 39 {
		return "", errors.New("pasering assembly file occur error , the line is too short , can't get call instruction args. content: " + s)
	}
	pos := strings.IndexByte(s[39:], '<')
	if pos == -1 {
		return "", errors.New("the call instruction doesn't has callee" + s)
	}
	idx := strings.LastIndex(s, "+0x")
	if idx != -1 {
		return s[pos+40 : idx], nil
	}
	//runtime.panicwrap+0x3a5
	//internal/bytealg.IndexByteString.abi0
	return s[pos+40 : len(s)-1], nil
}

func (ate *ATExtract) functionName(s string) (string, error) {
	if len(s) < 18 {
		return "", errors.New("pasering assembly file occur error , the line is too short , can't get function name. content: " + s)
	}
	return s[18 : len(s)-2], nil
}

func (ate *ATExtract) verb(s string) (string, error) {
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

func (ate *ATExtract) removeleading(scan *bufio.Scanner) {
	for i := 0; i < 6; i++ {
		scan.Scan()
	}
}

func (ate *ATExtract) segmentFeatures(scan *bufio.Scanner) (*flowmapmaker.FuncFeatures, bool) {
	funcfeatures := flowmapmaker.NewFuncFeatures()
	for scan.Scan() {
		line := scan.Text()
		if len(line) < 4 {
			return funcfeatures, false
		}
		ate.parseInstruction(line, funcfeatures)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	return funcfeatures, true
}

func (ate *ATExtract) parseInstruction(inst string, funcfeatures *flowmapmaker.FuncFeatures) {
	if strings.HasPrefix(inst, " ") {
		ate.recordinstructionInfo(inst, funcfeatures)
	} else {
		var err error
		funcfeatures.FuncName, err = ate.functionName(inst)
		if err != nil {
			logger.Info(err.Error())
		}
	}
}

func (ate *ATExtract) recordinstructionInfo(inst string, funcfeatures *flowmapmaker.FuncFeatures) {
	action, err := ate.verb(inst)
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
		callee, err := ate.callInstArgs(inst)
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
