package hardware

import (
	"os/exec"
	"strings"
)

// GPUInfo 包含了显卡的信息
type GPUInfo struct {
	Models    []string
	UUIDs     []string
	Memory    []string
}

// getGPUInfo 获取显卡的信息
func GetGPUInfo() (GPUInfo, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,uuid,memory.total", "--format=csv,noheader")
	out, err := cmd.Output()
	if err != nil {
		return GPUInfo{}, err
	}

	outStr := string(out)
	lines := strings.Split(outStr, "\n")

	models := make([]string, 0, len(lines))
	uuids := make([]string, 0, len(lines))
	memory := make([]string, 0, len(lines))

	for _, line := range lines {
		if line != "" {
			parts := strings.Split(line, ",")
			model := strings.TrimSpace(parts[0])
			uuid := strings.TrimSpace(parts[1])
			mem := strings.TrimSpace(parts[2])
			models = append(models, model)
			uuids = append(uuids, uuid)
			memory = append(memory, mem)
		}
	}

	info := GPUInfo{
		Models: models,
		UUIDs:  uuids,
		Memory: memory,
	}

	return info, nil
}