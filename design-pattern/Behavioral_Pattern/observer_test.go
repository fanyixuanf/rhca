/*
@Time : 2022/10/30 17:08
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : observer_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type Observer interface {
	Update(*Subject)
}

type Subject struct {
	observer []Observer
	context string
}

func NewSubject() *Subject {
	return &Subject{
		observer: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observer = append(s.observer, o)
}

func (s *Subject) notify() {
	for _, o := range s.observer {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s \n", r.name, s.context)
}

func TestObserver(t *testing.T) {
	sub := NewSubject()
	r1 := NewReader("reader1")
	r2 := NewReader("reader2")
	r3 := NewReader("reader3")

	sub.Attach(r1)
	sub.Attach(r2)
	sub.Attach(r3)
	sub.UpdateContext("observer mode")
}
