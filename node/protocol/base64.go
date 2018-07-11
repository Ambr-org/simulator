///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"encoding/base64"
	"strings"
)

func Base64Encode(src []byte) string {
	ret := base64.StdEncoding.EncodeToString(src)
	return strings.Replace(ret, "/", "-", -1)
}

func Base64Decode(src string) ([]byte, error) {
	dst := strings.Replace(src, "-", "/", -1)
	return base64.StdEncoding.DecodeString(dst)
}
