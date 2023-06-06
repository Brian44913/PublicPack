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
func GetLocalIP(network string) (net.IP, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    for _, iface := range ifaces {
        if iface.Flags&net.FlagUp == 0 {
            continue // interface down
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue // loopback interface
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return nil, err
        }
        for _, addr := range addrs {
            ip := getIpFromAddr(addr)
            if ip == nil {
                continue
            }
			if(network == "public" && strings.Index(ip.String(), "192.168")== -1){
				return ip, nil
			}
			if(network == "intranet" && strings.Index(ip.String(), "192.168")!= -1){
				return ip, nil
			}
            //return ip, nil
        }
    }
    return nil, fmt.Errorf("connected to the network?")
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