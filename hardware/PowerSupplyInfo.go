package hardware

import (
	"os/exec"
	"strings"
)

// PowerSupplyInfo 包含电源的信息
type PowerSupplyInfo struct {
	Manufacturer string
	Model        string
	SN           string
}

// getPowerSupplyInfo 获取电源的型号和SN号
func GetPowerSupplyInfo() ([]PowerSupplyInfo, error) {
	cmd := exec.Command("dmidecode", "-t", "39")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	powerSupplyOutStr := string(out)
	powerSupplySections := strings.Split(powerSupplyOutStr, "\n\n")
	info := make([]PowerSupplyInfo, 0)

	for _, section := range powerSupplySections {
		if strings.Contains(section, "System Power Supply") {
			lines := strings.Split(section, "\n")

			powerSupply := PowerSupplyInfo{}

			for _, line := range lines {
				if strings.Contains(line, "Manufacturer:") {
					powerSupply.Manufacturer = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
				} else if strings.Contains(line, "Name:") {
					powerSupply.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
				} else if strings.Contains(line, "Serial Number:") {
					powerSupply.SN = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
				}
			}

			if powerSupply.Manufacturer != "" || powerSupply.Model != "" || powerSupply.SN != "" {
				info = append(info, powerSupply)
			}
		}
	}

	return info, nil
}