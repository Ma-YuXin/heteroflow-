package cfw

import (
	"fmt"
	"heterflow/pkg/codeaid/def"
)

type techInstCounter struct {
	Total          int
	VIRTUALIZATION int
	GP             int
	GP_EXT         int
	GP_IN_OUT      int
	FPU            int
	MMX            int
	STATE          int
	SIMD           int
	SSE            int
	SCALAR         int
	CRYPTO_HASH    int
	AVX            int
	AVX512         int
	MASK           int
	AMX            int
}

type behaviorInstCounter struct {
	Total            int
	Transmission     int
	IO               int
	Arithmetic       int
	Logical          int
	String           int
	ProgramTransfer  int
	Interrupt        int
	Pseudo           int
	ProcessorControl int
}
type ProgramFeatures struct {
	Counter        techInstCounter
	ProgrammerName string
	// ControlFlowGraphRoots map[Cstring]*Node
}

type Node struct {
	Counter     techInstCounter
	FuncName    string
	Callee      map[string]struct{}
	CalledTimes int
	OutDegree   int
	flag        bool
}

func NewNode() *Node {
	return &Node{
		Callee: make(map[string]struct{}),
	}
}

func NewProgramFeatures() *ProgramFeatures {
	return &ProgramFeatures{
		// ControlFlowGraphRoots: make(map[string]*Node),
	}
}

func (bic *behaviorInstCounter) get() []int {
	return []int{
		bic.Total,
		bic.Transmission,
		bic.IO,
		bic.Arithmetic,
		bic.Logical,
		bic.String,
		bic.ProgramTransfer,
		bic.Interrupt,
		bic.Pseudo,
		bic.ProcessorControl,
	}
}

func (bic *behaviorInstCounter) degmentsPerInterval() []int {
	return def.BehaviorInstCounterHierarchies
}

func (bic *behaviorInstCounter) deepCopy() behaviorInstCounter {
	return behaviorInstCounter{
		Total:            bic.Total,
		Transmission:     bic.Transmission,
		IO:               bic.IO,
		Arithmetic:       bic.Arithmetic,
		Logical:          bic.Logical,
		String:           bic.String,
		ProgramTransfer:  bic.ProgramTransfer,
		Interrupt:        bic.Interrupt,
		Pseudo:           bic.Pseudo,
		ProcessorControl: bic.ProcessorControl,
	}
}

func (bic *behaviorInstCounter) add(in *behaviorInstCounter) {
	bic.Total += in.Total
	bic.Transmission += in.Transmission
	bic.IO += in.IO
	bic.Arithmetic += in.Arithmetic
	bic.Logical += in.Logical
	bic.String += in.String
	bic.ProgramTransfer += in.ProgramTransfer
	bic.Interrupt += in.Interrupt
	bic.Pseudo += in.Pseudo
	bic.ProcessorControl += in.ProcessorControl
}

func (bic *behaviorInstCounter) classification(action string) def.BehaviorInstManager {
	if _, ok := def.X86X64ProcessorActionOrientedSet[action]; !ok {
		fmt.Println("missed action", action)
	}
	return def.X86X64ProcessorActionOrientedSet[action]
}

func (bic *behaviorInstCounter) record(action string) {
	class := bic.classification(action)
	bic.Total++
	for _, v := range class {
		switch v {
		case def.Transmission:
			bic.Transmission++
		case def.IO:
			bic.IO++
		case def.Arithmetic:
			bic.Arithmetic++
		case def.Logical:
			bic.Logical++
		case def.String:
			bic.String++
		case def.ProgramTransfer:
			bic.ProgramTransfer++
		case def.Interrupt:
			bic.Interrupt++
		case def.Pseudo:
			bic.Pseudo++
		case def.ProcessorControl:
			bic.ProcessorControl++
		}
	}
}

func (tic *techInstCounter) get() []int {
	return []int{
		tic.Total,
		tic.VIRTUALIZATION,
		tic.GP,
		tic.GP_EXT,
		tic.GP_IN_OUT,
		tic.FPU,
		tic.MMX,
		tic.STATE,
		tic.SIMD,
		tic.SSE,
		tic.SCALAR,
		tic.CRYPTO_HASH,
		tic.AVX,
		tic.AVX512,
		tic.MASK,
		tic.AMX,
	}
}

func (tic *techInstCounter) segmentsPerInterval() []int {
	return def.TechInstCounterHierarchies
}

func (tic *techInstCounter) deepCopy() techInstCounter {
	return techInstCounter{
		Total:          tic.Total,
		VIRTUALIZATION: tic.VIRTUALIZATION,
		GP:             tic.GP,
		GP_EXT:         tic.GP_EXT,
		GP_IN_OUT:      tic.GP_IN_OUT,
		FPU:            tic.FPU,
		MMX:            tic.MMX,
		STATE:          tic.STATE,
		SIMD:           tic.SIMD,
		SSE:            tic.SSE,
		SCALAR:         tic.SCALAR,
		CRYPTO_HASH:    tic.CRYPTO_HASH,
		AVX:            tic.AVX,
		AVX512:         tic.AVX512,
		MASK:           tic.MASK,
		AMX:            tic.AMX,
	}
}

