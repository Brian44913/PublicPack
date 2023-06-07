package PublicPackageHardware

import (
    "fmt"
    "net"
	"os/exec"
	"strings"
	"github.com/safchain/ethtool"
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