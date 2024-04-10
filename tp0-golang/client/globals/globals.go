package globals

type Config struct {
	Ip      string `json:"ip"`
	Puerto  int    `json:"puerto"`
	Mensaje string `json:"mensaje"`
}

var ClientConfig *Config

// ------------------------------------------------------------------------------
type ConfigKernel struct {
	Port                int      `json:"port"`
	Ip_memory           string   `json:"ip_memory"`
	Port_memory         int      `json:"port_memory"`
	Ip_cpu              string   `json:"ip_cpu"`
	Port_cpu            int      `json:"port_cpu"`
	Planning_algorithm  string   `json:"planning_algorithm"`
	Quantum             int      `json:"quantum"`
	Resources           []string `json:"resources"`
	Resources_instances []int    `json:"resources_instances"`
	Multiprogramming    int      `json:"multiprogramming"`
}

var KernelConfig *ConfigKernel

type ConfigCpu struct {
	Port               int    `json:"port"`
	Ip_memory          string `json:"ip_memory"`
	Port_memory        int    `json:"port_memory"`
	Number_felling_tlb int    `json:"number_felling_tlb"`
	Algorithm_tlb      int    `json:"algorithm_tlb"`
}

var CpuConfig *ConfigCpu

type ConfigMemory struct {
	Port             int    `json:"port"`
	Memory_size      int    `json:"memory_size"`
	Page_size        int    `json:"page_size"`
	Instruction_path string `json:"instruction_path"`
	Delay_response   int    `json:"delay_response"`
}

var MemoryConfig *ConfigMemory

type ConfigInterface struct {
	Port               int    `json:"port"`
	Type               string `json:"type"`
	Unit_work_time     int    `json:"unit_work_time"`
	Ip_kernel          string `json:"ip_kernel"`
	Port_kernel        int    `json:"port_kernel"`
	Ip_memory          string `json:"ip_memory"`
	Port_memory        int    `json:"port_memory"`
	Dialfs_path        string `json:"dialfs_path"`
	Dialfs_block_size  int    `json:"dialfs_block_size"`
	Dialfs_block_count int    `json:"dialfs_block_count"`
}

var InterfaceConfig *ConfigInterface
