package hardware

import (
	"encoding/json"
    "io/ioutil"
    "os"
    "time"
)

type JsonHardware struct {
	BoardInfo				MotherboardInfo			`json:"BoardInfo"`
	CPUInfo					CPUInfo					`json:"CPUInfo"`
	GPUInfo					GPUInfo					`json:"GPUInfo"`
	DiskInfo				[]DiskInfo					`json:"DiskInfo"`
	MemoryInfo				[]MemoryInfo				`json:"MemoryInfo"`
	PowerSupplyInfo			[]PowerSupplyInfo			`json:"PowerSupplyInfo"`
}

func GetHardwareInfo() (JsonHardware, error){
	BoardInfo, err := GetMotherboardInfo()
	if err != nil {
		return JsonHardware{}, err
	}

	DiskInfo, err := GetDiskInfo()
	if err != nil {
		return JsonHardware{}, err
	}
	
	CPUInfo, err := GetCPUInfo()
	if err != nil {
		return JsonHardware{}, err
	}
	
	GPUInfo, err := GetGPUInfo()
	if err != nil {
		return JsonHardware{}, err
	}
	
	MemoryInfo, err := GetMemoryInfo()
	if err != nil {
		return JsonHardware{}, err
	}
	
	PowerSupplyInfo, err := GetPowerSupplyInfo()
	if err != nil {
		return JsonHardware{}, err
	}
	
	jsonHardware := JsonHardware{
		BoardInfo:			BoardInfo,
		CPUInfo:			CPUInfo,
		GPUInfo:			GPUInfo,
		DiskInfo:			DiskInfo,
		MemoryInfo:			MemoryInfo,
		PowerSupplyInfo:	PowerSupplyInfo,
	}
	return jsonHardware, nil
}
// 本地缓存硬件信息
func GetHardwareInfo2(filePath string, cacheDuration int) (JsonHardware, error) {
    var hardwareInfo JsonHardware

    fileInfo, err := os.Stat(filePath)
    if os.IsNotExist(err) {
        // 文件不存在，获取新的硬件信息
        newHardwareInfo, err := GetHardwareInfo()
        if err != nil {
            return hardwareInfo, err
        }

        // 将新的硬件信息转换为 JSON 并写入新文件
        jsonBytes, err := json.Marshal(newHardwareInfo)
        if err != nil {
            return hardwareInfo, err
        }

        err = ioutil.WriteFile(filePath, jsonBytes, 0644)
        if err != nil {
            return hardwareInfo, err
        }
        return newHardwareInfo, nil
    } else if err != nil {
        // 其他错误
        return hardwareInfo, err
    }

    // 如果文件存在且最后修改时间超过指定的缓存时间
    if time.Since(fileInfo.ModTime()) > time.Duration(cacheDuration)*time.Minute {
        newHardwareInfo, err := GetHardwareInfo()
        if err != nil {
            return hardwareInfo, err
        }

        jsonBytes, err := json.Marshal(newHardwareInfo)
        if err != nil {
            return hardwareInfo, err
        }

        fileContents, err := ioutil.ReadFile(filePath)
        if err != nil {
            return hardwareInfo, err
        }

        if string(fileContents) == string(jsonBytes) {
            // 内容相同，更新时间戳
            current := time.Now().Local()
            os.Chtimes(filePath, current, current)
        } else {
            // 内容不同，更新文件
            err = ioutil.WriteFile(filePath, jsonBytes, 0644)
            if err != nil {
                return hardwareInfo, err
            }
        }
        return newHardwareInfo, nil
    }

    // 文件最后修改时间在指定的缓存时间以内，读取并返回文件内容
    return JsonHardware{}, nil
}