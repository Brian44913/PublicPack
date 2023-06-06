package PublicPackageOther

import (
	"regexp"
    "fmt"
	"os"
	"io/ioutil"
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