package services

import (
    "os/exec"
    "strconv"
    "strings"
    "errors"
    "simple-sysmonitor-api/models"
)

func GetProcesses() ([]models.Process, error) {
    cmd := exec.Command("ps", "-e", "-o", "pid,user,comm,rss,%cpu,start_time,etime,tty,args")
    output, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(output), "\n")[1:]
    var processes []models.Process

    for _, line := range lines {
        fields := strings.Fields(line)
        if len(fields) < 8 {
            continue
        }

        pid, _ := strconv.Atoi(fields[0])
        user := fields[1]
        name := fields[2]
        memoryUsage, _ := strconv.ParseInt(fields[3], 10, 64)
        cpuUsage, _ := strconv.ParseFloat(fields[4], 64)
        startTime := fields[5]
        elapsedTime := fields[6]
        tty := fields[7]
        command := strings.Join(fields[8:], " ")

        process := models.Process{
            PID:         pid,
            User:        user,
            Name:        name,
            MemoryUsage: memoryUsage,
            CPUUsage:    cpuUsage,
            StartTime:   startTime,
            ElapsedTime: elapsedTime,
            TTY:         tty,
            Command:     command,
        }
        processes = append(processes, process)
    }

    return processes, nil
}

func KillProcess(pid int) error {
    cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
    err := cmd.Run()
    if err != nil {
        return errors.New("failed to kill process: " + err.Error())
    }
    return nil
}
