package Model

import (
	"github.com/shirou/gopsutil/net"
)

func GetNetInfo() []net.IOCountersStat {
	nv, _ := net.IOCounters(true)
	return nv
}
