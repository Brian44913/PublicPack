package hardware

import (
	"fmt"
	"os/exec"
	"strings"
)

// MotherboardInfo 包含了主板的信息
type MotherboardInfo struct {
	Manufacturer string
	Model        string
	SerialNumber string
}

// GetMotherboardInfo 获取主板的信息
func GetMotherboardInfo() (MotherboardInfo, error) {
	cmd := exec.Command("/usr/sbin/dmidecode", "-t", "baseboard")
	out, err := cmd.Output()
	if err != nil {
		return MotherboardInfo{}, err
	}

	outStr := string(out)
	lines := strings.Split(outStr, "\n")
	
	info := MotherboardInfo{}

	for _, line := range lines {
		if strings.HasPrefix(line, "\tManufacturer:") {
			info.Manufacturer = strings.TrimPrefix(line, "\tManufacturer:")
			info.Manufacturer = strings.TrimSpace(info.Manufacturer)
		} else if strings.HasPrefix(line, "\tProduct Name:") {
			info.Model = strings.TrimPrefix(line, "\tProduct Name:")
			info.Model = strings.TrimSpace(info.Model)
		} else if strings.HasPrefix(line, "\tSerial Number:") {
			info.SerialNumber = strings.TrimPrefix(line, "\tSerial Number:")
			info.SerialNumber = strings.TrimSpace(info.SerialNumber)
		}
	}

	return info, nil
}
func GetMotherboardName() (string, error) {
	BoardInfo, err := GetMotherboardInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}
	return fmt.Sprintf("%s %s", BoardInfo.Manufacturer, BoardInfo.Model), nil
}