package flowmapmaker

type FuncFeatures struct {
	FuncName                               string
	TransmissionInstruction                int
	IOInstruction                          int
	ArithmeticInstruction                  int
	LogicalInstruction                     int
	StringInstruction                      int
	ProgramTransferInstructionsInstruction int
	InterruptInstruction                   int
	PseudoInstruction                      int
	ProcessorControlInstruction            int
	Callee                                 map[string]*FuncFeatures
}

type FileFeatures struct {
	ProgrammerName                              string
	TotalTransmissionInstruction                int
	TotalIOInstruction                          int
	TotalArithmeticInstruction                  int
	TotalLogicalInstruction                     int
	TotalStringInstruction                      int
	TotalProgramTransferInstructionsInstruction int
	TotalInterruptInstruction                   int
	TotalPseudoInstruction                      int
	TotalProcessorControlInstruction            int
	ControlFlowGraphRoots                       map[string]*FuncFeatures
}

func NewFuncFeatures() *FuncFeatures {
	return &FuncFeatures{
		Callee: make(map[string]*FuncFeatures),
	}
}
func NewFileFeatures() *FileFeatures {
	return &FileFeatures{
		ControlFlowGraphRoots: make(map[string]*FuncFeatures),
	}
}
func (ff *FuncFeatures) AddCallee(funcName string) {
	ff.Callee[funcName] = nil
}

func (ff *FileFeatures) AddFileFeatures(funcFeatures *FuncFeatures) {
	ff.TotalTransmissionInstruction += funcFeatures.TransmissionInstruction
	ff.TotalIOInstruction += funcFeatures.IOInstruction
	ff.TotalArithmeticInstruction += funcFeatures.ArithmeticInstruction
	ff.TotalLogicalInstruction += funcFeatures.LogicalInstruction
	ff.TotalStringInstruction += funcFeatures.StringInstruction
	ff.TotalProgramTransferInstructionsInstruction += funcFeatures.ProgramTransferInstructionsInstruction
	ff.TotalInterruptInstruction += funcFeatures.InterruptInstruction
	ff.TotalPseudoInstruction += funcFeatures.PseudoInstruction
	ff.TotalProcessorControlInstruction += funcFeatures.ProcessorControlInstruction
}