func (tic *techInstCounter) add(in *techInstCounter) {
	tic.Total += in.Total
	tic.VIRTUALIZATION += in.VIRTUALIZATION
	tic.GP += in.GP
	tic.GP_EXT += in.GP_EXT
	tic.GP_IN_OUT += in.GP_IN_OUT
	tic.FPU += in.FPU
	tic.MMX += in.MMX
	tic.STATE += in.STATE
	tic.SIMD += in.SIMD
	tic.SSE += in.SSE
	tic.SCALAR += in.SCALAR
	tic.CRYPTO_HASH += in.CRYPTO_HASH
	tic.AVX += in.AVX
	tic.AVX512 += in.AVX512
	tic.MASK += in.MASK
	tic.AMX += in.AMX
}

func (tic *techInstCounter) classification(action string) def.TechInstManager {
	if _, ok := def.X86X64HardwareFeatureInstSet[action]; !ok {
		fmt.Println("missed action", action)
	}
	return def.X86X64HardwareFeatureInstSet[action]
}

func (tic *techInstCounter) record(action string) {
	class := tic.classification(action)
	tic.Total++
	for _, v := range class {
		switch v {
		case def.VIRTUALIZATION:
			tic.VIRTUALIZATION++
		case def.GP:
			tic.GP++
		case def.GP_EXT:
			tic.GP_EXT++
		case def.GP_IN_OUT:
			tic.GP_IN_OUT++
		case def.FPU:
			tic.FPU++
		case def.MMX:
			tic.MMX++
		case def.STATE:
			tic.STATE++
		case def.SIMD:
			tic.SIMD++
		case def.SSE:
			tic.SSE++
		case def.SCALAR:
			tic.SCALAR++
		case def.CRYPTO_HASH:
			tic.CRYPTO_HASH++
		case def.AVX:
			tic.AVX++
		case def.AVX512:
			tic.AVX512++
		case def.MASK:
			tic.MASK++
		case def.AMX:
			tic.AMX++
		}
	}
}

func (node *Node) AddCallee(funcName string) {
	node.Callee[funcName] = struct{}{}
}

func (node *Node) Neighbors() map[string]struct{} {
	return node.Callee
}

func (node *Node) Features() []int {
	return append(node.Counter.get(), node.CalledTimes, node.CalledTimes)
}

func (node *Node) SegmentsPerInterval() []int {
	return append(node.Counter.segmentsPerInterval(), 4, 4)
}

func (node *Node) Name() string {
	return node.FuncName
}

func (node *Node) MetricsNumber() int {
	return len(node.SegmentsPerInterval())
}

func (node *Node) DeepCopy() *Node {
	return &Node{
		FuncName:    node.FuncName,
		Counter:     node.Counter.deepCopy(),
		Callee:      node.Callee,
		CalledTimes: node.CalledTimes,
		flag:        node.flag,
		OutDegree:   node.CalledTimes,
	}
}

func (node *Node) Add(in *Node) {
	node.Counter.add(&in.Counter)
	node.CalledTimes += in.CalledTimes
	node.CalledTimes += in.CalledTimes
}

func (node *Node) Record(action string) {
	node.Counter.record(action)
}

func (node *Node) SetName(name string) {
	node.FuncName = name
}

func (profeature *ProgramFeatures) Features() []int {
	return profeature.Counter.get()
}
func (profeature *ProgramFeatures) FeaturesWeight() []float64 {
	return []float64{
		16, //tic.Total,
		8,  //tic.VIRTUALIZATION,
		1,  //tic.GP,
		1,  //tic.GP_EXT,
		8,  //tic.GP_IN_OUT,
		4,  //tic.FPU,
		8,  //tic.MMX,
		8,  //tic.STATE,
		8,  //tic.SIMD,
		8,  //tic.SSE,
		8,  //tic.SCALAR,
		8,  //tic.CRYPTO_HASH,
		8,  //tic.AVX,
		8,  //tic.AVX512,
		8,  //tic.MASK,
		8,  //tic.AMX,
	}
}

func (profeature *ProgramFeatures) Name() string {
	return profeature.ProgrammerName
}

func (profeature *ProgramFeatures) Add(in *Node) {
	profeature.Counter.add(&in.Counter)
}

func (profeature *ProgramFeatures) SetName(name string) {
	profeature.ProgrammerName = name
}

func (profeature *ProgramFeatures) DeepCopy() ProgramFeatures {
	return ProgramFeatures{
		ProgrammerName: profeature.ProgrammerName,
		Counter:        profeature.Counter.deepCopy(),
	}
}
