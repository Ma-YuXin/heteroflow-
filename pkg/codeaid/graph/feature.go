package graph

import (
	"heterflow/pkg/logger"
)

type FeatureType int

const (
	ProgrammerFeature FeatureType = iota
	FuncFeatures
	Directedh GraphType = iota
	Undirected
)

type Features interface {
	Features() []int
	Name() string
	Nodes() map[string]Features
	AddInfo(interface{})
}

type GraphType int

type Programmer struct {
	ProgrammerName                   string
	TotalInstruction                 int
	TotalTransmissionInstruction     int
	TotalIOInstruction               int
	TotalArithmeticInstruction       int
	TotalLogicalInstruction          int
	TotalStringInstruction           int
	TotalProgramTransferInstruction  int
	TotalInterruptInstruction        int
	TotalPseudoInstruction           int
	TotalProcessorControlInstruction int
	ControlFlowGraphRoots            map[string]Features
}

type Node struct {
	FuncName                    string
	TotalInstruction            int
	TransmissionInstruction     int
	IOInstruction               int
	ArithmeticInstruction       int
	LogicalInstruction          int
	StringInstruction           int
	ProgramTransferInstruction  int
	InterruptInstruction        int
	PseudoInstruction           int
	ProcessorControlInstruction int
	OtherInstruction            int
	Callee                      map[string]Features
	CalledTimes                 int
	flag                        bool
}

func NewFeatures(class FeatureType) Features {
	switch class {
	case ProgrammerFeature:
		return &Programmer{
			ControlFlowGraphRoots: make(map[string]Features),
		}
	case FuncFeatures:
		return &Node{
			Callee: make(map[string]Features),
		}
	default:
		return nil
	}
}

func (ff *Node) AddCallee(funcName string) {
	ff.Callee[funcName] = nil
}
func (ff *Node) Nodes() map[string]Features {
	return ff.Callee
}
func (ff *Node) Features() []int {
	return []int{ff.TotalInstruction, ff.TransmissionInstruction, ff.IOInstruction, ff.ArithmeticInstruction, ff.LogicalInstruction, ff.StringInstruction, ff.ProgramTransferInstruction, ff.InterruptInstruction, ff.PseudoInstruction, ff.ProcessorControlInstruction}
}

func (ff *Node) Name() string {
	return ff.FuncName
}

func (ff *Node) AddInfo(funcFeatures interface{}) {
	if funcfeat, ok := funcFeatures.(string); ok {
		ff.FuncName = funcfeat
	}
}
func (ff *Programmer) Features() []int {
	return []int{ff.TotalInstruction, ff.TotalTransmissionInstruction, ff.TotalIOInstruction, ff.TotalArithmeticInstruction, ff.TotalLogicalInstruction, ff.TotalStringInstruction, ff.TotalProgramTransferInstruction, ff.TotalInterruptInstruction, ff.TotalPseudoInstruction, ff.TotalProcessorControlInstruction}
}

func (ff *Programmer) Nodes() map[string]Features {
	return ff.ControlFlowGraphRoots
}

func (ff *Programmer) Name() string {
	return ff.ProgrammerName
}

func (ff *Programmer) AddInfo(funcFeatures interface{}) {
	if funcfeat, ok := funcFeatures.(*Node); ok {
		ff.TotalInstruction += funcfeat.TotalInstruction
		ff.TotalTransmissionInstruction += funcfeat.TransmissionInstruction
		ff.TotalIOInstruction += funcfeat.IOInstruction
		ff.TotalArithmeticInstruction += funcfeat.ArithmeticInstruction
		ff.TotalLogicalInstruction += funcfeat.LogicalInstruction
		ff.TotalStringInstruction += funcfeat.StringInstruction
		ff.TotalProgramTransferInstruction += funcfeat.ProgramTransferInstruction
		ff.TotalInterruptInstruction += funcfeat.InterruptInstruction
		ff.TotalPseudoInstruction += funcfeat.PseudoInstruction
		ff.TotalProcessorControlInstruction += funcfeat.ProcessorControlInstruction
	} else if funcfeat, ok := funcFeatures.(string); ok {
		ff.ProgrammerName = funcfeat
	} else if funcfeat, ok := funcFeatures.(*Programmer); ok {
		ff.TotalInstruction += funcfeat.TotalInstruction
		ff.TotalTransmissionInstruction += funcfeat.TotalTransmissionInstruction
		ff.TotalIOInstruction += funcfeat.TotalIOInstruction
		ff.TotalArithmeticInstruction += funcfeat.TotalArithmeticInstruction
		ff.TotalLogicalInstruction += funcfeat.TotalLogicalInstruction
		ff.TotalStringInstruction += funcfeat.TotalStringInstruction
		ff.TotalProgramTransferInstruction += funcfeat.TotalProgramTransferInstruction
		ff.TotalInterruptInstruction += funcfeat.TotalInterruptInstruction
		ff.TotalPseudoInstruction += funcfeat.TotalPseudoInstruction
		ff.TotalProcessorControlInstruction += funcfeat.TotalProcessorControlInstruction
	} else {
		logger.Error("can't add to Programmer Feature,the type is unfit ")
	}
}
