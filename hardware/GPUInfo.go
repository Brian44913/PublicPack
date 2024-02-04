package hardware

import (
	"fmt"
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

    // 检查命令执行是否出错
    if err != nil {
        // 不论错误的原因是什么（包括 nvidia-smi 命令不存在或执行失败）
        // 返回空的 GPUInfo 对象而不是错误
        return GPUInfo{}, nil
    }

    // 以下是正常解析输出的代码
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
func GetGPUName() (string, error) {
	GPUInfo, err := GetGPUInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	return strings.Join(GPUInfo.Models, ", "), nil
}