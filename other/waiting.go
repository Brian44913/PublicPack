package other

import (
	"fmt"
	"net"
	"time"
)

func ipToLong(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 + uint32(ip[1])<<16 + uint32(ip[2])<<8 + uint32(ip[3])
}

func Waiting(host string,maxTime uint32) {
	ip := net.ParseIP(host)
	if ip == nil {
		fmt.Println("无效的IP地址")
		return
	}
	ipLong := ipToLong(ip)

	// 计算除以maxTime的余数
	waitSeconds := ipLong % maxTime
	// fmt.Printf("等待时间: %d 秒\n", waitSeconds)
	time.Sleep(time.Duration(waitSeconds) * time.Second)
}