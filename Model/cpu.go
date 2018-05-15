package Model

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CpuInfo struct {
	Name string  `json:"name,omitempty"`
	Used float64 `json:"used,omitempty"`
}

func GetCpuInfo() CpuInfo {
	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false)
	return CpuInfo{
		Name: c[0].ModelName,
		Used: cc[0],
	}
}
