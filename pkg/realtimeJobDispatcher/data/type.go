package data

// Result represents a single benchmark result entry
type BenchmarkResult struct {
	Benchmark             string `json:"Benchmark"`
	HardwareVendor        string `json:"Hardware Vendor\t"`
	System                string `json:"System"`
	Cores                 string `json:"# Cores"`
	Chips                 string `json:"# Chips"`
	EnabledThreadsPerCore string `json:"# Enabled Threads Per Core"`
	Processor             string `json:"Processor"`
	ProcessorMHz          string `json:"Processor MHz"`
	CPUsOrderable         string `json:"CPU(s) Orderable"`
	Parallel              string `json:"Parallel \t"`
	BasePointerSize       string `json:"Base Pointer Size"`
	PeakPointerSize       string `json:"Peak Pointer Size "`
	FirstLevelCache       string `json:"1st Level Cache"`
	SecondLevelCache      string `json:"2nd Level Cache"`
	ThirdLevelCache       string `json:"3rd Level Cache"`
	OtherCache            string `json:"Other Cache"`
	Memory                string `json:"Memory"`
	Storage               string `json:"Storage\t"`
	OperatingSystem       string `json:"Operating System"`
	FileSystem            string `json:"File System"`
	Compiler              string `json:"Compiler"`
	HWAvail               string `json:"HW Avail"`
	SWAvail               string `json:"SW Avail"`
	Result                string `json:"Result"`
	Baseline              string `json:"Baseline"`
	EnergyPeakResult      string `json:"Energy Peak Result"`
	EnergyBaseResult      string `json:"Energy Base Result\t"`
	Peak603               string `json:"603 Peak"`
	Base603               string `json:"603 Base"`
	Peak607               string `json:"607 Peak"`
	Base607               string `json:"607 Base"`
	Peak619               string `json:"619 Peak"`
	Base619               string `json:"619 Base"`
	Peak621               string `json:"621 Peak"`
	Base621               string `json:"621 Base"`
	Peak627               string `json:"627 Peak"`
	Base627               string `json:"627 Base"`
	Peak628               string `json:"628 Peak"`
	Base628               string `json:"628 Base"`
	Peak638               string `json:"638 Peak"`
	Base638               string `json:"638 Base"`
	Peak644               string `json:"644 Peak"`
	Base644               string `json:"644 Base"`
	Peak649               string `json:"649 Peak"`
	Base649               string `json:"649 Base"`
	Peak654               string `json:"654 Peak"`
	Base654               string `json:"654 Base"`
	License               string `json:"License"`
	TestedBy              string `json:"Tested By"`
	TestSponsor           string `json:"Test Sponsor\t"`
	TestDate              string `json:"Test Date"`
	Published             string `json:"Published"`
	Updated               string `json:"Updated "`
	Disclosure            string `json:"Disclosure"`
	Disclosures           string `json:"Disclosures"`
}
