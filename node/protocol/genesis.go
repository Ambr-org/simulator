///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

const (
	TOTAL = 6300 * 10000 * 10000
)

//private
//Kf+BAwEBC1ByaXZhdGVEYXRhAf+CAAECAQNLZXkB-4QAAQFEAf+GAAAAH-+DAwEBA0tleQH-hAABAgEBWAH-hgABAVkB-4YAAAAK-4UFAQL-iAAAAG7-ggEBIQK48XD0trZV1u7zKfaV1Cfkby6bb7Lf0shiXErFsosmawEhAssrZfXyaK9gSjk5EwoTBCZ9VGczNeWqSrK9oBPXPA7eAAEhAoNdTtiL8+rG75hLtfqGdJjh2IgQ21YDCf-QzglHxyQ1AA==
//public
//3zWsYZm8oKbErrhKt38CA7go1CupRk2i5zHrvLqf1Y2t37WN1jzGHBGuXA5Cgi1wKuNVUsefunbMRHu1rc3eRmV3sdX31b8WosGMUriqF7yrEzgUffe3ZLhvNpqvB4vrFqMmVWq4gKN7143AWm4YK4UP1cniesSP
var (
	genesisUnit *Unit
)

func initGenesis() error {
	u, e := NewUnit(nil, HashKeyType{Value: ""}, UnitGenesis, TOTAL)
	if e != nil {
		return e
	}

	genesisUnit = u
	return nil
}

func GetGenesisUnit() (*Unit, error) {
	if genesisUnit == nil {
		err := initGenesis()
		if err != nil {
			return nil, err
		}
	}

	return genesisUnit, nil
}
