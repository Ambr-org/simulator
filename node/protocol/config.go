///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.
package protocol

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*
<app>
	<general>
		<Item1>xx</Item1>
	</general>

	<network>
		<port>7777</port>
		<maxpeers>5</maxpeers>
	</network>
	<Seeds>
		<seed>
			<dns></dns>
			<ip>192.168.1.2</ip>
			<port>7777</port>
		</seed>
	</disks>
	<cache></cache>
</app>
*/

//Network ...
type P2p struct {
	Port     uint32 `xml:"port"`
	MaxPeers uint32 `xml:"maxpeers"`
}

//General ...
type General struct {
	Item string `xml:"item"`
	//ServerName string `xml:"serverName"`
}

//Seed ...
type Seed struct {
	DNS string `xml:"dns"`
	//ServerName string `xml:"serverName"`
	IP   string `xml:"ip"`
	Port uint32 `xml:"port"`
}

//Config ...
type Config struct {
	General   General `xml:"general"`
	P2p       P2p     `xml:"p2p"`
	Seeds     []Seed  `xml:"seeds>seed"`
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
