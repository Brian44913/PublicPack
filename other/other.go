package PublicPackOther

import (
	"regexp"
    "fmt"
	"os"
	"strings"
	"io/ioutil"
	"crypto/md5"
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