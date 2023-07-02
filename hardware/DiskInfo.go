package PublicPackHardware

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
		cmd = exec.Command("smartctl", "-i", disk)
		out, err = cmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("smartctl error for %s: %v, output: %s", disk, err, string(out))
		}

		diskOutStr := string(out)
		diskLines := strings.Split(diskOutStr, "\n")
		diskInfo := DiskInfo{
			Size: blockDevice.Size,
		}

		for _, line := range diskLines {
			if strings.Contains(line, "Model Number") {
				diskInfo.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			} else if strings.Contains(line, "Serial Number") {
				diskInfo.SN = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			}
		}

		info = append(info, diskInfo)
	}

	return info, nil
}