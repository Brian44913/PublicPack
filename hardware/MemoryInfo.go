package PublicPackHardware

import (
	"os/exec"
	"strings"
)

// MemoryInfo 包含了内存的信息
type MemoryInfo struct {
	Model  string
	SN     string
	Size   string
	Speed  string
	PartNumber string
}

// getMemoryInfo 获取内存的信息
func GetMemoryInfo() ([]MemoryInfo, error) {
	cmd := exec.Command("dmidecode", "-t", "memory")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	memoryOutStr := string(out)
	memoryLines := strings.Split(memoryOutStr, "\n")

	info := make([]MemoryInfo, 0)
	memInfo := MemoryInfo{}

	for _, line := range memoryLines {
		if strings.Contains(line, "Size:") && !strings.Contains(line, "No Module Installed") {
			memInfo.Size = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		} else if strings.Contains(line, "Manufacturer:") {
			memInfo.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		} else if strings.Contains(line, "Serial Number:") {
			memInfo.SN = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		} else if strings.Contains(line, "Speed:") {
			memInfo.Speed = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		} else if strings.Contains(line, "Part Number:") {
			memInfo.PartNumber = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		}

		// 如果我们找到了一个完整的内存模块信息，将其添加到列表中并重置内存信息
		if memInfo.Model != "" && memInfo.SN != "" && memInfo.Size != "" && memInfo.Speed != "" {
			info = append(info, memInfo)
			memInfo = MemoryInfo{}
		}
	}

	return info, nil
}