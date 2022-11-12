/*
@Time : 2022/11/10 20:29
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : hash-table_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package hash_table

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {

	hash := make(map[string]string, 3)
	hash["test"] = "test"
	fmt.Println(hash)
}