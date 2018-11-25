/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package main is the entrypoint for the orderer binary
// and calls only into the server.Main() function.  No other
// function should be included in this package.
package main

import "github.com/hyperledger/fabric/orderer/common/server"

/*
	HHH fabric/orderer 这个模块下面，有main.go, 同时也有 main 函数，应该就是编译时候的一个单位。
    orderer/consensus/kafka  这个目录，实现一个一些 consenter 接口等，会被调用。
    那个通用的 consenter 接口，solo 模式应该也会被调用
 */
func main() {
	server.Main()
}
