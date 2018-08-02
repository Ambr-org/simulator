///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

type NetworkHelper struct {
}

func NewNetworkHelper() *NetworkHelper {
	return &NetworkHelper{}
}

func (p *NetworkHelper) Publish(msg string) error {
	return nil
}

func (p *NetworkHelper) Vote(msg string) error {
	return nil
}
