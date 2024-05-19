package definition

const (
	CommendOfSharedLibs             = "ldd "
	CommendOfDisassembly            = "objdump -M intel -d "
	BinaryFilePathPrefix            = "/mnt/data/nfs/myx/tmp/datasets/"
	BasePath                        = "/mnt/data/nfs/myx/tmp/"
	CudaFlags                       = "cudaFree|cudaMemcpy|cudaMalloc"
	GraphKernalDefaultIteratorTimes = 2
	MetricsNumber                   = 12
	FeatureDividBase                = 3
	FileWeight, CFWWeight           = 0.7, 0.3
	JsonDatabase                    = "/mnt/data/nfs/myx/tmp/json/Asteria-Pro/buildroot-elf-5arch/X86"
	PercentageDecline               = 0.8
	Delta                           = 1.0
	Alpha                           = 1.0
)
