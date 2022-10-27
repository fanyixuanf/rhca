/*
@Time : 2022/10/26 21:46
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : composite_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Structural_Pattern

import (
	"fmt"
	"testing"
)

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

type component struct {
	parent Component
	name string
}

func (c *component)Parent() Component {
	return c.parent
}

func (c *component)SetParent(p Component) {
	c.parent = p
}

func (c *component)Name() string {
	return c.name
}

func (c *component)SetName(s string) {
	c.name = s
}

func (c *component)AddChild(co Component) {

}

func (c *component)Print(p string) {

}

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (l *Leaf)Print(p string) {
	fmt.Printf("%s-%s\n", p, l.name)
}

type Composite struct {
	component
	childs []Component
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func (c *Composite)AddChild(child Component){
	child.SetParent(c)
	c.childs = append(c.childs, child)
}

func (c *Composite)Print(p string) {
	fmt.Printf("%s+%s\n", p, c.name)
	p += " "
	for _, comp := range c.childs {
		comp.Print(p)
	}
}

func TestMyfunc(t *testing.T) {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(l1)
	c1.AddChild(l2)
	c2.AddChild(c3)
	c2.AddChild(l3)
	root.Print("")

}

