package setting

import (
	"io/ioutil"
	"os"

	"simpleblockchain/modules/log"

	"github.com/BurntSushi/toml"
)

// Config : 配置对象
type Config struct {
	Logs map[string]map[string]log.LogService
}

// Conf : 配置对象
var Conf Config

// InitConf : 初始化
func InitConf(confPath string) (err error) {
	contents, err := ReplaceEnvsFile(confPath)
	if err != nil {
		return err
	}

	if _, err = toml.Decode(contents, &Conf); err != nil {
		return err
	}

	return nil
}

// ReplaceEnvsFile : 匹配配置信息
func ReplaceEnvsFile(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return os.ExpandEnv(string(contents)), nil
}
