package PublicPackHardware

import (
	"os/exec"
	"strings"
)

// CPUInfo 包含了CPU的信息
type CPUInfo struct {
	Model string
}

// getCPUInfo 获取CPU的信息
func GetCPUInfo() (CPUInfo, error) {
	cmd := exec.Command("lscpu")
	out, err := cmd.Output()
	if err != nil {
		return CPUInfo{}, err
	}

	outStr := string(out)
	lines := strings.Split(outStr, "\n")

	info := CPUInfo{}

	for _, line := range lines {
		if strings.HasPrefix(line, "Model name:") {
			info.Model = strings.TrimPrefix(line, "Model name:")
			info.Model = strings.TrimSpace(info.Model)
			break
		}
	}

	return info, nil
}