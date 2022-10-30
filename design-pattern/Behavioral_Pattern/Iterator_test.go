/*
@Time : 2022/10/30 15:11
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : Iterator_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

type Numbers struct {
	start,
	end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end: end,
	}
}

type NumbersIterator struct {
	numbers *Numbers
	next int
}

func (i *NumbersIterator)First() {
	i.next = i.numbers.start
}

func (i *NumbersIterator) IsDone() bool {
	return i.next > i.numbers.end
}

func (i *NumbersIterator) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}

func TestIterator(t *testing.T) {
	var agg Aggregate
	agg = NewNumbers(1, 10)
	IteratorPrint(agg.Iterator())

}