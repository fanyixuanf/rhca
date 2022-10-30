/*
@Time : 2022/10/30 14:37
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : interpreter_test.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type Node interface {
	Interpret() int
}

type ValNode struct {
	val int
}

func (v *ValNode) Interpret() int {
	return v.val
}

type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

type MainNode struct {
	left, right Node
}

func (n *MainNode)Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}


type Parser struct {
	exp []string
	index int
	prev Node
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")
	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newMinNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{
		left: p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newMinNode() Node {
	p.index++
	return &MainNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newValNode() Node {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{val: v}
}

func (p *Parser) Result() Node {
	return p.prev
}

func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 2 + 3 + 4 - 5")
	res := p.Result().Interpret()
	fmt.Println(res)
}