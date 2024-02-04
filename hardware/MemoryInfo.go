package hardware

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
    cmd := exec.Command("/usr/sbin/dmidecode", "-t", "memory")
    out, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    memoryOutStr := string(out)
    memoryLines := strings.Split(memoryOutStr, "\n")

    info := make([]MemoryInfo, 0)
    var memInfo *MemoryInfo

    for _, line := range memoryLines {
        if strings.Contains(line, "Memory Device") {
            if memInfo != nil && memInfo.Size != "" && memInfo.Size != "No Module Installed" {
                info = append(info, *memInfo)
            }
            memInfo = &MemoryInfo{}
        }

        if memInfo == nil {
            continue
        }

        if strings.Contains(line, "Size:") {
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
    }

    // 添加最后一个内存模块
    if memInfo != nil && memInfo.Size != "" && memInfo.Size != "No Module Installed" {
        info = append(info, *memInfo)
    }

    return info, nil
}
