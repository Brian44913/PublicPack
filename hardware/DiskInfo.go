package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type lsblkOutput struct {
	BlockDevices []struct {
		Name string `json:"name"`
		Size string `json:"size"`
	} `json:"blockdevices"`
}

// DiskInfo 包含了磁盘的信息
type DiskInfo struct {
	Model string
	SN    string
	Size  string
}

// GetDiskInfo 获取磁盘的信息
func GetDiskInfo() ([]DiskInfo, error) {
    cmd := exec.Command("lsblk", "-d", "-o", "NAME,SIZE", "--json")
    out, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    var output lsblkOutput
    if err := json.NewDecoder(bytes.NewReader(out)).Decode(&output); err != nil {
        return nil, err
    }

    info := make([]DiskInfo, 0, len(output.BlockDevices))

    for _, blockDevice := range output.BlockDevices {
        if strings.HasPrefix(blockDevice.Name, "loop") || strings.HasPrefix(blockDevice.Name, "md") ||
            strings.HasPrefix(blockDevice.Name, "sr") || strings.HasPrefix(blockDevice.Name, "ram") {
            continue
        }
		if blockDevice.Size == "0B" {
			continue
		}

        disk := "/dev/" + blockDevice.Name
        cmd = exec.Command("/usr/sbin/smartctl", "-i", disk)
        out, err = cmd.CombinedOutput()
        if err != nil {
            if strings.Contains(string(out), "megaraid") {
                for i := 1; i <= 2; i++ {
                    cmd = exec.Command("/usr/sbin/smartctl", "-i", disk, "-d", fmt.Sprintf("megaraid,%d", i))
                    out, err = cmd.CombinedOutput()
                    if err != nil {
                        return nil, fmt.Errorf("smartctl error for %s with megaraid,%d: %v, output: %s", disk, i, err, string(out))
                    }
                    info = append(info, parseDiskInfo(blockDevice.Size,out))
                }
            } else {
                return nil, fmt.Errorf("smartctl error for %s: %v, output: %s", disk, err, string(out))
            }
        } else {
            info = append(info, parseDiskInfo(blockDevice.Size,out))
        }
    }

    return info, nil
}

// parseDiskInfo 解析磁盘信息
func parseDiskInfo(size string, out []byte) DiskInfo {
    diskOutStr := string(out)
    diskLines := strings.Split(diskOutStr, "\n")
    diskInfo := DiskInfo{
		Size: size,
	}

    for _, line := range diskLines {
        if strings.Contains(line, "Model Number") {
            diskInfo.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
        } else if strings.Contains(line, "Device Model") && diskInfo.Model == "" {
            // 只有在 Model Number 为空的情况下才考虑 Device Model
            diskInfo.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
        } else if strings.Contains(line, "Serial Number") {
            diskInfo.SN = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
        }
    }

    return diskInfo
}

/*
// getDiskInfo 获取磁盘的信息
func GetDiskInfo() ([]DiskInfo, error) {
	cmd := exec.Command("lsblk", "-d", "-o", "NAME,SIZE", "--json")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var output lsblkOutput
	if err := json.NewDecoder(bytes.NewReader(out)).Decode(&output); err != nil {
		return nil, err
	}

	info := make([]DiskInfo, 0, len(output.BlockDevices))

	for _, blockDevice := range output.BlockDevices {
		if strings.HasPrefix(blockDevice.Name, "loop") || strings.HasPrefix(blockDevice.Name, "md") ||
			strings.HasPrefix(blockDevice.Name, "sr") || strings.HasPrefix(blockDevice.Name, "ram") {
			continue
		}

		disk := "/dev/" + blockDevice.Name
		cmd = exec.Command("/usr/sbin/smartctl", "-i", disk)
		out, err = cmd.CombinedOutput()
		if err != nil {
			// return nil, fmt.Errorf("smartctl error for %s: %v, output: %s", disk, err, string(out))
			
			// 检查错误输出中是否包含 "megaraid" 关键字
			if strings.Contains(string(out), "megaraid") {
				for i := 1; i <= 2; i++ {
					cmd = exec.Command("/usr/sbin/smartctl", "-i", disk, "-d", fmt.Sprintf("megaraid,%d", i))
					out, err = cmd.CombinedOutput()
					if err == nil {
						// break // 如果没有错误，跳出循环
						diskOutStr := string(out)
						diskLines := strings.Split(diskOutStr, "\n")
						diskInfo := DiskInfo{
							Size: blockDevice.Size,
						}

						for _, line := range diskLines {
							if strings.Contains(line, "Model Number") {
								diskInfo.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
							} else if strings.Contains(line, "Device Model") {
								diskInfo.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
							} else if strings.Contains(line, "Serial Number") {
								diskInfo.SN = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
							}
						}

						info = append(info, diskInfo)
					}
				}

				// 如果尝试了所有 megaraid 参数后仍然出错，则返回错误
				if err != nil {
					return nil, fmt.Errorf("smartctl error for %s with megaraid: %v, output: %s", disk, err, string(out))
				}
			} else {
				// 如果错误不包含 "megaraid" 关键字，直接返回错误
				return nil, fmt.Errorf("smartctl error for %s: %v, output: %s", disk, err, string(out))
			}
			
		}

		diskOutStr := string(out)
		diskLines := strings.Split(diskOutStr, "\n")
		diskInfo := DiskInfo{
			Size: blockDevice.Size,
		}

		for _, line := range diskLines {
			if strings.Contains(line, "Model Number") {
				diskInfo.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			} else if strings.Contains(line, "Device Model") {
				diskInfo.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			} else if strings.Contains(line, "Serial Number") {
				diskInfo.SN = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			}
		}

		info = append(info, diskInfo)
	}

	return info, nil
}
*/