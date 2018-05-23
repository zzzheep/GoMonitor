package Model

import (
	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	MemoryPercent float32 `json:"memory_percent"`
	Name          string  `json:"name"`
	Id            int32   `json:"id"`
	CPUPercent    float64 `json:"cpu_percent"`
	Status        string  `json:"status"`
}

func GetProcessInfo() []ProcessInfo {
	p, _ := process.Processes()
	processList := make([]ProcessInfo, 0)
	for _, pchild := range p {
		if pchild.Pid != 0 {
			memoryPercent, _ := pchild.MemoryPercent()
			name, _ := pchild.Name()
			status, err := pchild.Status()
			if err != nil {
				status = "不支持"
			}
			cpuPercent, _ := pchild.CPUPercent()
			processList = append(processList, ProcessInfo{
				MemoryPercent: memoryPercent,
				Name:          name,
				Id:            pchild.Pid,
				Status:        status,
				CPUPercent:    cpuPercent,
			})
		}
	}
	return processList
}
