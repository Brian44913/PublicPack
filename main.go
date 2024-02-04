package main
 
import (
	"fmt"
	"strings"
	"github.com/Brian44913/PublicPack/hardware"
	"github.com/Brian44913/PublicPack/code"
	"github.com/Brian44913/PublicPack/other"
)

func main() {
	// hardware
	OS,_ := hardware.GetOS()
	fmt.Println("OS:", OS)
	
	GPUName,_ := hardware.GetGPUName()
	fmt.Println("GPUName:", GPUName)
	
	CPUName := hardware.GetCPUName()
	fmt.Println("CPUName:", CPUName)
	
	BoardName,_ := hardware.GetMotherboardName()
	fmt.Println("BoardName:", BoardName)
	
	Speed := hardware.GetSpeed()
	fmt.Println("Speed:", Speed)
	
	totalUsedGB, _ := hardware.GetUsedMemory()
	fmt.Println("totalUsedGB:", totalUsedGB)
	
	Public_IP := hardware.GetLocalIP(`public`)
	fmt.Println("Public_IP:", Public_IP)
	Intranet_IP := hardware.GetLocalIP("intranet")
	fmt.Println("Intranet_IP:", Intranet_IP)
	Gateway,_ := hardware.GetDefaultGateway()
	fmt.Println("Gateway:", Gateway)
	
	// new 
	BoardInfo, err := hardware.GetMotherboardInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Motherboard Manufacturer: %s\n", BoardInfo.Manufacturer)
	fmt.Printf("Motherboard Model: %s\n", BoardInfo.Model)
	fmt.Printf("Motherboard Serial Number: %s\n", BoardInfo.SerialNumber)
	
	CPUInfo, err := hardware.GetCPUInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("CPU Model: %s\n", CPUInfo.Model)
	
	GPUInfo, err := hardware.GetGPUInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("GPU Models: %s\n", strings.Join(GPUInfo.Models, ", "))
	for i, model := range GPUInfo.Models {
		fmt.Printf("GPU #%d Model: %s\n", i+1, model)
		fmt.Printf("GPU #%d UUID: %s\n", i+1, GPUInfo.UUIDs[i])
		fmt.Printf("GPU #%d Memory: %s\n", i+1, GPUInfo.Memory[i])
	}
	
	DiskInfo, err := hardware.GetDiskInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, disk := range DiskInfo {
		fmt.Printf("Disk #%d Model: %s\n", i+1, disk.Model)
		fmt.Printf("Disk #%d SN: %s\n", i+1, disk.SN)
		fmt.Printf("Disk #%d Size: %s\n", i+1, disk.Size)
	}
	
	MemoryInfo, err := hardware.GetMemoryInfo()
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
	
	PowerSupplyInfo, err := hardware.GetPowerSupplyInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, powerSupply := range PowerSupplyInfo {
		fmt.Printf("Power Supply #%d Manufacturer: %s\n", i+1, powerSupply.Manufacturer)
		fmt.Printf("Power Supply #%d Model: %s\n", i+1, powerSupply.Model)
		fmt.Printf("Power Supply #%d SN: %s\n", i+1, powerSupply.SN)
	}
	
	// code
	Base64UrlEncode := code.Base64UrlEncode("https://github.com/Brian44913/PublicPack")
	fmt.Println("Base64UrlEncode:", Base64UrlEncode)
	Base64UrlDecode,_ := code.Base64UrlDecode(Base64UrlEncode)
	fmt.Println("Base64UrlDecode:", string(Base64UrlDecode))
	
	// other
	hostname, _ := other.ReadAll("/etc/hostname")
	fmt.Println("hostname:", string(hostname))
}