package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*
<vm>
	<general>
		<dbpath></dbpath>
		<abifile>xx</abifile>
		<binfile>xx</binfile>
	</general>
	<cache></cache>
</vm>
*/

//General ...
type General struct {
	DBFile string `xml:"dbpath"`
	ABI    string `xml:"abifile"`
	BIN    string `xml:"binfile"`
	//ServerName string `xml:"serverName"`
}

//Config ...
type Config struct {
	General   General `xml:"general"`
	CacheSize int     `xml:"cache"`
}

const (
	DefaultConfigFile = "cfg.txt"
)

var (
	config *Config
)

//InitConfig ...
func InitConfig(file string) bool {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	config = new(Config)
	err = xml.Unmarshal(body, config)
	fmt.Println(err, config)
	return err == nil
}

//GetConfig ...
func GetConfig() *Config {
	if config == nil {
		if !InitConfig(DefaultConfigFile) {
			return nil
		}
	}
	return config
}

//GetConfig ...
func GetConfig_mock(file string) *Config {
	if config == nil {
		if !InitConfig(file) {
			return nil
		}
	}
	return config
}
