package main
 
import (
	"fmt"
	"PublicPackageHardware"
	"PublicPackageCode"
	"PublicPackageOther"
)

func main() {
	// PublicPackageHardware
	OS,_ := PublicPackageHardware.GetOS()
	fmt.Println("OS:", OS)
	
	Speed := PublicPackageHardware.GetSpeed()
	fmt.Println("Speed:", Speed)
	
	Public_IP, err := PublicPackageHardware.GetLocalIP(`public`)
	if err != nil {
        fmt.Println(err)
    }
	fmt.Println("Public_IP:", Public_IP)
	Intranet_IP, err := PublicPackageHardware.GetLocalIP("intranet")
	if err != nil {
        fmt.Println(err)
    }
	fmt.Println("Intranet_IP:", Intranet_IP)
	
	// PublicPackageCode
	Base64UrlEncode := PublicPackageCode.Base64UrlEncode("https://filecoin.plus/")
	fmt.Println("Base64UrlEncode:", Base64UrlEncode)
	Base64UrlDecode,_ := PublicPackageCode.Base64UrlDecode(Base64UrlEncode)
	fmt.Println("Base64UrlDecode:", string(Base64UrlDecode))
	
	// PublicPackageOther
	hostname, _ := PublicPackageOther.ReadAll("/etc/hostname")
	fmt.Println("hostname:", string(hostname))
	
}