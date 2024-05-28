package def

const (
	CommendOfSharedLibs             = "ldd "
	CommendOfDisassembly            = "objdump -M intel -d "
	BinaryFilePathPrefix            = "/mnt/data/nfs/myx/tmp/datasets/"
	BasePath                        = "/mnt/data/nfs/myx/tmp/"
	CudaFlags                       = "cudaFree|cudaMemcpy|cudaMalloc"
	GraphKernalDefaultIteratorTimes = 2
	// MetricsNumber                       = 18
	FeatureDividBase                    = 3
	FileWeight, CFWWeight               = 0.7, 0.3
	JsonDatabase                        = "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86"
	PercentageDecline                   = 0.8
	Delta                               = 1.0
	Alpha                               = 1.0
	MaxTotalInstDiff                    = 100000000.0
	StatisticalVectorIntersectionWeight = 0.5
	StatisticalVectorDisjointWeight     = 0.5
	TotalInstWeight                     = 0.5
	ProgramFeatureWeight                = 0.5
	MaxGoroutines                       = 1
	Debug                               = false
)

var (
	// Hierarchies = []int{12, 6, 2, 6, 6, 6, 8, 4, 2, 6, 4, 6}
	Weight = []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}
)

// func Buckets() int {
// 	length := 1
// 	for _, v := range Hierarchies {
// 		length *= (v + 1)
// 	}
// 	return length
// }
