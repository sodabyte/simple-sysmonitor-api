package handlers

import (
    "net/http"
    "simple-sysmonitor-api/services"
    "github.com/gin-gonic/gin"
)

func GetDiskUsageHandler(c *gin.Context) {
    diskUsage, err := services.GetDiskUsage()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch disk usage"})
        return
    }
    c.JSON(http.StatusOK, diskUsage)
}

func GetDiskIOHandler(c *gin.Context) {
    diskUsage, err := services.GetDiskIO()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch disk usage"})
        return
    }
    c.JSON(http.StatusOK, diskUsage)
}

func GetPartitionsHandler(c *gin.Context) {
    partitions, err := services.GetDiskPartitions()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch partitions"})
        return
    }
    c.JSON(http.StatusOK, partitions)
}

func GetMemoryUsageHandler(c *gin.Context) {
    memoryUsage, err := services.GetMemoryUsage()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch memory usage"})
        return
    }
    c.JSON(http.StatusOK, memoryUsage)
}

func GetCPUUsageHandler(c *gin.Context) {
    cpuUsage, err := services.GetCPUUsage()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch CPU usage"})
        return
    }
    c.JSON(http.StatusOK, cpuUsage)
}
