/*
@Time : 2022/10/19 20:39
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : factorymethod
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package creation_pattern

import (
	"fmt"
	"testing"
)

type Log interface {
	WriteLog()
}

type LogFactory interface {
	CreateLog() Log
}

type FileLog struct {
	Log
}

func (f *FileLog)WriteLog() {
	fmt.Println("File log ...")
}

type DatabaseLog struct {

}

func (d *DatabaseLog)WriteLog() {
	fmt.Println("Database Log ....")
}

type FileLogFactory struct {
	LogFactory
}

func (ff *FileLogFactory)CreateLog() Log {
	var l Log
	l = new(FileLog)
	return l
}

type DatabaseLogFactory struct {

}

func (df *DatabaseLogFactory) CreateLog() Log {
	var d Log
	d = new(DatabaseLog)
	return d
}

func TestFactoryMethod(t *testing.T) {
	var ff LogFactory
	ff = new(FileLogFactory)
	var f Log
	f = ff.CreateLog()
	f.WriteLog()
	t.Log("-----------")
	var bf LogFactory
	bf = new(DatabaseLogFactory)
	var b Log
	b = bf.CreateLog()
	b.WriteLog()
}

