# PublicPack

```go
package main
 
import (
	"fmt"
	"github.com/Brian44913/PublicPack/hardware"
	"github.com/Brian44913/PublicPack/code"
	"github.com/Brian44913/PublicPack/other"
)

func main() {
	// PublicPackHardware
	OS,_ := PublicPackHardware.GetOS()
	fmt.Println("OS:", OS)
	
	Speed := PublicPackHardware.GetSpeed()
	fmt.Println("Speed:", Speed)
	
	Public_IP := PublicPackHardware.GetLocalIP(`public`)
	fmt.Println("Public_IP:", Public_IP)
	Intranet_IP := PublicPackHardware.GetLocalIP("intranet")
	fmt.Println("Intranet_IP:", Intranet_IP)
	
	// PublicPackCode
	Base64UrlEncode := PublicPackCode.Base64UrlEncode("https://github.com/Brian44913/PublicPack")
	fmt.Println("Base64UrlEncode:", Base64UrlEncode)
	Base64UrlDecode,_ := PublicPackCode.Base64UrlDecode(Base64UrlEncode)
	fmt.Println("Base64UrlDecode:", string(Base64UrlDecode))
	
	// PublicPackOther
	hostname, _ := PublicPackOther.ReadAll("/etc/hostname")
	fmt.Println("hostname:", string(hostname))
}
```
