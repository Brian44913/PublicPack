package main
 
import (
	"fmt"
	"strings"
	"github.com/Brian44913/PublicPack/hardware"
	"github.com/Brian44913/PublicPack/code"
	"github.com/Brian44913/PublicPack/other"
)

func main() {
	// PublicPackHardware
	OS,_ := PublicPackHardware.GetOS()
	fmt.Println("OS:", OS)
	
	GPUName,_ := PublicPackHardware.GetGPUName()
	fmt.Println("GPUName:", GPUName)
	
	CPUName := PublicPackHardware.GetCPUName()
	fmt.Println("CPUName:", CPUName)
	
	BoardName,_ := PublicPackHardware.GetMotherboardName()
	fmt.Println("BoardName:", BoardName)
	
	Speed := PublicPackHardware.GetSpeed()
	fmt.Println("Speed:", Speed)
	
	totalUsedGB, _ := PublicPackHardware.GetUsedMemory()
	fmt.Println("totalUsedGB:", totalUsedGB)
	
	Public_IP := PublicPackHardware.GetLocalIP(`public`)
	fmt.Println("Public_IP:", Public_IP)
	Intranet_IP := PublicPackHardware.GetLocalIP("intranet")
	fmt.Println("Intranet_IP:", Intranet_IP)
	Gateway,_ := PublicPackHardware.GetDefaultGateway()
	fmt.Println("Gateway:", Gateway)
	
	// new 
	BoardInfo, err := PublicPackHardware.GetMotherboardInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Motherboard Manufacturer: %s\n", BoardInfo.Manufacturer)
	fmt.Printf("Motherboard Model: %s\n", BoardInfo.Model)
	fmt.Printf("Motherboard Serial Number: %s\n", BoardInfo.SerialNumber)
	
	CPUInfo, err := PublicPackHardware.GetCPUInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("CPU Model: %s\n", CPUInfo.Model)
	
	GPUInfo, err := PublicPackHardware.GetGPUInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("GPU Models: %s\n", strings.Join(GPUInfo.Models, ", "))
	
	DiskInfo, err := PublicPackHardware.GetDiskInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, disk := range DiskInfo {
		fmt.Printf("Disk #%d Model: %s\n", i+1, disk.Model)
		fmt.Printf("Disk #%d SN: %s\n", i+1, disk.SN)
		fmt.Printf("Disk #%d Size: %s\n", i+1, disk.Size)
	}
	
	MemoryInfo, err := PublicPackHardware.GetMemoryInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, memory := range MemoryInfo {
		fmt.Printf("Memory #%d Model: %s\n", i+1, memory.Model)
		fmt.Printf("Memory #%d Part Number: %s\n", i+1, memory.PartNumber)
		fmt.Printf("Memory #%d Speed: %s\n", i+1, memory.Speed)
		fmt.Printf("Memory #%d SN: %s\n", i+1, memory.SN)
		fmt.Printf("Memory #%d Size: %s\n", i+1, memory.Size)
	}
	
	PowerSupplyInfo, err := PublicPackHardware.GetPowerSupplyInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, powerSupply := range PowerSupplyInfo {
		fmt.Printf("Power Supply #%d Manufacturer: %s\n", i+1, powerSupply.Manufacturer)
		fmt.Printf("Power Supply #%d Model: %s\n", i+1, powerSupply.Model)
		fmt.Printf("Power Supply #%d SN: %s\n", i+1, powerSupply.SN)
	}
	
	// PublicPackCode
	Base64UrlEncode := PublicPackCode.Base64UrlEncode("https://github.com/Brian44913/PublicPack")
	fmt.Println("Base64UrlEncode:", Base64UrlEncode)
	Base64UrlDecode,_ := PublicPackCode.Base64UrlDecode(Base64UrlEncode)
	fmt.Println("Base64UrlDecode:", string(Base64UrlDecode))
	
	// PublicPackOther
	hostname, _ := PublicPackOther.ReadAll("/etc/hostname")
	fmt.Println("hostname:", string(hostname))
	
	lotus_v, _ := PublicPackOther.GetBinV("/root/sh/.bash/lotus","-v")
	fmt.Println("lotus_v:", lotus_v)
}