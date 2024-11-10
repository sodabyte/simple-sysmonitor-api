package services

import (
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
    "simple-sysmonitor-api/models"
    "time"
    "os"
)

func GetDiskUsage() (models.DiskUsage, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return models.DiskUsage{}, err
    }

    usage, err := disk.Usage(homeDir)
    if err != nil {
        return models.DiskUsage{}, err
    }
    return models.DiskUsage{
        Total:       usage.Total,
        Used:        usage.Used,
        Free:        usage.Free,
        UsedPercent: usage.UsedPercent,
    }, nil
}

func GetDiskIO() (models.DiskIO, error) {
    ioStats, err := disk.IOCounters("")
    if err != nil {
        return models.DiskIO{}, err
    }

    disk := ioStats["0"]

    return models.DiskIO{
        ReadCount:  disk.ReadCount,
        WriteCount: disk.WriteCount,
        ReadBytes:  disk.ReadBytes,
        WriteBytes: disk.WriteBytes,
        ReadTime:   disk.ReadTime,
        WriteTime:  disk.WriteTime,
    }, nil
}

func GetDiskPartitions() ([]models.DiskPartition, error) {
    partitions, err := disk.Partitions(true)
    if err != nil {
        return nil, err
    }

    var partitionModels []models.DiskPartition
    for _, part := range partitions {
        partitionModels = append(partitionModels, models.DiskPartition{
            Device:     part.Device,
            Mountpoint: part.Mountpoint,
            Fstype:     part.Fstype,
        })
    }

    return partitionModels, nil
}

func GetMemoryUsage() (models.MemoryUsage, error) {
    vmStat, err := mem.VirtualMemory()
    if err != nil {
        return models.MemoryUsage{}, err
    }
    return models.MemoryUsage{
        Total:       vmStat.Total,
        Used:        vmStat.Used,
        Free:        vmStat.Free,
        UsedPercent: vmStat.UsedPercent,
    }, nil
}

func GetCPUUsage() (models.CPUUsage, error) {
    percent, err := cpu.Percent(time.Second, false)
    if err != nil {
        return models.CPUUsage{}, err
    }

    times, err := cpu.Times(false)
    if err != nil {
        return models.CPUUsage{}, err
    }

    info, err := cpu.Info()
    if err != nil {
        return models.CPUUsage{}, err
    }

    numCores := int(info[0].Cores)

    return models.CPUUsage{
        UsedPercent:  percent[0],
        CPUInfo:      info[0].ModelName,
        CPUNumCores:  numCores,
        CPUUserTime:  times[0].User,
        CPUSystemTime: times[0].System,
        CPUIDLETime:  times[0].Idle,
    }, nil
}
