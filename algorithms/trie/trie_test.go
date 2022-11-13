/*
@Time : 2022/11/12 21:24
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : trie_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package trie

import (
	"fmt"
	"testing"
)

func TestTrie(tt *testing.T) {
	t := new(TrieTree)
	t.AddTerm("hi")
	t.AddTerm("how")
	t.AddTerm("her")
	t.AddTerm("hello")
	t.AddTerm("see")
	t.AddTerm("so")

	terms := t.Retrieve("he")
	fmt.Println(terms)
	terms = t.Retrieve("bu")
	fmt.Println(terms)
}
