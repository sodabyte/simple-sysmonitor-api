package main

import (
    "simple-sysmonitor-api/handlers"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    r.Use(cors.Default())

    r.GET("/processes", handlers.GetProcessesHandler)
    r.GET("/disk-usage", handlers.GetDiskUsageHandler)
    r.GET("/disk-io", handlers.GetDiskIOHandler)
    r.GET("/partitions", handlers.GetPartitionsHandler)
    r.GET("/memory-usage", handlers.GetMemoryUsageHandler)
    r.GET("/cpu-usage", handlers.GetCPUUsageHandler)
    r.Run(":8080")
}
