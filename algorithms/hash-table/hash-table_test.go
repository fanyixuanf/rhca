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
	"bytes"
	"crypto/sha256"
	"encoding"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {

	hash := make(map[string]string, 3)
	hash["test"] = "test"
	fmt.Println(hash)
}

func TestHash(t *testing.T) {

	input1 := "The tunneling gopher digs downwards, "
	input2 := "unaware of what he will find."

	first := sha256.New()
	first.Write([]byte(input1))

	marshaler, ok := first.(encoding.BinaryMarshaler)
	if !ok {
		fmt.Println("first does not implement encoding.BinaryMarshaler")
	}
	state, err := marshaler.MarshalBinary()
	if err != nil {
		fmt.Println("unable to marshal hash:", err)
	}

	second := sha256.New()

	unmarshaler, ok := second.(encoding.BinaryUnmarshaler)
	if !ok {
		fmt.Println("second does not implement encoding.BinaryUnmarshaler")
	}
	if err := unmarshaler.UnmarshalBinary(state); err != nil {
		fmt.Println("unable to unmarshal hash:", err)
	}

	first.Write([]byte(input2))
	second.Write([]byte(input2))

	fmt.Printf("%x\n", first.Sum(nil))
	fmt.Printf("%x\n", second.Sum(nil))
	fmt.Println(bytes.Equal(first.Sum(nil), second.Sum(nil)))
}

func TestHash2(t *testing.T) {
	h := NewHashTable(10)
	h.Put("name", "Alice")
	h.Put("age", 30)
	h.Put("address", "123 Main St")
	fmt.Println(h.Get("name"))     // Output: Alice true
	fmt.Println(h.Get("age"))      // Output: 30 true
	fmt.Println(h.Get("address"))  // Output: 123 Main St true
	fmt.Println(h.Get("phone"))    // Output: <nil> false
	fmt.Println(h.Remove("phone")) // Output: false
	fmt.Println(h.Remove("age"))   // Output: true
	fmt.Println(h.Get("age"))      // Output: <nil> false
}
