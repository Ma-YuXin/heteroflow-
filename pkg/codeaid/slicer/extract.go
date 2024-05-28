package slicer

import (
	"bufio"
	"errors"

	"heterflow/pkg/codeaid/cfw"
	"heterflow/pkg/codeaid/def"
	"heterflow/pkg/logger"
	"log"
	"strings"
)

type extracter interface {
	callInstArgs(s string) (string, error)
	functionName(s string) (string, error)
	verb(s string) (string, error)
	removeleading(scan *bufio.Scanner)
	readAsmFunc(scan *bufio.Scanner) (*cfw.Node, bool)
	parseInstruction(inst string, funcfeatures *cfw.Node)
	recordinstructionInfo(inst string, funcfeatures *cfw.Node)
}

type intelExtract struct{}

func (ie intelExtract) callInstArgs(s string) (string, error) {

	pos := strings.Index(s, ":")
	if pos == -1 {
		return "", errors.New("inst is not illegal " + s)
	}
	s = s[pos+1:]
	s = s[23:]
	start := strings.IndexByte(s, '<')
	if start != -1 {
		s = s[start+1:]
	} else {
		return "", errors.New("inst doesn't has callee ")
	}
	end := strings.LastIndexByte(s, '>')
	if end == -1 {
		return "", errors.New("inst doesn't has callee ")
	}
	idx := strings.LastIndex(s, "+0x")
	if idx != -1 {
		end = idx
	}
	//runtime.panicwrap+0x3a5
	//internal/bytealg.IndexByteString.abi0
	return s[:end], nil
}

func (ie intelExtract) functionName(s string) (string, error) {
	_, preProcess, ok := strings.Cut(s, "<")
	if !ok {
		return "", errors.New("not proper function header instruction")
	}
	name, ok := strings.CutSuffix(preProcess, ">:")
	if !ok {
		return "", errors.New("function header instruction format error")
	}
	return name, nil
}

func (ie intelExtract) verb(s string) (string, error) {
	// fmt.Println(s)
	pos := strings.Index(s, ":")
	if pos == -1 {
		return "", errors.New("inst is not illegal " + s)
	}
	s = s[pos+1:]
	if len(s) < 23 {
		return "", errors.New("inst doesn't has verb " + s)
	}
	s = s[23:]
	end := strings.IndexByte(s, ' ')
	if end == -1 {
		end = len(s)
	}
	return s[:end], nil
}

func (ie intelExtract) removeleading(scan *bufio.Scanner) {
	for i := 0; i < 6; i++ {
		scan.Scan()
	}
}

func (ie intelExtract) readAsmFunc(scan *bufio.Scanner) (*cfw.Node, bool) {
	node := cfw.NewNode()
	for scan.Scan() {
		line := scan.Text()
		if len(line) < 4 {
			return node, false
		}
		ie.parseInstruction(line, node)
	}
	if err := scan.Err(); err != nil {
		log.Println(err)
	}
	return node, true
}

func (ie intelExtract) parseInstruction(inst string, funcfeatures *cfw.Node) {
	if strings.HasPrefix(inst, " ") {
		ie.recordinstructionInfo(inst, funcfeatures)
	} else {
		name, err := ie.functionName(inst)
		if err != nil {
			logger.Info(err.Error())
		}
		funcfeatures.SetName(name)
	}
}

func (ie intelExtract) recordinstructionInfo(inst string, node *cfw.Node) {
	action, err := ie.verb(inst)
	if err != nil {
		logger.Info(err.Error())
		return
	}
	if _, ok := def.X86X64ControlFlowInst[action]; ok {
		callee, err := ie.callInstArgs(inst)
		if err != nil {
			logger.Info((err.Error()))
			return
		}
		node.AddCallee(callee)
	}
	node.Record(action)
}
