package Model

import (
	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	MemoryPercent float32 `json:"memory_percent,omitempty"`
	Name          string  `json:"name,omitempty"`
	Id            int32   `json:"id,omitempty"`
	CPUPercent    float64 `json:"cpu_percent,omitempty"`
	Status        string  `json:"status,omitempty"`
}

func GetProcessInfo() []ProcessInfo {
	p, _ := process.Processes()
	processList := make([]ProcessInfo, len(p))
	for _, pchild := range p {
		memoryPercent, _ := pchild.MemoryPercent()
		name, _ := pchild.Name()
		status, _ := pchild.Status()
		cpuPercent, _ := pchild.CPUPercent()
		processList = append(processList, ProcessInfo{
			MemoryPercent: memoryPercent,
			Name:          name,
			Id:            pchild.Pid,
			Status:        status,
			CPUPercent:    cpuPercent,
		})
	}
	return processList
}
