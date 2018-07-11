///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

//unify compare interface definition
//refs: public key or private key
type Compare interface {
	Equals(a interface{}) bool
}
