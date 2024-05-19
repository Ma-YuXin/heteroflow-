package assemblyslicer

import (
	"bufio"
	"errors"
	"heterflow/pkg/codeaid/graph"
	"heterflow/pkg/logger"
	"log"
	"strings"
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

type IntelExtract struct{}

type Class int8

var (
	UsedAction      = map[string]struct{}{}
	MissedAction    = map[string]struct{}{}
	X86Instrcutions = map[string]Class{"mov": TransmissionInstruction, "movsxd": TransmissionInstruction,
		"movsx": TransmissionInstruction, "movzx": TransmissionInstruction, "push": TransmissionInstruction, "pop": TransmissionInstruction,
		"pusha": TransmissionInstruction, "popa": TransmissionInstruction, "popad": TransmissionInstruction,
		"bswap": TransmissionInstruction, "xchg": TransmissionInstruction, "cmpxchg": TransmissionInstruction,
		"xadd": TransmissionInstruction, "xlat": TransmissionInstruction, "in": IOInstruction, "out": IOInstruction,
		"lea": TransmissionInstruction, "lds": TransmissionInstruction, "les": TransmissionInstruction,
		"lfs": TransmissionInstruction, "lgs": TransmissionInstruction, "lss": TransmissionInstruction,
		"lahf": TransmissionInstruction, "sahf": TransmissionInstruction, "pushf": TransmissionInstruction,
		"popf": TransmissionInstruction, "popd": TransmissionInstruction, "pushd": TransmissionInstruction,
		"add": ArithmeticInstruction, "adc": ArithmeticInstruction, "inc": ArithmeticInstruction, "aaa": ArithmeticInstruction,
		"daa": ArithmeticInstruction, "sub": ArithmeticInstruction, "sbb": ArithmeticInstruction,
		"dec": ArithmeticInstruction, "nec": ArithmeticInstruction, "cmp": ArithmeticInstruction, "aas": ArithmeticInstruction,
		"das": ArithmeticInstruction, "mul": ArithmeticInstruction, "imul": ArithmeticInstruction,
		"div": ArithmeticInstruction, "idiv": ArithmeticInstruction, "aam": ArithmeticInstruction,
		"aad": ArithmeticInstruction, "cbw": ArithmeticInstruction, "cwd": ArithmeticInstruction, "cwde": ArithmeticInstruction,
		"cdq": ArithmeticInstruction, "and": LogicalInstruction, "or": LogicalInstruction, "xor": LogicalInstruction,
		"not": LogicalInstruction, "test": LogicalInstruction, "shl": LogicalInstruction, "sal": LogicalInstruction,
		"shr": LogicalInstruction, "sar": LogicalInstruction, "rol": LogicalInstruction, "ror": LogicalInstruction,
		"rcl": LogicalInstruction, "rcr": LogicalInstruction, "ds": StringInstruction, "es": StringInstruction,
		"cx": StringInstruction, "al": StringInstruction, "ax": StringInstruction, "movs": StringInstruction,
		"movsb": StringInstruction, "movsw": StringInstruction, "movsd": StringInstruction, "cmps": StringInstruction,
		"cmpsb": StringInstruction, "cmpsw": StringInstruction, "scas": StringInstruction, "lods": StringInstruction,
		"lodsb": StringInstruction, "lodsw": StringInstruction, "lodsd": StringInstruction, "stos": StringInstruction,
		"rep": StringInstruction, "repe": StringInstruction, "repne": StringInstruction, "repc": StringInstruction,
		"repnc": StringInstruction, "int": InterruptInstruction, "into": InterruptInstruction, "iret": InterruptInstruction,
		"hlt": InterruptInstruction, "wait": InterruptInstruction, "esc": InterruptInstruction, "lock": InterruptInstruction,
		"nop": InterruptInstruction, "stc": InterruptInstruction, "clc": InterruptInstruction, "cmc": InterruptInstruction,
		"std": InterruptInstruction, "cld": InterruptInstruction, "sti": InterruptInstruction, "cli": InterruptInstruction,
		"dw": PseudoInstruction, "proc": PseudoInstruction, "endp": PseudoInstruction, "segment": PseudoInstruction,
		"assume": PseudoInstruction, "ends": PseudoInstruction, "end": PseudoInstruction, "loop": ProcessorControlInstruction,
		"loope": ProcessorControlInstruction, "loopz": ProcessorControlInstruction, "loopne": ProcessorControlInstruction,
		"loopnz": ProcessorControlInstruction, "jcxz": ProcessorControlInstruction, "jecxz": ProcessorControlInstruction,
		"jmp": ProgramTransferInstruction, "call": ProgramTransferInstruction, "ret": ProgramTransferInstruction,
		"retf": ProgramTransferInstruction, "jae": ProgramTransferInstruction, "jnb": ProgramTransferInstruction,
		"jb": ProgramTransferInstruction, "jnae": ProgramTransferInstruction, "jbe": ProgramTransferInstruction,
		"jna": ProgramTransferInstruction, "jg": ProgramTransferInstruction, "jnle": ProgramTransferInstruction,
		"jge": ProgramTransferInstruction, "jnl": ProgramTransferInstruction, "jl": ProgramTransferInstruction,
		"jnge": ProgramTransferInstruction, "jle": ProgramTransferInstruction, "jng": ProgramTransferInstruction,
		"je": ProgramTransferInstruction, "jz": ProgramTransferInstruction, "jne": ProgramTransferInstruction,
		"jnz": ProgramTransferInstruction, "jc": ProgramTransferInstruction, "jnc": ProgramTransferInstruction,
		"jno": ProgramTransferInstruction, "jnp": ProgramTransferInstruction, "jpo": ProgramTransferInstruction,
		"jns": ProgramTransferInstruction, "jo": ProgramTransferInstruction, "jp": ProgramTransferInstruction,
		"jpe": ProgramTransferInstruction, "js": ProgramTransferInstruction, "adcx": ArithmeticInstruction,
		"addl": ArithmeticInstruction, "addq": ArithmeticInstruction, "addsd": ArithmeticInstruction,
		"addss": ArithmeticInstruction, "adox": ArithmeticInstruction, "andn": LogicalInstruction,
		"andpd": LogicalInstruction, "andq": LogicalInstruction, "bt": ArithmeticInstruction,
		"btc": ArithmeticInstruction, "btr": ArithmeticInstruction, "bts": ArithmeticInstruction,
		"cltd": ArithmeticInstruction, "cmpb": ArithmeticInstruction, "cmpl": ArithmeticInstruction,
		"cmpnltsd": ArithmeticInstruction, "cmpq": ArithmeticInstruction, "cmpw": ArithmeticInstruction,
		"comisd": ArithmeticInstruction, "cqto": ArithmeticInstruction, "cvtsd2si": ArithmeticInstruction,
		"cvtsi2sd": ArithmeticInstruction, "cvtsi2ss": ArithmeticInstruction, "cvtss2sd": ArithmeticInstruction,
		"decl": ArithmeticInstruction, "decq": ArithmeticInstruction, "divsd": ArithmeticInstruction,
		"divss": ArithmeticInstruction, "incl": ArithmeticInstruction, "incq": ArithmeticInstruction,
		"paddd": ArithmeticInstruction, "palignr": ArithmeticInstruction, "shld": ArithmeticInstruction,
		"shrd": ArithmeticInstruction, "subq": ArithmeticInstruction, "subsd": ArithmeticInstruction,
		"subss": ArithmeticInstruction, "mulq": ArithmeticInstruction, "mulsd": ArithmeticInstruction,
		"mulx": ArithmeticInstruction, "psubd": ArithmeticInstruction, "ucomisd": ArithmeticInstruction,
		"ucomiss": ArithmeticInstruction, "vaddsd": ArithmeticInstruction, "vfmadd213sd": ArithmeticInstruction,
		"vfnmadd231sd": ArithmeticInstruction, "vpaddd": ArithmeticInstruction, "vpaddq": ArithmeticInstruction,
		"bsf": TransmissionInstruction, "vpcmpeqb": ArithmeticInstruction, "mulss": ArithmeticInstruction,
		"bsr": TransmissionInstruction, "cmova": TransmissionInstruction, "cmovae": TransmissionInstruction,
		"cmovb": TransmissionInstruction, "cmovbe": TransmissionInstruction, "cmove": TransmissionInstruction,
		"cmovg": TransmissionInstruction, "cmovge": TransmissionInstruction, "cmovl": TransmissionInstruction,
		"cmovle": TransmissionInstruction, "cmovne": TransmissionInstruction, "lfence": TransmissionInstruction,
		"movabs": TransmissionInstruction, "movapd": TransmissionInstruction, "movb": TransmissionInstruction,
		"movdqa": TransmissionInstruction, "movdqu": TransmissionInstruction, "movl": TransmissionInstruction,
		"movq": TransmissionInstruction, "movsbq": TransmissionInstruction, "movslq": TransmissionInstruction,
		"movss": TransmissionInstruction, "movswq": TransmissionInstruction, "movups": TransmissionInstruction,
		"movw": TransmissionInstruction, "movzbl": TransmissionInstruction, "movzwl": TransmissionInstruction,
		"pinsrd": TransmissionInstruction, "pinsrq": TransmissionInstruction, "pinsrw": TransmissionInstruction,
		"pmovmskb": TransmissionInstruction, "popfq": TransmissionInstruction, "prefetchnta": TransmissionInstruction,
		"pushfq": TransmissionInstruction, "sto": TransmissionInstruction, "mfence": TransmissionInstruction,
		"sfence": TransmissionInstruction, "vinserti128": TransmissionInstruction, "movd": TransmissionInstruction,
		"vmovdqa": TransmissionInstruction, "vmovdqu": TransmissionInstruction, "vmovntdq": TransmissionInstruction,
		"vpalignr": TransmissionInstruction, "vpbroadcastb": TransmissionInstruction, "vpshufd": LogicalInstruction,
		"vperm2f128": TransmissionInstruction, "vperm2i128": TransmissionInstruction, "vpmovmskb": TransmissionInstruction,
		"neg": LogicalInstruction, "orpd": LogicalInstruction, "orq": LogicalInstruction, "pand": LogicalInstruction,
		"pandn": LogicalInstruction, "pblendw": LogicalInstruction, "pcmpeqb": LogicalInstruction, "pcmpeqd": LogicalInstruction,
		"popcnt": LogicalInstruction, "punpcklbw": LogicalInstruction, "pxor": LogicalInstruction, "rex.W": LogicalInstruction,
		"h": LogicalInstruction, "rex.WXB": LogicalInstruction, "rorx": LogicalInstruction, "testb": LogicalInstruction,
		"vpand": LogicalInstruction, "vpblendd": LogicalInstruction, "vpor": LogicalInstruction, "vpshufb": LogicalInstruction,
		"vpslld": LogicalInstruction, "vpslldq": LogicalInstruction, "vpsllq": LogicalInstruction, "vpsrld": LogicalInstruction,
		"vpsrldq": LogicalInstruction, "vpsrlq": LogicalInstruction, "vptest": LogicalInstruction, "vpxor": LogicalInstruction,
		"xorps": LogicalInstruction, "callq": ProgramTransferInstruction, "ja": ProgramTransferInstruction,
		"jmpq": ProgramTransferInstruction, "lret": ProgramTransferInstruction, "retq": ProgramTransferInstruction,
		"syscall": ProgramTransferInstruction, "cpuid": ProcessorControlInstruction, "nopl": ProcessorControlInstruction,
		"nopw": ProcessorControlInstruction, "pause": ProcessorControlInstruction, "pshufb": ProcessorControlInstruction,
		"pshufd": ProcessorControlInstruction, "pshufhw": ProcessorControlInstruction, "rdtsc": ProcessorControlInstruction,
		"seta": ProcessorControlInstruction, "setae": ProcessorControlInstruction, "setb": ProcessorControlInstruction,
		"setbe": ProcessorControlInstruction, "sete": ProcessorControlInstruction, "setg": ProcessorControlInstruction,
		"setge": ProcessorControlInstruction, "setl": ProcessorControlInstruction, "setle": ProcessorControlInstruction,
		"setne": ProcessorControlInstruction, "setnp": ProcessorControlInstruction, "seto": ProcessorControlInstruction,
		"setp": ProcessorControlInstruction, "vzeroupper": ProcessorControlInstruction, "xgetbv": ProcessorControlInstruction,
		"pcmpestri": StringInstruction, "int3": InterruptInstruction, "ud2": InterruptInstruction, "cvtsd2ss": ArithmeticInstruction,
		"rdtscp": ProcessorControlInstruction, "prefetcht0": TransmissionInstruction, "sha256msg1": ArithmeticInstruction,
		"sha256msg2": ArithmeticInstruction, "sha256rnds2": ArithmeticInstruction, "aesenc": ArithmeticInstruction, "cvttsd2si": ArithmeticInstruction,
		"aesdec": ArithmeticInstruction, "aesdeclast": ArithmeticInstruction, "aesenclast": ArithmeticInstruction, "aesimc": ArithmeticInstruction,
		"aeskeygenassist": ArithmeticInstruction, "andl": LogicalInstruction, "andnpd": LogicalInstruction,
		"cmpltsd": ArithmeticInstruction, "cvttss2si": TransmissionInstruction, "movaps": TransmissionInstruction,
		"orl": LogicalInstruction, "pclmulhqhqdq": ArithmeticInstruction, "pclmulhqlqdq": ArithmeticInstruction,
		"pclmullqhqdq": ArithmeticInstruction, "pclmullqlqdq": ArithmeticInstruction, "pextrb": TransmissionInstruction,
		"pextrd": TransmissionInstruction, "pinsrb": TransmissionInstruction, "pslld": LogicalInstruction,
		"pslldq": LogicalInstruction, "psrad": LogicalInstruction, "psrld": LogicalInstruction, "psrldq": LogicalInstruction,
		"psrlq": LogicalInstruction, "roundsd": ArithmeticInstruction, "shufps": ProcessorControlInstruction,
		"sqrtsd": ArithmeticInstruction, "vbroadcasti128": TransmissionInstruction, "cqo": TransmissionInstruction,
		"leave": ProgramTransferInstruction,
	}
)

const (
	TransmissionInstruction Class = iota
	IOInstruction
	ArithmeticInstruction
	LogicalInstruction
	StringInstruction
	ProgramTransferInstruction
	InterruptInstruction
	PseudoInstruction
	ProcessorControlInstruction
	OtherInstruction
)

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
		log.Println(err)
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
	class := actionClassify(action)
	if ff, ok := funcfeatures.(*graph.Node); ok {
		ff.TotalInstruction++
		switch class {
		case TransmissionInstruction:
			ff.TransmissionInstruction++
		case IOInstruction:
			ff.IOInstruction++
		case ArithmeticInstruction:
			ff.ArithmeticInstruction++
		case LogicalInstruction:
			ff.LogicalInstruction++
		case StringInstruction:
			ff.StringInstruction++
		case ProgramTransferInstruction:
			ff.ProgramTransferInstruction++
			callee, err := ie.callInstArgs(inst)
			if err != nil {
				logger.Info((err.Error()))
				return
			}
			ff.AddCallee(callee)
		case InterruptInstruction:
			ff.InterruptInstruction++
		case PseudoInstruction:
			ff.PseudoInstruction++
		case ProcessorControlInstruction:
			ff.ProcessorControlInstruction++
		case OtherInstruction:
			ff.OtherInstruction++
		}
	}
}

func actionClassify(action string) Class {
	if v, ok := X86Instrcutions[action]; ok {
		UsedAction[action] = struct{}{}
		return v
	} else {
		MissedAction[action] = struct{}{}
		return OtherInstruction
	}
}