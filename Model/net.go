package Model

import (
	"github.com/shirou/gopsutil/net"
)

type NetInfo struct {
	BytesRecv uint64 `json:"bytes_recv,omitempty"`
	BytesSent uint64 `json:"bytes_sent,omitempty"`
}

func GetNetInfo() NetInfo {
	nv, _ := net.IOCounters(true)
	return NetInfo{
		BytesRecv: nv[0].BytesRecv,
		BytesSent: nv[0].BytesSent,
	}
	
}
