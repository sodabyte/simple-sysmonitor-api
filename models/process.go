package models

type Process struct {
    Name          string `json:"name"`
    PID           int    `json:"pid"`
    MemoryUsage   int64  `json:"memory_usage,omitempty"`
    User          string `json:"user,omitempty"`
    CPUUsage      float64 `json:"cpu_usage,omitempty"`
    StartTime     string `json:"start_time,omitempty"`
    ElapsedTime   string `json:"elapsed_time,omitempty"`
    TTY           string `json:"tty,omitempty"`
    Command       string `json:"command,omitempty"`
}
