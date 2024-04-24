package hardware

import (
    "fmt"
    "net"
	"os/exec"
	"strings"
	"github.com/safchain/ethtool"
	"math"
	"time"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetOS() (string, error){
	cmd := exec.Command("lsb_release", "-ds")
	output, err := cmd.Output()
	if err != nil {
		return "",fmt.Errorf("getOS error", err)
	}

	version := strings.TrimSpace(string(output))
	return version, nil
}
func GetSpeed() int {
	e, err := ethtool.NewEthtool()
	if err != nil {
		panic(err.Error())
	}
	defer e.Close()
	cmdGet, err := e.CmdGetMapped("bond0")
	if err != nil {
		ifconfig, _ := exec.Command("/bin/bash", "-c", "/sbin/ifconfig | grep RUNNING | grep -vE 'lo|bond0' | awk '{print $1}' | sed 's/://'").Output()
		for _, device := range strings.Split(string(ifconfig), "\n") {
			if len(device) > 1 {
				stats, err := e.CmdGetMapped(device)
				if err != nil {
					panic(err.Error())
				}
				if stats["speed"]>0{
					return int(stats["speed"])
				}
				//fmt.Printf("LinkName: %s LinkState: %d\n", device, stats["speed"])
			}

		}
		return 0
	}
	return int(cmdGet["speed"])
}
func GetLocalIP(network string) (ip string) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, addr := range addrs {
        ipAddr, ok := addr.(*net.IPNet)
		//fmt.Println(ipAddr)
        if !ok {
            continue
        }
        if ipAddr.IP.IsLoopback() {
            continue
        }
        if !ipAddr.IP.IsGlobalUnicast() {
            continue
        }
		if(network == "public" && strings.Index(ipAddr.IP.String(), "192.168")== -1){
			return ipAddr.IP.String()
		}
		if(network == "intranet" && strings.Index(ipAddr.IP.String(), "192.168")!= -1){
			return ipAddr.IP.String()
		}
    }
    return ""
}
func getIpFromAddr(addr net.Addr) net.IP {
    var ip net.IP
    switch v := addr.(type) {
    case *net.IPNet:
        ip = v.IP
    case *net.IPAddr:
        ip = v.IP
    }
    if ip == nil || ip.IsLoopback() {
        return nil
    }
    ip = ip.To4()
    if ip == nil {
        return nil // not an ipv4 address
    }

    return ip
}
func BootTime() string{
	boottime, _ := host.BootTime()
    return time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")
}
func GetDefaultGateway() (string, error) {
	cmd := exec.Command("bash", "-c", "ip route | grep default | awk '{print $3}' | head -n 1")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// 删除尾部的换行符
	gateway := strings.TrimSpace(string(output))
	return gateway, nil
}
func CheckRoute() (int, error) {
	cmd := exec.Command("ip", "route")
	output, err := cmd.Output()

	if err != nil {
		return 0, err
	}

	if strings.Contains(string(output), "8.0.0.0") {
		return 1, nil
	}

	return 0, nil
}
func CheckUFW(ip string) (bool, error) {
    cmd := exec.Command("/usr/sbin/ufw", "status")
    output, err := cmd.Output()

    if err != nil {
        return false, err
    }

    outputStr := string(output)

    if strings.Contains(outputStr, "Status: active") && strings.Contains(outputStr, ip) {
        return true, nil
    }

    return false, nil
}

func RateDisk(dir string) int {
	info, _ := disk.Usage(dir)
	return int(math.Floor(float64(info.UsedPercent) + 0.5))
}
func GetUsedMemory() (int, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	totalUsedGB := int(vmStat.Used) / 1024 / 1024 / 1024
	return totalUsedGB, nil
}