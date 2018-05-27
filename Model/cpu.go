package Model

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CpuInfo struct {
	Name string  `json:"name,omitempty"`
	Used float64 `json:"used,omitempty"`
}

func GetCpuInfo() []CpuInfo {
	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false)
	cpuList := make([]CpuInfo, 0, len(c))
	for i, _ := range c {
		cpuList = append(cpuList, CpuInfo{
			Name: c[i].ModelName,
			Used: cc[i],
		})
	}
	return cpuList
}
