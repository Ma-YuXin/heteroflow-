package def

const (
	CommendOfSharedLibs             = "ldd "
	CommendOfDisassembly            = "objdump -M intel -d "
	BinaryFilePathPrefix            = "/mnt/data/myx/tmp/datasets/"
	BasePath                        = "/mnt/data/myx/tmp/"
	CudaFlags                       = "cudaFree|cudaMemcpy|cudaMalloc"
	GraphKernalDefaultIteratorTimes = 2
	// MetricsNumber                       = 18
	FeatureDividBase = 3
	// FileWeight, CFWWeight               = 0.7, 0.3
	JsonDatabase                        = "/mnt/data/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X64"
	PercentageDecline                   = 0.8
	Delta                               = 1.0
	Alpha                               = 1.0
	MaxTotalInstDiff                    = 100000000.0
	StatisticalVectorIntersectionWeight = 0.5
	StatisticalVectorDisjointWeight     = 0.5
	TotalInstWeight                     = 0.5
	ProgramFeatureWeight                = 0.5
	MaxGoroutines                       = 70
	Debug                               = false
	LargestElementsToFind               = 10
)

var (
	BehaviorInstCounterHierarchies = []int{
		4, // Total:
		4, // Transmission:
		2, // IO:
		4, // Arithmetic:
		4, // Logical:
		4, // String:
		4, // ProgramTransfer:
		2, // Interrupt:
		2, // Pseudo:
		4, // ProcessorControl:
	}
	TechInstCounterHierarchies = []int{
		5, // Total:
		5, // VIRTUALIZATION:
		5, // GP:
		5, // GP_EXT:
		5, // GP_IN_OUT:
		5, // FPU:
		5, // MMX:
		5, // STATE:
		5, // SIMD:
		5, // SSE:
		5, // SCALAR:
		5, // CRYPTO_HASH:
		5, // AVX:
		5, // AVX512:
		5, // MASK:
		5, // AMX:
	}
	//  = []int{12, 6, 2, 6, 6, 6, 8, 4, 2, 6, 4, 6}
	Weight = []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}
)

// func Buckets() int {
// 	length := 1
// 	for _, v := range Hierarchies {
// 		length *= (v + 1)
// 	}
// 	return length
// }
