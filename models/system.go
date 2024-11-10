package models

type DiskUsage struct {
    Total       uint64  `json:"total"`
    Used        uint64  `json:"used"`
    Free        uint64  `json:"free"`
    UsedPercent float64 `json:"used_percent"`
}

type DiskIO struct {
    ReadCount  uint64 `json:"read_count"`
    WriteCount uint64 `json:"write_count"`
    ReadBytes  uint64 `json:"read_bytes"`
    WriteBytes uint64 `json:"write_bytes"`
    ReadTime   uint64 `json:"read_time"`
    WriteTime  uint64 `json:"write_time"`
}

type DiskPartition struct {
    Device     string `json:"device"`
    Mountpoint string `json:"mount_point"`
    Fstype     string `json:"fstype"`
}

type MemoryUsage struct {
    Total       uint64  `json:"total"`
    Used        uint64  `json:"used"`
    Free        uint64  `json:"free"`
    UsedPercent float64 `json:"used_percent"`
}

type CPUUsage struct {
    UsedPercent    float64 `json:"used_percent"`
    CPUInfo        string  `json:"cpu_info"`
    CPUNumCores    int     `json:"cpu_num_cores"`
    CPUUserTime    float64 `json:"cpu_user_time"`
    CPUSystemTime  float64 `json:"cpu_system_time"`
    CPUIDLETime    float64 `json:"cpu_idle_time"`
}
