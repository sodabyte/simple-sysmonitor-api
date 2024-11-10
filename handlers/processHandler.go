package handlers

import (
    "net/http"
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
