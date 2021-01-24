package utils

import (
	"encoding/json"
	"github.com/feivxs/cbtool/common"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path/filepath"
)

var config *common.CnBlogAccount

func GetDefaultConfigPath() string {
	home, _ := homedir.Dir()
	if exist, _ := IsPathExists(filepath.Join(home, common.DefaultConfigDir)); !exist {
		HandleErr(os.Mkdir(filepath.Join(home, common.DefaultConfigDir), 0777))
	}
	return filepath.Join(home, common.DefaultConfigDir, common.DefaultConfigName)
}

func GetConfig() common.CnBlogAccount {
	if config == nil {
		config = &common.CnBlogAccount{}
		configBytes, _ := ioutil.ReadFile(GetDefaultConfigPath())
		HandleErr(json.Unmarshal(configBytes, config))
	}
	return *config
}
