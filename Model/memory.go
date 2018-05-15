package Model

import (
	"github.com/shirou/gopsutil/mem"
)

type MemoryInfo struct {
	Total       uint64  `json:"total,omitempty"`
	Available   uint64  `json:"available,omitempty"`
	Used        uint64  `json:"used,omitempty"`
	UsedPercent float64 `json:"used_percent,omitempty"`
}

func GetMemoryInfo() MemoryInfo {
	v, _ := mem.VirtualMemory()
	return MemoryInfo{
		Total:       v.Total / 1024 / 1024,
		Available:   v.Available / 1024 / 1024,
		Used:        v.Used / 1024 / 1024,
		UsedPercent: v.UsedPercent,
	}
}
