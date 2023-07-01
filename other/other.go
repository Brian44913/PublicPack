package PublicPackOther

import (
	"regexp"
    "fmt"
	"os"
	"time"
	"strings"
	"io/ioutil"
	"encoding/hex"
	"io"
	"crypto/md5"
	"path/filepath"
	"os/exec"
)

func DomainHosts(domain string,IP string) {
	HOSTS, _ := ReadAll("/etc/hosts")
	if ok, _ := regexp.MatchString(IP+" "+domain, string(HOSTS)); ok {
		return
    }else{
		if ok, _ = regexp.MatchString(domain, string(HOSTS)); ok {
			// 需要替换
			re, _ := regexp.Compile("(.+) "+domain);
			newContent := re.ReplaceAllString(string(HOSTS), IP+" "+domain);
			ioutil.WriteFile("/etc/hosts", []byte(newContent), 0644)
		}else{
			// 加一行即可
			newContent := string(HOSTS)+IP+" "+domain+"\n"
			ioutil.WriteFile("/etc/hosts", []byte(newContent), 0644)
		}
		return
	}
}
func FileExist(path string) bool {
  _, err := os.Lstat(path)
  return !os.IsNotExist(err)
}
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
func ReplaceStringByRegex(str, rule, replace string) (string, error) {
    reg, err := regexp.Compile(rule)
    if reg == nil || err != nil {
        return "", fmt.Errorf("正则MustCompile错误:" + err.Error())
    }
    return reg.ReplaceAllString(str, replace), nil
}
func GetStringFromFile(path string) string {
	//读取文件全部内容
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.Replace(string(b), "\n", "", -1)
}
func generateMD5(paths []string) (string, error) {
	hash := md5.New()
	for _, path := range paths {
		_, err := io.WriteString(hash, path)
		if err != nil {
			return "", err
		}
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
func GetDirMD5(dir string) (string, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			modTime := info.ModTime().Format(time.RFC3339)
			paths = append(paths, modTime)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	md5Value, err := generateMD5(paths)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return md5Value, nil
}
func IsCacheEx(dir string) bool {
	MD5file  := "/tmp/cacheMD5"
	MD5Info, err := os.Stat(MD5file)
	if err != nil {
		MD5,err  := GetDirMD5(dir)
		if err != nil {
			fmt.Println("Error:", err)
		}
		create, _ := os.Create(MD5file)
		_, err = create.Write([]byte(MD5))
	}else{
		modTime := MD5Info.ModTime()
		elapsed := time.Since(modTime)
		if elapsed > 40*time.Minute {
			MD5cache := GetStringFromFile(MD5file)
			MD5,err  := GetDirMD5(dir)
			if err != nil {
				fmt.Println("Error:", err)
			}
			if MD5cache==MD5 {
				_ = os.Remove(MD5file)
				return true
			}
		}else{
			fmt.Println(elapsed)
		}
	}
	return false // 未超时
}
func GetBinV(file string, cmds string) (string, error) {
	bin_stat, _ := os.Stat(file)
	bin_md5 := fmt.Sprintf("%x", md5.Sum([]byte(file+bin_stat.ModTime().Format("2006-01-02 15:04:05"))))
	bin_v, err := ReadAll("/tmp/"+bin_md5)
	if err != nil {
		cmd := exec.Command("bash", "-c", file+" "+cmds)
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}

		// 删除尾部的换行符
		bin_v2 := strings.TrimSpace(string(output))
		ioutil.WriteFile("/tmp/"+bin_md5, []byte(bin_v2), 0666)
		return bin_v2, nil
	}
	return string(bin_v), nil
}