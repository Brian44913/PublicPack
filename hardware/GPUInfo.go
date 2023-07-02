package PublicPackHardware

import (
	"os/exec"
	"strings"
)

// GPUInfo 包含了显卡的信息
type GPUInfo struct {
	Models []string
}

// getGPUInfo 获取显卡的信息
func GetGPUInfo() (GPUInfo, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=name", "--format=csv,noheader")
	out, err := cmd.Output()
	if err != nil {
		return GPUInfo{}, err
	}

	outStr := string(out)
	lines := strings.Split(outStr, "\n")

	models := make([]string, 0, len(lines))
	for _, line := range lines {
		if line != "" {
			models = append(models, strings.TrimSpace(line))
		}
	}

	info := GPUInfo{
		Models: models,
	}

	return info, nil
}