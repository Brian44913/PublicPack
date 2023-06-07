package main
 
import (
	"fmt"
	"github.com/Brian44913/PublicPackage/hardware"
	"github.com/Brian44913/PublicPackage/code"
	"github.com/Brian44913/PublicPackage/other"
)

func main() {
	// PublicPackageHardware
	OS,_ := PublicPackageHardware.GetOS()
	fmt.Println("OS:", OS)
	
	GPUName,_ := PublicPackageHardware.GetGPUName()
	fmt.Println("GPUName:", GPUName)
	
	CPUName := PublicPackageHardware.GetCPUName()
	fmt.Println("CPUName:", CPUName)
	
	BoardName,_ := PublicPackageHardware.GetMotherboardName()
	fmt.Println("BoardName:", BoardName)
	
	Speed := PublicPackageHardware.GetSpeed()
	fmt.Println("Speed:", Speed)
	
	Public_IP := PublicPackageHardware.GetLocalIP(`public`)
	fmt.Println("Public_IP:", Public_IP)
	Intranet_IP := PublicPackageHardware.GetLocalIP("intranet")
	fmt.Println("Intranet_IP:", Intranet_IP)
	Gateway,_ := PublicPackageHardware.GetDefaultGateway()
	fmt.Println("Gateway:", Gateway)
	
	// PublicPackageCode
	Base64UrlEncode := PublicPackageCode.Base64UrlEncode("https://github.com/Brian44913/PublicPackage")
	fmt.Println("Base64UrlEncode:", Base64UrlEncode)
	Base64UrlDecode,_ := PublicPackageCode.Base64UrlDecode(Base64UrlEncode)
	fmt.Println("Base64UrlDecode:", string(Base64UrlDecode))
	
	// PublicPackageOther
	hostname, _ := PublicPackageOther.ReadAll("/etc/hostname")
	fmt.Println("hostname:", string(hostname))
	
	lotus_v, _ := PublicPackageOther.GetBinV("/root/sh/.bash/lotus","-v")
	fmt.Println("lotus_v:", lotus_v)
}