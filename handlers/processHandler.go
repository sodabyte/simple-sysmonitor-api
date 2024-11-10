package handlers

import (
    "net/http"
    "strconv"
    "simple-sysmonitor-api/services"
    "github.com/gin-gonic/gin"
)

func GetProcessesHandler(c *gin.Context) {
    processes, err := services.GetProcesses()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch processes"})
        return
    }

    c.JSON(http.StatusOK, processes)
}

func KillProcessHandler(c *gin.Context) {
    pidStr := c.Param("pid")
    pid, err := strconv.Atoi(pidStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PID"})
        return
    }

    err = services.KillProcess(pid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to kill process", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Process killed successfully"})
}
